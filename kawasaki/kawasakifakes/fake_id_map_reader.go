// This file was generated by counterfeiter
package kawasakifakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/kawasaki"
)

type FakeIdMapReader struct {
	ReadRootIdStub        func(path string) (int, error)
	readRootIdMutex       sync.RWMutex
	readRootIdArgsForCall []struct {
		path string
	}
	readRootIdReturns struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIdMapReader) ReadRootId(path string) (int, error) {
	fake.readRootIdMutex.Lock()
	fake.readRootIdArgsForCall = append(fake.readRootIdArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("ReadRootId", []interface{}{path})
	fake.readRootIdMutex.Unlock()
	if fake.ReadRootIdStub != nil {
		return fake.ReadRootIdStub(path)
	}
	return fake.readRootIdReturns.result1, fake.readRootIdReturns.result2
}

func (fake *FakeIdMapReader) ReadRootIdCallCount() int {
	fake.readRootIdMutex.RLock()
	defer fake.readRootIdMutex.RUnlock()
	return len(fake.readRootIdArgsForCall)
}

func (fake *FakeIdMapReader) ReadRootIdArgsForCall(i int) string {
	fake.readRootIdMutex.RLock()
	defer fake.readRootIdMutex.RUnlock()
	return fake.readRootIdArgsForCall[i].path
}

func (fake *FakeIdMapReader) ReadRootIdReturns(result1 int, result2 error) {
	fake.ReadRootIdStub = nil
	fake.readRootIdReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeIdMapReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readRootIdMutex.RLock()
	defer fake.readRootIdMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeIdMapReader) recordInvocation(key string, args []interface{}) {
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

var _ kawasaki.IdMapReader = new(FakeIdMapReader)
