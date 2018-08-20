// Code generated by counterfeiter. DO NOT EDIT.
package runruncfakes

import (
	"sync"

	"code.cloudfoundry.org/garden"
	"github.com/concourse/guardian/rundmc/runrunc"
)

type FakeStatsNotifier struct {
	OnStatStub        func(handle string, cpuStat garden.ContainerCPUStat, memoryStat garden.ContainerMemoryStat)
	onStatMutex       sync.RWMutex
	onStatArgsForCall []struct {
		handle     string
		cpuStat    garden.ContainerCPUStat
		memoryStat garden.ContainerMemoryStat
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStatsNotifier) OnStat(handle string, cpuStat garden.ContainerCPUStat, memoryStat garden.ContainerMemoryStat) {
	fake.onStatMutex.Lock()
	fake.onStatArgsForCall = append(fake.onStatArgsForCall, struct {
		handle     string
		cpuStat    garden.ContainerCPUStat
		memoryStat garden.ContainerMemoryStat
	}{handle, cpuStat, memoryStat})
	fake.recordInvocation("OnStat", []interface{}{handle, cpuStat, memoryStat})
	fake.onStatMutex.Unlock()
	if fake.OnStatStub != nil {
		fake.OnStatStub(handle, cpuStat, memoryStat)
	}
}

func (fake *FakeStatsNotifier) OnStatCallCount() int {
	fake.onStatMutex.RLock()
	defer fake.onStatMutex.RUnlock()
	return len(fake.onStatArgsForCall)
}

func (fake *FakeStatsNotifier) OnStatArgsForCall(i int) (string, garden.ContainerCPUStat, garden.ContainerMemoryStat) {
	fake.onStatMutex.RLock()
	defer fake.onStatMutex.RUnlock()
	return fake.onStatArgsForCall[i].handle, fake.onStatArgsForCall[i].cpuStat, fake.onStatArgsForCall[i].memoryStat
}

func (fake *FakeStatsNotifier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.onStatMutex.RLock()
	defer fake.onStatMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStatsNotifier) recordInvocation(key string, args []interface{}) {
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

var _ runrunc.StatsNotifier = new(FakeStatsNotifier)
