// This file was generated by counterfeiter
package guardiancmdfakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/gardener"
	"code.cloudfoundry.org/guardian/guardiancmd"
)

type FakeSystemConfigurer struct {
	StartStub        func([]gardener.Starter) error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 []gardener.Starter
	}
	startReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSystemConfigurer) Start(arg1 []gardener.Starter) error {
	var arg1Copy []gardener.Starter
	if arg1 != nil {
		arg1Copy = make([]gardener.Starter, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 []gardener.Starter
	}{arg1Copy})
	fake.recordInvocation("Start", []interface{}{arg1Copy})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub(arg1)
	}
	return fake.startReturns.result1
}

func (fake *FakeSystemConfigurer) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeSystemConfigurer) StartArgsForCall(i int) []gardener.Starter {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return fake.startArgsForCall[i].arg1
}

func (fake *FakeSystemConfigurer) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSystemConfigurer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSystemConfigurer) recordInvocation(key string, args []interface{}) {
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

var _ guardiancmd.SystemConfigurer = new(FakeSystemConfigurer)