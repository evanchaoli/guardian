// Code generated by counterfeiter. DO NOT EDIT.
package processwaiterfakes

import (
	"sync"

	"github.com/concourse/guardian/rundmc/peas/processwaiter"
)

type FakeProcessWaiter struct {
	Stub        func(pid int) error
	mutex       sync.RWMutex
	argsForCall []struct {
		pid int
	}
	returns struct {
		result1 error
	}
	returnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProcessWaiter) Spy(pid int) error {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		pid int
	}{pid})
	fake.recordInvocation("ProcessWaiter", []interface{}{pid})
	fake.mutex.Unlock()
	if fake.Stub != nil {
		return fake.Stub(pid)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.returns.result1
}

func (fake *FakeProcessWaiter) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *FakeProcessWaiter) ArgsForCall(i int) int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].pid
}

func (fake *FakeProcessWaiter) Returns(result1 error) {
	fake.Stub = nil
	fake.returns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProcessWaiter) ReturnsOnCall(i int, result1 error) {
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeProcessWaiter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProcessWaiter) recordInvocation(key string, args []interface{}) {
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

var _ processwaiter.ProcessWaiter = new(FakeProcessWaiter).Spy
