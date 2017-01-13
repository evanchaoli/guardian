// This file was generated by counterfeiter
package imagepluginfakes

import (
	"os/exec"
	"sync"

	"code.cloudfoundry.org/garden-shed/rootfs_provider"
	"code.cloudfoundry.org/guardian/imageplugin"
	"code.cloudfoundry.org/lager"
)

type FakeCommandCreator struct {
	CreateCommandStub        func(log lager.Logger, handle string, spec rootfs_provider.Spec) *exec.Cmd
	createCommandMutex       sync.RWMutex
	createCommandArgsForCall []struct {
		log    lager.Logger
		handle string
		spec   rootfs_provider.Spec
	}
	createCommandReturns struct {
		result1 *exec.Cmd
	}
	DestroyCommandStub        func(log lager.Logger, handle string) (*exec.Cmd, error)
	destroyCommandMutex       sync.RWMutex
	destroyCommandArgsForCall []struct {
		log    lager.Logger
		handle string
	}
	destroyCommandReturns struct {
		result1 *exec.Cmd
		result2 error
	}
	MetricsCommandStub        func(log lager.Logger, handle string) (*exec.Cmd, error)
	metricsCommandMutex       sync.RWMutex
	metricsCommandArgsForCall []struct {
		log    lager.Logger
		handle string
	}
	metricsCommandReturns struct {
		result1 *exec.Cmd
		result2 error
	}
	GCCommandStub        func(log lager.Logger) (*exec.Cmd, error)
	gCCommandMutex       sync.RWMutex
	gCCommandArgsForCall []struct {
		log lager.Logger
	}
	gCCommandReturns struct {
		result1 *exec.Cmd
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCommandCreator) CreateCommand(log lager.Logger, handle string, spec rootfs_provider.Spec) *exec.Cmd {
	fake.createCommandMutex.Lock()
	fake.createCommandArgsForCall = append(fake.createCommandArgsForCall, struct {
		log    lager.Logger
		handle string
		spec   rootfs_provider.Spec
	}{log, handle, spec})
	fake.recordInvocation("CreateCommand", []interface{}{log, handle, spec})
	fake.createCommandMutex.Unlock()
	if fake.CreateCommandStub != nil {
		return fake.CreateCommandStub(log, handle, spec)
	} else {
		return fake.createCommandReturns.result1
	}
}

func (fake *FakeCommandCreator) CreateCommandCallCount() int {
	fake.createCommandMutex.RLock()
	defer fake.createCommandMutex.RUnlock()
	return len(fake.createCommandArgsForCall)
}

func (fake *FakeCommandCreator) CreateCommandArgsForCall(i int) (lager.Logger, string, rootfs_provider.Spec) {
	fake.createCommandMutex.RLock()
	defer fake.createCommandMutex.RUnlock()
	return fake.createCommandArgsForCall[i].log, fake.createCommandArgsForCall[i].handle, fake.createCommandArgsForCall[i].spec
}

func (fake *FakeCommandCreator) CreateCommandReturns(result1 *exec.Cmd) {
	fake.CreateCommandStub = nil
	fake.createCommandReturns = struct {
		result1 *exec.Cmd
	}{result1}
}

func (fake *FakeCommandCreator) DestroyCommand(log lager.Logger, handle string) (*exec.Cmd, error) {
	fake.destroyCommandMutex.Lock()
	fake.destroyCommandArgsForCall = append(fake.destroyCommandArgsForCall, struct {
		log    lager.Logger
		handle string
	}{log, handle})
	fake.recordInvocation("DestroyCommand", []interface{}{log, handle})
	fake.destroyCommandMutex.Unlock()
	if fake.DestroyCommandStub != nil {
		return fake.DestroyCommandStub(log, handle)
	} else {
		return fake.destroyCommandReturns.result1, fake.destroyCommandReturns.result2
	}
}

func (fake *FakeCommandCreator) DestroyCommandCallCount() int {
	fake.destroyCommandMutex.RLock()
	defer fake.destroyCommandMutex.RUnlock()
	return len(fake.destroyCommandArgsForCall)
}

func (fake *FakeCommandCreator) DestroyCommandArgsForCall(i int) (lager.Logger, string) {
	fake.destroyCommandMutex.RLock()
	defer fake.destroyCommandMutex.RUnlock()
	return fake.destroyCommandArgsForCall[i].log, fake.destroyCommandArgsForCall[i].handle
}

func (fake *FakeCommandCreator) DestroyCommandReturns(result1 *exec.Cmd, result2 error) {
	fake.DestroyCommandStub = nil
	fake.destroyCommandReturns = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeCommandCreator) MetricsCommand(log lager.Logger, handle string) (*exec.Cmd, error) {
	fake.metricsCommandMutex.Lock()
	fake.metricsCommandArgsForCall = append(fake.metricsCommandArgsForCall, struct {
		log    lager.Logger
		handle string
	}{log, handle})
	fake.recordInvocation("MetricsCommand", []interface{}{log, handle})
	fake.metricsCommandMutex.Unlock()
	if fake.MetricsCommandStub != nil {
		return fake.MetricsCommandStub(log, handle)
	} else {
		return fake.metricsCommandReturns.result1, fake.metricsCommandReturns.result2
	}
}

func (fake *FakeCommandCreator) MetricsCommandCallCount() int {
	fake.metricsCommandMutex.RLock()
	defer fake.metricsCommandMutex.RUnlock()
	return len(fake.metricsCommandArgsForCall)
}

func (fake *FakeCommandCreator) MetricsCommandArgsForCall(i int) (lager.Logger, string) {
	fake.metricsCommandMutex.RLock()
	defer fake.metricsCommandMutex.RUnlock()
	return fake.metricsCommandArgsForCall[i].log, fake.metricsCommandArgsForCall[i].handle
}

func (fake *FakeCommandCreator) MetricsCommandReturns(result1 *exec.Cmd, result2 error) {
	fake.MetricsCommandStub = nil
	fake.metricsCommandReturns = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeCommandCreator) GCCommand(log lager.Logger) (*exec.Cmd, error) {
	fake.gCCommandMutex.Lock()
	fake.gCCommandArgsForCall = append(fake.gCCommandArgsForCall, struct {
		log lager.Logger
	}{log})
	fake.recordInvocation("GCCommand", []interface{}{log})
	fake.gCCommandMutex.Unlock()
	if fake.GCCommandStub != nil {
		return fake.GCCommandStub(log)
	} else {
		return fake.gCCommandReturns.result1, fake.gCCommandReturns.result2
	}
}

func (fake *FakeCommandCreator) GCCommandCallCount() int {
	fake.gCCommandMutex.RLock()
	defer fake.gCCommandMutex.RUnlock()
	return len(fake.gCCommandArgsForCall)
}

func (fake *FakeCommandCreator) GCCommandArgsForCall(i int) lager.Logger {
	fake.gCCommandMutex.RLock()
	defer fake.gCCommandMutex.RUnlock()
	return fake.gCCommandArgsForCall[i].log
}

func (fake *FakeCommandCreator) GCCommandReturns(result1 *exec.Cmd, result2 error) {
	fake.GCCommandStub = nil
	fake.gCCommandReturns = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeCommandCreator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCommandMutex.RLock()
	defer fake.createCommandMutex.RUnlock()
	fake.destroyCommandMutex.RLock()
	defer fake.destroyCommandMutex.RUnlock()
	fake.metricsCommandMutex.RLock()
	defer fake.metricsCommandMutex.RUnlock()
	fake.gCCommandMutex.RLock()
	defer fake.gCCommandMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCommandCreator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ imageplugin.CommandCreator = new(FakeCommandCreator)
