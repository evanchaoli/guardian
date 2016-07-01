package kawasaki

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/cloudfoundry-incubator/garden"
	"github.com/cloudfoundry-incubator/guardian/gardener"
	"github.com/cloudfoundry-incubator/guardian/kawasaki/subnets"
	"github.com/pivotal-golang/lager"
)

// generic gardener properties
const containerIpKey = gardener.ContainerIPKey
const bridgeIpKey = gardener.BridgeIPKey
const externalIpKey = gardener.ExternalIPKey

// kawasaki-specific state properties
const hostIntfKey = "kawasaki.host-interface"
const containerIntfKey = "kawasaki.container-interface"
const bridgeIntfKey = "kawasaki.bridge-interface"
const subnetKey = "kawasaki.subnet"
const iptablePrefixKey = "kawasaki.iptable-prefix"
const iptableInstanceKey = "kawasaki.iptable-inst"
const mtuKey = "kawasaki.mtu"
const dnsServerKey = "kawasaki.dns-servers"

//go:generate counterfeiter . SpecParser

type SpecParser interface {
	Parse(log lager.Logger, spec string) (subnets.SubnetSelector, subnets.IPSelector, error)
}

//go:generate counterfeiter . ConfigCreator

type ConfigCreator interface {
	Create(log lager.Logger, handle string, subnet *net.IPNet, ip net.IP) (NetworkConfig, error)
}

//go:generate counterfeiter . Configurer

type Configurer interface {
	Apply(log lager.Logger, cfg NetworkConfig, nsPath, bundlePath string) error
	DestroyBridge(log lager.Logger, cfg NetworkConfig) error
	DestroyIPTablesRules(log lager.Logger, cfg NetworkConfig) error
}

//go:generate counterfeiter . ConfigStore

type ConfigStore interface {
	Set(handle string, name string, value string)
	Get(handle string, name string) (string, bool)
}

//go:generate counterfeiter . PortPool

type PortPool interface {
	Acquire() (uint32, error)
	Release(uint32)
	Remove(uint32) error
}

//go:generate counterfeiter . PortForwarder

type PortForwarder interface {
	Forward(spec PortForwarderSpec) error
}

type PortForwarderSpec struct {
	InstanceID  string
	FromPort    uint32
	ToPort      uint32
	ContainerIP net.IP
	ExternalIP  net.IP
}

//go:generate counterfeiter . FirewallOpener

type FirewallOpener interface {
	Open(log lager.Logger, instance string, rule garden.NetOutRule) error
}

//go:generate counterfeiter . NetworkHooker

type NetworkHooker interface {
	Hooks(log lager.Logger, containerSpec garden.ContainerSpec) (gardener.Hooks, error)
}

//go:generate counterfeiter . Networker

type Networker interface {
	Capacity() uint64
	Network(log lager.Logger, spec garden.ContainerSpec, pid int, bundlePath string) error
	Destroy(log lager.Logger, handle string) error
	NetIn(log lager.Logger, handle string, externalPort, containerPort uint32) (uint32, uint32, error)
	NetOut(log lager.Logger, handle string, rule garden.NetOutRule) error
	Restore(log lager.Logger, handle string) error
}

type networker struct {
	kawasakiBinPath string // path to a binary that will apply the configuration
	iptablesBin     string

	specParser     SpecParser
	subnetPool     subnets.Pool
	configCreator  ConfigCreator
	configStore    ConfigStore
	portForwarder  PortForwarder
	portPool       PortPool
	firewallOpener FirewallOpener
	configurer     Configurer
}

func New(
	kawasakiBinPath string,
	iptablesBin string,
	specParser SpecParser,
	subnetPool subnets.Pool,
	configCreator ConfigCreator,
	configStore ConfigStore,
	configurer Configurer,
	portPool PortPool,
	portForwarder PortForwarder,
	firewallOpener FirewallOpener,
) *networker {
	return &networker{
		kawasakiBinPath: kawasakiBinPath,
		iptablesBin:     iptablesBin,

		specParser:    specParser,
		subnetPool:    subnetPool,
		configCreator: configCreator,
		configStore:   configStore,
		configurer:    configurer,

		portForwarder: portForwarder,
		portPool:      portPool,

		firewallOpener: firewallOpener,
	}
}

func (n *networker) Network(log lager.Logger, containerSpec garden.ContainerSpec, pid int, bundlePath string) error {
	log = log.Session("network", lager.Data{
		"handle": containerSpec.Handle,
		"spec":   containerSpec.Network,
	})

	log.Info("started")
	defer log.Info("finished")

	subnetReq, ipReq, err := n.specParser.Parse(log, containerSpec.Network)
	if err != nil {
		log.Error("parse-failed", err)
		return err
	}

	subnet, ip, err := n.subnetPool.Acquire(log, subnetReq, ipReq)
	if err != nil {
		log.Error("acquire-failed", err)
		return err
	}

	config, err := n.configCreator.Create(log, containerSpec.Handle, subnet, ip)
	if err != nil {
		log.Error("create-config-failed", err)
		return fmt.Errorf("create network config: %s", err)
	}
	log.Info("config-create", lager.Data{"config": config})

	if err := save(n.configStore, containerSpec.Handle, config); err != nil {
		return err
	}

	if err := n.configurer.Apply(log, config, fmt.Sprintf("/proc/%d/ns/net", pid), bundlePath); err != nil {
		return err
	}
	return nil
}

// Capacity returns the number of subnets this network can host
func (n *networker) Capacity() uint64 {
	return uint64(n.subnetPool.Capacity())
}

func (n *networker) NetIn(log lager.Logger, handle string, externalPort, containerPort uint32) (uint32, uint32, error) {
	cfg, err := load(n.configStore, handle)
	if err != nil {
		return 0, 0, err
	}

	if externalPort == 0 {
		externalPort, err = n.portPool.Acquire()
		if err != nil {
			return 0, 0, err
		}
	}

	if containerPort == 0 {
		containerPort = externalPort
	}

	err = n.portForwarder.Forward(PortForwarderSpec{
		InstanceID:  cfg.IPTableInstance,
		FromPort:    externalPort,
		ToPort:      containerPort,
		ContainerIP: cfg.ContainerIP,
		ExternalIP:  cfg.ExternalIP,
	})

	if err != nil {
		return 0, 0, err
	}

	if err := addPortMapping(log, n.configStore, handle, garden.PortMapping{
		HostPort:      externalPort,
		ContainerPort: containerPort,
	}); err != nil {
		return 0, 0, err
	}

	return externalPort, containerPort, nil
}

func (n *networker) NetOut(log lager.Logger, handle string, rule garden.NetOutRule) error {
	cfg, err := load(n.configStore, handle)
	if err != nil {
		return err
	}

	return n.firewallOpener.Open(log, cfg.IPTableInstance, rule)
}

func (n *networker) Destroy(log lager.Logger, handle string) error {
	cfg, err := load(n.configStore, handle)
	if err != nil {
		log.Debug("no-properties-for-container-skipping-destroy-network")
		return nil
	}

	if err := n.subnetPool.Release(cfg.Subnet, cfg.ContainerIP); err != nil && err != subnets.ErrReleasedUnallocatedSubnet {
		log.Error("release-failed", err)
		return err
	}

	if ports, ok := n.configStore.Get(handle, gardener.MappedPortsKey); ok {
		mappings, err := portsFromJson(ports)
		if err != nil {
			return err
		}

		for _, m := range mappings {
			n.portPool.Release(m.HostPort)
		}
	}

	if err := n.configurer.DestroyIPTablesRules(log, cfg); err != nil {
		return err
	}

	n.subnetPool.RunIfFree(cfg.Subnet, func() error {
		n.configurer.DestroyBridge(log, cfg)
		return nil
	})

	return nil
}

func (n *networker) Restore(log lager.Logger, handle string) error {
	networkConfig, err := load(n.configStore, handle)
	if err != nil {
		return fmt.Errorf("loading %s: %v", handle, err)
	}

	err = n.subnetPool.Remove(networkConfig.Subnet, networkConfig.ContainerIP)
	if err != nil {
		return fmt.Errorf("subnet pool removing %s: %v", handle, err)
	}

	currentMappingsJson, ok := n.configStore.Get(handle, gardener.MappedPortsKey)
	if !ok {
		return nil
	}

	currentMappings, err := portsFromJson(currentMappingsJson)
	if err != nil {
		return fmt.Errorf("unmarshaling port mappings %s: %v", handle, err)
	}

	for _, mapping := range currentMappings {
		if err = n.portPool.Remove(mapping.HostPort); err != nil {
			return fmt.Errorf("port pool removing %s: %v", handle, err)
		}
	}

	return nil
}

func addPortMapping(logger lager.Logger, configStore ConfigStore, handle string, newMapping garden.PortMapping) error {
	var currentMappings portMappingList
	if currentMappingsJson, ok := configStore.Get(handle, gardener.MappedPortsKey); ok {
		var err error
		currentMappings, err = portsFromJson(currentMappingsJson)
		if err != nil {
			return err
		}
	}

	updatedMappings := append(currentMappings, newMapping)
	configStore.Set(handle, gardener.MappedPortsKey, updatedMappings.toJson())
	return nil
}

func getAll(config ConfigStore, handle string, key ...string) (vals []string, err error) {
	for _, k := range key {
		v, ok := config.Get(handle, k)
		if !ok {
			return nil, fmt.Errorf("property not found: %s", k)
		}

		vals = append(vals, v)
	}

	return vals, nil
}

func save(config ConfigStore, handle string, netConfig NetworkConfig) error {
	config.Set(handle, hostIntfKey, netConfig.HostIntf)
	config.Set(handle, containerIntfKey, netConfig.ContainerIntf)
	config.Set(handle, bridgeIntfKey, netConfig.BridgeName)
	config.Set(handle, bridgeIpKey, netConfig.BridgeIP.String())
	config.Set(handle, containerIpKey, netConfig.ContainerIP.String())
	config.Set(handle, subnetKey, netConfig.Subnet.String())
	config.Set(handle, iptablePrefixKey, netConfig.IPTablePrefix)
	config.Set(handle, iptableInstanceKey, netConfig.IPTableInstance)
	config.Set(handle, mtuKey, strconv.Itoa(netConfig.Mtu))
	config.Set(handle, externalIpKey, netConfig.ExternalIP.String())

	var dnsServers []string
	for _, dnsServer := range netConfig.DNSServers {
		dnsServers = append(dnsServers, dnsServer.String())
	}

	config.Set(handle, dnsServerKey, strings.Join(dnsServers, ", "))

	return nil
}

func appendIfNotNil(errors []error, err error) []error {
	if err != nil {
		return append(errors, err)
	}
	return errors
}

func load(config ConfigStore, handle string) (NetworkConfig, error) {
	vals, err := getAll(config, handle, hostIntfKey, containerIntfKey, bridgeIntfKey, bridgeIpKey, containerIpKey, subnetKey, iptablePrefixKey, iptableInstanceKey, mtuKey, externalIpKey, dnsServerKey)

	if err != nil {
		return NetworkConfig{}, err
	}

	_, ipnet, err := net.ParseCIDR(vals[5])
	if err != nil {
		return NetworkConfig{}, err
	}

	mtu, err := strconv.Atoi(vals[8])
	if err != nil {
		return NetworkConfig{}, err
	}

	var dnsServers []net.IP
	for _, dnsServerName := range strings.Split(vals[10], ",") {
		dnsServerName = strings.TrimSpace(dnsServerName)
		if dnsServerName == "" {
			continue
		}
		ip := net.ParseIP(dnsServerName)
		if ip == nil {
			return NetworkConfig{}, fmt.Errorf("Failed to parse DNS server IP address %s", dnsServerName)
		}
		dnsServers = append(dnsServers, ip)
	}

	return NetworkConfig{
		HostIntf:        vals[0],
		ContainerIntf:   vals[1],
		BridgeName:      vals[2],
		BridgeIP:        net.ParseIP(vals[3]),
		ContainerIP:     net.ParseIP(vals[4]),
		ExternalIP:      net.ParseIP(vals[9]),
		Subnet:          ipnet,
		IPTablePrefix:   vals[6],
		IPTableInstance: vals[7],
		Mtu:             mtu,
		DNSServers:      dnsServers,
	}, nil
}

type portMappingList []garden.PortMapping

func (l portMappingList) toJson() string {
	b, err := json.Marshal(l)
	if err != nil {
		panic(err) // impossible, since []PortMapping is always encodable
	}

	return string(b)
}

func portsFromJson(s string) (portMappingList, error) {
	var mappings portMappingList
	if err := json.Unmarshal([]byte(s), &mappings); err != nil {
		return nil, err
	}

	return mappings, nil
}
