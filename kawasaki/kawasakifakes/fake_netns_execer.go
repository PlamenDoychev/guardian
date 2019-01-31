// Code generated by counterfeiter. DO NOT EDIT.
package kawasakifakes

import (
	"os"
	"sync"

	"code.cloudfoundry.org/guardian/kawasaki"
)

type FakeNetnsExecer struct {
	ExecStub        func(*os.File, func() error) error
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		arg1 *os.File
		arg2 func() error
	}
	execReturns struct {
		result1 error
	}
	execReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNetnsExecer) Exec(arg1 *os.File, arg2 func() error) error {
	fake.execMutex.Lock()
	ret, specificReturn := fake.execReturnsOnCall[len(fake.execArgsForCall)]
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		arg1 *os.File
		arg2 func() error
	}{arg1, arg2})
	fake.recordInvocation("Exec", []interface{}{arg1, arg2})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.execReturns
	return fakeReturns.result1
}

func (fake *FakeNetnsExecer) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeNetnsExecer) ExecCalls(stub func(*os.File, func() error) error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = stub
}

func (fake *FakeNetnsExecer) ExecArgsForCall(i int) (*os.File, func() error) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	argsForCall := fake.execArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNetnsExecer) ExecReturns(result1 error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetnsExecer) ExecReturnsOnCall(i int, result1 error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	if fake.execReturnsOnCall == nil {
		fake.execReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.execReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetnsExecer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNetnsExecer) recordInvocation(key string, args []interface{}) {
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

var _ kawasaki.NetnsExecer = new(FakeNetnsExecer)
