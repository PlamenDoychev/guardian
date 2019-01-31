// Code generated by counterfeiter. DO NOT EDIT.
package stopperfakes

import (
	"sync"
	"syscall"

	"code.cloudfoundry.org/guardian/rundmc/stopper"
)

type FakeKiller struct {
	KillStub        func(syscall.Signal, ...int)
	killMutex       sync.RWMutex
	killArgsForCall []struct {
		arg1 syscall.Signal
		arg2 []int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKiller) Kill(arg1 syscall.Signal, arg2 ...int) {
	fake.killMutex.Lock()
	fake.killArgsForCall = append(fake.killArgsForCall, struct {
		arg1 syscall.Signal
		arg2 []int
	}{arg1, arg2})
	fake.recordInvocation("Kill", []interface{}{arg1, arg2})
	fake.killMutex.Unlock()
	if fake.KillStub != nil {
		fake.KillStub(arg1, arg2...)
	}
}

func (fake *FakeKiller) KillCallCount() int {
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	return len(fake.killArgsForCall)
}

func (fake *FakeKiller) KillCalls(stub func(syscall.Signal, ...int)) {
	fake.killMutex.Lock()
	defer fake.killMutex.Unlock()
	fake.KillStub = stub
}

func (fake *FakeKiller) KillArgsForCall(i int) (syscall.Signal, []int) {
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	argsForCall := fake.killArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeKiller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeKiller) recordInvocation(key string, args []interface{}) {
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

var _ stopper.Killer = new(FakeKiller)
