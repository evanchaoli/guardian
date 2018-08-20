// Code generated by counterfeiter. DO NOT EDIT.
package rundmcfakes

import (
	"io"
	"sync"

	"github.com/concourse/guardian/rundmc"
	"code.cloudfoundry.org/lager"
)

type FakeNstarRunner struct {
	StreamInStub        func(log lager.Logger, pid int, path string, user string, tarStream io.Reader) error
	streamInMutex       sync.RWMutex
	streamInArgsForCall []struct {
		log       lager.Logger
		pid       int
		path      string
		user      string
		tarStream io.Reader
	}
	streamInReturns struct {
		result1 error
	}
	streamInReturnsOnCall map[int]struct {
		result1 error
	}
	StreamOutStub        func(log lager.Logger, pid int, path string, user string) (io.ReadCloser, error)
	streamOutMutex       sync.RWMutex
	streamOutArgsForCall []struct {
		log  lager.Logger
		pid  int
		path string
		user string
	}
	streamOutReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	streamOutReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNstarRunner) StreamIn(log lager.Logger, pid int, path string, user string, tarStream io.Reader) error {
	fake.streamInMutex.Lock()
	ret, specificReturn := fake.streamInReturnsOnCall[len(fake.streamInArgsForCall)]
	fake.streamInArgsForCall = append(fake.streamInArgsForCall, struct {
		log       lager.Logger
		pid       int
		path      string
		user      string
		tarStream io.Reader
	}{log, pid, path, user, tarStream})
	fake.recordInvocation("StreamIn", []interface{}{log, pid, path, user, tarStream})
	fake.streamInMutex.Unlock()
	if fake.StreamInStub != nil {
		return fake.StreamInStub(log, pid, path, user, tarStream)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.streamInReturns.result1
}

func (fake *FakeNstarRunner) StreamInCallCount() int {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return len(fake.streamInArgsForCall)
}

func (fake *FakeNstarRunner) StreamInArgsForCall(i int) (lager.Logger, int, string, string, io.Reader) {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return fake.streamInArgsForCall[i].log, fake.streamInArgsForCall[i].pid, fake.streamInArgsForCall[i].path, fake.streamInArgsForCall[i].user, fake.streamInArgsForCall[i].tarStream
}

func (fake *FakeNstarRunner) StreamInReturns(result1 error) {
	fake.StreamInStub = nil
	fake.streamInReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNstarRunner) StreamInReturnsOnCall(i int, result1 error) {
	fake.StreamInStub = nil
	if fake.streamInReturnsOnCall == nil {
		fake.streamInReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.streamInReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNstarRunner) StreamOut(log lager.Logger, pid int, path string, user string) (io.ReadCloser, error) {
	fake.streamOutMutex.Lock()
	ret, specificReturn := fake.streamOutReturnsOnCall[len(fake.streamOutArgsForCall)]
	fake.streamOutArgsForCall = append(fake.streamOutArgsForCall, struct {
		log  lager.Logger
		pid  int
		path string
		user string
	}{log, pid, path, user})
	fake.recordInvocation("StreamOut", []interface{}{log, pid, path, user})
	fake.streamOutMutex.Unlock()
	if fake.StreamOutStub != nil {
		return fake.StreamOutStub(log, pid, path, user)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.streamOutReturns.result1, fake.streamOutReturns.result2
}

func (fake *FakeNstarRunner) StreamOutCallCount() int {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return len(fake.streamOutArgsForCall)
}

func (fake *FakeNstarRunner) StreamOutArgsForCall(i int) (lager.Logger, int, string, string) {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return fake.streamOutArgsForCall[i].log, fake.streamOutArgsForCall[i].pid, fake.streamOutArgsForCall[i].path, fake.streamOutArgsForCall[i].user
}

func (fake *FakeNstarRunner) StreamOutReturns(result1 io.ReadCloser, result2 error) {
	fake.StreamOutStub = nil
	fake.streamOutReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeNstarRunner) StreamOutReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.StreamOutStub = nil
	if fake.streamOutReturnsOnCall == nil {
		fake.streamOutReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.streamOutReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeNstarRunner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNstarRunner) recordInvocation(key string, args []interface{}) {
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

var _ rundmc.NstarRunner = new(FakeNstarRunner)
