// Code generated by counterfeiter. DO NOT EDIT.
package signalsfakes

import (
	"sync"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/rundmc/signals"
)

type FakeSignaller struct {
	SignalStub        func(garden.Signal) error
	signalMutex       sync.RWMutex
	signalArgsForCall []struct {
		arg1 garden.Signal
	}
	signalReturns struct {
		result1 error
	}
	signalReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSignaller) Signal(arg1 garden.Signal) error {
	fake.signalMutex.Lock()
	ret, specificReturn := fake.signalReturnsOnCall[len(fake.signalArgsForCall)]
	fake.signalArgsForCall = append(fake.signalArgsForCall, struct {
		arg1 garden.Signal
	}{arg1})
	fake.recordInvocation("Signal", []interface{}{arg1})
	fake.signalMutex.Unlock()
	if fake.SignalStub != nil {
		return fake.SignalStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.signalReturns
	return fakeReturns.result1
}

func (fake *FakeSignaller) SignalCallCount() int {
	fake.signalMutex.RLock()
	defer fake.signalMutex.RUnlock()
	return len(fake.signalArgsForCall)
}

func (fake *FakeSignaller) SignalCalls(stub func(garden.Signal) error) {
	fake.signalMutex.Lock()
	defer fake.signalMutex.Unlock()
	fake.SignalStub = stub
}

func (fake *FakeSignaller) SignalArgsForCall(i int) garden.Signal {
	fake.signalMutex.RLock()
	defer fake.signalMutex.RUnlock()
	argsForCall := fake.signalArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSignaller) SignalReturns(result1 error) {
	fake.signalMutex.Lock()
	defer fake.signalMutex.Unlock()
	fake.SignalStub = nil
	fake.signalReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSignaller) SignalReturnsOnCall(i int, result1 error) {
	fake.signalMutex.Lock()
	defer fake.signalMutex.Unlock()
	fake.SignalStub = nil
	if fake.signalReturnsOnCall == nil {
		fake.signalReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.signalReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSignaller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.signalMutex.RLock()
	defer fake.signalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSignaller) recordInvocation(key string, args []interface{}) {
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

var _ signals.Signaller = new(FakeSignaller)
