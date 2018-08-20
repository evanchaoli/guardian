// Code generated by counterfeiter. DO NOT EDIT.
package fake_subnet_pool

import (
	"net"
	"sync"

	"github.com/concourse/guardian/kawasaki/subnets"
	"code.cloudfoundry.org/lager"
)

type FakePool struct {
	AcquireStub        func(lager.Logger, subnets.SubnetSelector, subnets.IPSelector) (*net.IPNet, net.IP, error)
	acquireMutex       sync.RWMutex
	acquireArgsForCall []struct {
		arg1 lager.Logger
		arg2 subnets.SubnetSelector
		arg3 subnets.IPSelector
	}
	acquireReturns struct {
		result1 *net.IPNet
		result2 net.IP
		result3 error
	}
	acquireReturnsOnCall map[int]struct {
		result1 *net.IPNet
		result2 net.IP
		result3 error
	}
	ReleaseStub        func(*net.IPNet, net.IP) error
	releaseMutex       sync.RWMutex
	releaseArgsForCall []struct {
		arg1 *net.IPNet
		arg2 net.IP
	}
	releaseReturns struct {
		result1 error
	}
	releaseReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveStub        func(*net.IPNet, net.IP) error
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		arg1 *net.IPNet
		arg2 net.IP
	}
	removeReturns struct {
		result1 error
	}
	removeReturnsOnCall map[int]struct {
		result1 error
	}
	CapacityStub        func() int
	capacityMutex       sync.RWMutex
	capacityArgsForCall []struct{}
	capacityReturns     struct {
		result1 int
	}
	capacityReturnsOnCall map[int]struct {
		result1 int
	}
	RunIfFreeStub        func(*net.IPNet, func() error) error
	runIfFreeMutex       sync.RWMutex
	runIfFreeArgsForCall []struct {
		arg1 *net.IPNet
		arg2 func() error
	}
	runIfFreeReturns struct {
		result1 error
	}
	runIfFreeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePool) Acquire(arg1 lager.Logger, arg2 subnets.SubnetSelector, arg3 subnets.IPSelector) (*net.IPNet, net.IP, error) {
	fake.acquireMutex.Lock()
	ret, specificReturn := fake.acquireReturnsOnCall[len(fake.acquireArgsForCall)]
	fake.acquireArgsForCall = append(fake.acquireArgsForCall, struct {
		arg1 lager.Logger
		arg2 subnets.SubnetSelector
		arg3 subnets.IPSelector
	}{arg1, arg2, arg3})
	fake.recordInvocation("Acquire", []interface{}{arg1, arg2, arg3})
	fake.acquireMutex.Unlock()
	if fake.AcquireStub != nil {
		return fake.AcquireStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.acquireReturns.result1, fake.acquireReturns.result2, fake.acquireReturns.result3
}

func (fake *FakePool) AcquireCallCount() int {
	fake.acquireMutex.RLock()
	defer fake.acquireMutex.RUnlock()
	return len(fake.acquireArgsForCall)
}

func (fake *FakePool) AcquireArgsForCall(i int) (lager.Logger, subnets.SubnetSelector, subnets.IPSelector) {
	fake.acquireMutex.RLock()
	defer fake.acquireMutex.RUnlock()
	return fake.acquireArgsForCall[i].arg1, fake.acquireArgsForCall[i].arg2, fake.acquireArgsForCall[i].arg3
}

func (fake *FakePool) AcquireReturns(result1 *net.IPNet, result2 net.IP, result3 error) {
	fake.AcquireStub = nil
	fake.acquireReturns = struct {
		result1 *net.IPNet
		result2 net.IP
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePool) AcquireReturnsOnCall(i int, result1 *net.IPNet, result2 net.IP, result3 error) {
	fake.AcquireStub = nil
	if fake.acquireReturnsOnCall == nil {
		fake.acquireReturnsOnCall = make(map[int]struct {
			result1 *net.IPNet
			result2 net.IP
			result3 error
		})
	}
	fake.acquireReturnsOnCall[i] = struct {
		result1 *net.IPNet
		result2 net.IP
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePool) Release(arg1 *net.IPNet, arg2 net.IP) error {
	fake.releaseMutex.Lock()
	ret, specificReturn := fake.releaseReturnsOnCall[len(fake.releaseArgsForCall)]
	fake.releaseArgsForCall = append(fake.releaseArgsForCall, struct {
		arg1 *net.IPNet
		arg2 net.IP
	}{arg1, arg2})
	fake.recordInvocation("Release", []interface{}{arg1, arg2})
	fake.releaseMutex.Unlock()
	if fake.ReleaseStub != nil {
		return fake.ReleaseStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.releaseReturns.result1
}

func (fake *FakePool) ReleaseCallCount() int {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return len(fake.releaseArgsForCall)
}

func (fake *FakePool) ReleaseArgsForCall(i int) (*net.IPNet, net.IP) {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return fake.releaseArgsForCall[i].arg1, fake.releaseArgsForCall[i].arg2
}

func (fake *FakePool) ReleaseReturns(result1 error) {
	fake.ReleaseStub = nil
	fake.releaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePool) ReleaseReturnsOnCall(i int, result1 error) {
	fake.ReleaseStub = nil
	if fake.releaseReturnsOnCall == nil {
		fake.releaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.releaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePool) Remove(arg1 *net.IPNet, arg2 net.IP) error {
	fake.removeMutex.Lock()
	ret, specificReturn := fake.removeReturnsOnCall[len(fake.removeArgsForCall)]
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		arg1 *net.IPNet
		arg2 net.IP
	}{arg1, arg2})
	fake.recordInvocation("Remove", []interface{}{arg1, arg2})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		return fake.RemoveStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeReturns.result1
}

func (fake *FakePool) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakePool) RemoveArgsForCall(i int) (*net.IPNet, net.IP) {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return fake.removeArgsForCall[i].arg1, fake.removeArgsForCall[i].arg2
}

func (fake *FakePool) RemoveReturns(result1 error) {
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePool) RemoveReturnsOnCall(i int, result1 error) {
	fake.RemoveStub = nil
	if fake.removeReturnsOnCall == nil {
		fake.removeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePool) Capacity() int {
	fake.capacityMutex.Lock()
	ret, specificReturn := fake.capacityReturnsOnCall[len(fake.capacityArgsForCall)]
	fake.capacityArgsForCall = append(fake.capacityArgsForCall, struct{}{})
	fake.recordInvocation("Capacity", []interface{}{})
	fake.capacityMutex.Unlock()
	if fake.CapacityStub != nil {
		return fake.CapacityStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.capacityReturns.result1
}

func (fake *FakePool) CapacityCallCount() int {
	fake.capacityMutex.RLock()
	defer fake.capacityMutex.RUnlock()
	return len(fake.capacityArgsForCall)
}

func (fake *FakePool) CapacityReturns(result1 int) {
	fake.CapacityStub = nil
	fake.capacityReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakePool) CapacityReturnsOnCall(i int, result1 int) {
	fake.CapacityStub = nil
	if fake.capacityReturnsOnCall == nil {
		fake.capacityReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.capacityReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakePool) RunIfFree(arg1 *net.IPNet, arg2 func() error) error {
	fake.runIfFreeMutex.Lock()
	ret, specificReturn := fake.runIfFreeReturnsOnCall[len(fake.runIfFreeArgsForCall)]
	fake.runIfFreeArgsForCall = append(fake.runIfFreeArgsForCall, struct {
		arg1 *net.IPNet
		arg2 func() error
	}{arg1, arg2})
	fake.recordInvocation("RunIfFree", []interface{}{arg1, arg2})
	fake.runIfFreeMutex.Unlock()
	if fake.RunIfFreeStub != nil {
		return fake.RunIfFreeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.runIfFreeReturns.result1
}

func (fake *FakePool) RunIfFreeCallCount() int {
	fake.runIfFreeMutex.RLock()
	defer fake.runIfFreeMutex.RUnlock()
	return len(fake.runIfFreeArgsForCall)
}

func (fake *FakePool) RunIfFreeArgsForCall(i int) (*net.IPNet, func() error) {
	fake.runIfFreeMutex.RLock()
	defer fake.runIfFreeMutex.RUnlock()
	return fake.runIfFreeArgsForCall[i].arg1, fake.runIfFreeArgsForCall[i].arg2
}

func (fake *FakePool) RunIfFreeReturns(result1 error) {
	fake.RunIfFreeStub = nil
	fake.runIfFreeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePool) RunIfFreeReturnsOnCall(i int, result1 error) {
	fake.RunIfFreeStub = nil
	if fake.runIfFreeReturnsOnCall == nil {
		fake.runIfFreeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runIfFreeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePool) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.acquireMutex.RLock()
	defer fake.acquireMutex.RUnlock()
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	fake.capacityMutex.RLock()
	defer fake.capacityMutex.RUnlock()
	fake.runIfFreeMutex.RLock()
	defer fake.runIfFreeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePool) recordInvocation(key string, args []interface{}) {
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

var _ subnets.Pool = new(FakePool)
