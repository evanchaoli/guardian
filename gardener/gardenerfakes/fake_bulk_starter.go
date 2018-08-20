// Code generated by counterfeiter. DO NOT EDIT.
package gardenerfakes

import (
	"sync"

	"github.com/concourse/guardian/gardener"
)

type FakeBulkStarter struct {
	StartAllStub        func() error
	startAllMutex       sync.RWMutex
	startAllArgsForCall []struct{}
	startAllReturns     struct {
		result1 error
	}
	startAllReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBulkStarter) StartAll() error {
	fake.startAllMutex.Lock()
	ret, specificReturn := fake.startAllReturnsOnCall[len(fake.startAllArgsForCall)]
	fake.startAllArgsForCall = append(fake.startAllArgsForCall, struct{}{})
	fake.recordInvocation("StartAll", []interface{}{})
	fake.startAllMutex.Unlock()
	if fake.StartAllStub != nil {
		return fake.StartAllStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.startAllReturns.result1
}

func (fake *FakeBulkStarter) StartAllCallCount() int {
	fake.startAllMutex.RLock()
	defer fake.startAllMutex.RUnlock()
	return len(fake.startAllArgsForCall)
}

func (fake *FakeBulkStarter) StartAllReturns(result1 error) {
	fake.StartAllStub = nil
	fake.startAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBulkStarter) StartAllReturnsOnCall(i int, result1 error) {
	fake.StartAllStub = nil
	if fake.startAllReturnsOnCall == nil {
		fake.startAllReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startAllReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBulkStarter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startAllMutex.RLock()
	defer fake.startAllMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBulkStarter) recordInvocation(key string, args []interface{}) {
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

var _ gardener.BulkStarter = new(FakeBulkStarter)
