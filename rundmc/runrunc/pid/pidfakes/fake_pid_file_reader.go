// Code generated by counterfeiter. DO NOT EDIT.
package pidfakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/rundmc/runrunc/pid"
)

type FakePidFileReader struct {
	PidStub        func(string) (int, error)
	pidMutex       sync.RWMutex
	pidArgsForCall []struct {
		arg1 string
	}
	pidReturns struct {
		result1 int
		result2 error
	}
	pidReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePidFileReader) Pid(arg1 string) (int, error) {
	fake.pidMutex.Lock()
	ret, specificReturn := fake.pidReturnsOnCall[len(fake.pidArgsForCall)]
	fake.pidArgsForCall = append(fake.pidArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Pid", []interface{}{arg1})
	fake.pidMutex.Unlock()
	if fake.PidStub != nil {
		return fake.PidStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.pidReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePidFileReader) PidCallCount() int {
	fake.pidMutex.RLock()
	defer fake.pidMutex.RUnlock()
	return len(fake.pidArgsForCall)
}

func (fake *FakePidFileReader) PidCalls(stub func(string) (int, error)) {
	fake.pidMutex.Lock()
	defer fake.pidMutex.Unlock()
	fake.PidStub = stub
}

func (fake *FakePidFileReader) PidArgsForCall(i int) string {
	fake.pidMutex.RLock()
	defer fake.pidMutex.RUnlock()
	argsForCall := fake.pidArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePidFileReader) PidReturns(result1 int, result2 error) {
	fake.pidMutex.Lock()
	defer fake.pidMutex.Unlock()
	fake.PidStub = nil
	fake.pidReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakePidFileReader) PidReturnsOnCall(i int, result1 int, result2 error) {
	fake.pidMutex.Lock()
	defer fake.pidMutex.Unlock()
	fake.PidStub = nil
	if fake.pidReturnsOnCall == nil {
		fake.pidReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.pidReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakePidFileReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pidMutex.RLock()
	defer fake.pidMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePidFileReader) recordInvocation(key string, args []interface{}) {
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

var _ pid.PidFileReader = new(FakePidFileReader)
