// Code generated by counterfeiter. DO NOT EDIT.
package kawasakifakes

import (
	"sync"

	"github.com/concourse/guardian/kawasaki"
	"code.cloudfoundry.org/lager"
)

type FakeHostConfigurer struct {
	ApplyStub        func(logger lager.Logger, cfg kawasaki.NetworkConfig, pid int) error
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		logger lager.Logger
		cfg    kawasaki.NetworkConfig
		pid    int
	}
	applyReturns struct {
		result1 error
	}
	applyReturnsOnCall map[int]struct {
		result1 error
	}
	DestroyStub        func(cfg kawasaki.NetworkConfig) error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		cfg kawasaki.NetworkConfig
	}
	destroyReturns struct {
		result1 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHostConfigurer) Apply(logger lager.Logger, cfg kawasaki.NetworkConfig, pid int) error {
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		logger lager.Logger
		cfg    kawasaki.NetworkConfig
		pid    int
	}{logger, cfg, pid})
	fake.recordInvocation("Apply", []interface{}{logger, cfg, pid})
	fake.applyMutex.Unlock()
	if fake.ApplyStub != nil {
		return fake.ApplyStub(logger, cfg, pid)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.applyReturns.result1
}

func (fake *FakeHostConfigurer) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeHostConfigurer) ApplyArgsForCall(i int) (lager.Logger, kawasaki.NetworkConfig, int) {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return fake.applyArgsForCall[i].logger, fake.applyArgsForCall[i].cfg, fake.applyArgsForCall[i].pid
}

func (fake *FakeHostConfigurer) ApplyReturns(result1 error) {
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHostConfigurer) ApplyReturnsOnCall(i int, result1 error) {
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHostConfigurer) Destroy(cfg kawasaki.NetworkConfig) error {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		cfg kawasaki.NetworkConfig
	}{cfg})
	fake.recordInvocation("Destroy", []interface{}{cfg})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub(cfg)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.destroyReturns.result1
}

func (fake *FakeHostConfigurer) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeHostConfigurer) DestroyArgsForCall(i int) kawasaki.NetworkConfig {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return fake.destroyArgsForCall[i].cfg
}

func (fake *FakeHostConfigurer) DestroyReturns(result1 error) {
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHostConfigurer) DestroyReturnsOnCall(i int, result1 error) {
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHostConfigurer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHostConfigurer) recordInvocation(key string, args []interface{}) {
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

var _ kawasaki.HostConfigurer = new(FakeHostConfigurer)
