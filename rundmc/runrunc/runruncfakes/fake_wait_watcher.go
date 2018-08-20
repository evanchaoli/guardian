// Code generated by counterfeiter. DO NOT EDIT.
package runruncfakes

import (
	"sync"

	"github.com/concourse/guardian/rundmc/runrunc"
	"code.cloudfoundry.org/lager"
)

type FakeWaitWatcher struct {
	OnExitStub        func(log lager.Logger, process runrunc.Waiter, onExit runrunc.Runner)
	onExitMutex       sync.RWMutex
	onExitArgsForCall []struct {
		log     lager.Logger
		process runrunc.Waiter
		onExit  runrunc.Runner
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWaitWatcher) OnExit(log lager.Logger, process runrunc.Waiter, onExit runrunc.Runner) {
	fake.onExitMutex.Lock()
	fake.onExitArgsForCall = append(fake.onExitArgsForCall, struct {
		log     lager.Logger
		process runrunc.Waiter
		onExit  runrunc.Runner
	}{log, process, onExit})
	fake.recordInvocation("OnExit", []interface{}{log, process, onExit})
	fake.onExitMutex.Unlock()
	if fake.OnExitStub != nil {
		fake.OnExitStub(log, process, onExit)
	}
}

func (fake *FakeWaitWatcher) OnExitCallCount() int {
	fake.onExitMutex.RLock()
	defer fake.onExitMutex.RUnlock()
	return len(fake.onExitArgsForCall)
}

func (fake *FakeWaitWatcher) OnExitArgsForCall(i int) (lager.Logger, runrunc.Waiter, runrunc.Runner) {
	fake.onExitMutex.RLock()
	defer fake.onExitMutex.RUnlock()
	return fake.onExitArgsForCall[i].log, fake.onExitArgsForCall[i].process, fake.onExitArgsForCall[i].onExit
}

func (fake *FakeWaitWatcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.onExitMutex.RLock()
	defer fake.onExitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWaitWatcher) recordInvocation(key string, args []interface{}) {
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

var _ runrunc.WaitWatcher = new(FakeWaitWatcher)
