// Code generated by counterfeiter. DO NOT EDIT.
package runruncfakes

import (
	"os"
	"sync"

	"github.com/concourse/guardian/rundmc/runrunc"
)

type FakeMkdirer struct {
	MkdirAsStub        func(rootFSPathFile string, uid, gid int, mode os.FileMode, recreate bool, path ...string) error
	mkdirAsMutex       sync.RWMutex
	mkdirAsArgsForCall []struct {
		rootFSPathFile string
		uid            int
		gid            int
		mode           os.FileMode
		recreate       bool
		path           []string
	}
	mkdirAsReturns struct {
		result1 error
	}
	mkdirAsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMkdirer) MkdirAs(rootFSPathFile string, uid int, gid int, mode os.FileMode, recreate bool, path ...string) error {
	fake.mkdirAsMutex.Lock()
	ret, specificReturn := fake.mkdirAsReturnsOnCall[len(fake.mkdirAsArgsForCall)]
	fake.mkdirAsArgsForCall = append(fake.mkdirAsArgsForCall, struct {
		rootFSPathFile string
		uid            int
		gid            int
		mode           os.FileMode
		recreate       bool
		path           []string
	}{rootFSPathFile, uid, gid, mode, recreate, path})
	fake.recordInvocation("MkdirAs", []interface{}{rootFSPathFile, uid, gid, mode, recreate, path})
	fake.mkdirAsMutex.Unlock()
	if fake.MkdirAsStub != nil {
		return fake.MkdirAsStub(rootFSPathFile, uid, gid, mode, recreate, path...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.mkdirAsReturns.result1
}

func (fake *FakeMkdirer) MkdirAsCallCount() int {
	fake.mkdirAsMutex.RLock()
	defer fake.mkdirAsMutex.RUnlock()
	return len(fake.mkdirAsArgsForCall)
}

func (fake *FakeMkdirer) MkdirAsArgsForCall(i int) (string, int, int, os.FileMode, bool, []string) {
	fake.mkdirAsMutex.RLock()
	defer fake.mkdirAsMutex.RUnlock()
	return fake.mkdirAsArgsForCall[i].rootFSPathFile, fake.mkdirAsArgsForCall[i].uid, fake.mkdirAsArgsForCall[i].gid, fake.mkdirAsArgsForCall[i].mode, fake.mkdirAsArgsForCall[i].recreate, fake.mkdirAsArgsForCall[i].path
}

func (fake *FakeMkdirer) MkdirAsReturns(result1 error) {
	fake.MkdirAsStub = nil
	fake.mkdirAsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMkdirer) MkdirAsReturnsOnCall(i int, result1 error) {
	fake.MkdirAsStub = nil
	if fake.mkdirAsReturnsOnCall == nil {
		fake.mkdirAsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mkdirAsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMkdirer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mkdirAsMutex.RLock()
	defer fake.mkdirAsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMkdirer) recordInvocation(key string, args []interface{}) {
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

var _ runrunc.Mkdirer = new(FakeMkdirer)
