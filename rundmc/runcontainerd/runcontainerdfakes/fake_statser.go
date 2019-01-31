// Code generated by counterfeiter. DO NOT EDIT.
package runcontainerdfakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/gardener"
	"code.cloudfoundry.org/guardian/rundmc/runcontainerd"
	"code.cloudfoundry.org/lager"
)

type FakeStatser struct {
	StatsStub        func(lager.Logger, string) (gardener.StatsContainerMetrics, error)
	statsMutex       sync.RWMutex
	statsArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	statsReturns struct {
		result1 gardener.StatsContainerMetrics
		result2 error
	}
	statsReturnsOnCall map[int]struct {
		result1 gardener.StatsContainerMetrics
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStatser) Stats(arg1 lager.Logger, arg2 string) (gardener.StatsContainerMetrics, error) {
	fake.statsMutex.Lock()
	ret, specificReturn := fake.statsReturnsOnCall[len(fake.statsArgsForCall)]
	fake.statsArgsForCall = append(fake.statsArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Stats", []interface{}{arg1, arg2})
	fake.statsMutex.Unlock()
	if fake.StatsStub != nil {
		return fake.StatsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.statsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStatser) StatsCallCount() int {
	fake.statsMutex.RLock()
	defer fake.statsMutex.RUnlock()
	return len(fake.statsArgsForCall)
}

func (fake *FakeStatser) StatsCalls(stub func(lager.Logger, string) (gardener.StatsContainerMetrics, error)) {
	fake.statsMutex.Lock()
	defer fake.statsMutex.Unlock()
	fake.StatsStub = stub
}

func (fake *FakeStatser) StatsArgsForCall(i int) (lager.Logger, string) {
	fake.statsMutex.RLock()
	defer fake.statsMutex.RUnlock()
	argsForCall := fake.statsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStatser) StatsReturns(result1 gardener.StatsContainerMetrics, result2 error) {
	fake.statsMutex.Lock()
	defer fake.statsMutex.Unlock()
	fake.StatsStub = nil
	fake.statsReturns = struct {
		result1 gardener.StatsContainerMetrics
		result2 error
	}{result1, result2}
}

func (fake *FakeStatser) StatsReturnsOnCall(i int, result1 gardener.StatsContainerMetrics, result2 error) {
	fake.statsMutex.Lock()
	defer fake.statsMutex.Unlock()
	fake.StatsStub = nil
	if fake.statsReturnsOnCall == nil {
		fake.statsReturnsOnCall = make(map[int]struct {
			result1 gardener.StatsContainerMetrics
			result2 error
		})
	}
	fake.statsReturnsOnCall[i] = struct {
		result1 gardener.StatsContainerMetrics
		result2 error
	}{result1, result2}
}

func (fake *FakeStatser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.statsMutex.RLock()
	defer fake.statsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStatser) recordInvocation(key string, args []interface{}) {
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

var _ runcontainerd.Statser = new(FakeStatser)
