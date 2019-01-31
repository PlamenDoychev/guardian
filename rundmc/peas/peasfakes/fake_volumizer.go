// Code generated by counterfeiter. DO NOT EDIT.
package peasfakes

import (
	"sync"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/rundmc/peas"
	"code.cloudfoundry.org/lager"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type FakeVolumizer struct {
	CreateStub        func(lager.Logger, garden.ContainerSpec) (specs.Spec, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 lager.Logger
		arg2 garden.ContainerSpec
	}
	createReturns struct {
		result1 specs.Spec
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 specs.Spec
		result2 error
	}
	DestroyStub        func(lager.Logger, string) error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	destroyReturns struct {
		result1 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolumizer) Create(arg1 lager.Logger, arg2 garden.ContainerSpec) (specs.Spec, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 lager.Logger
		arg2 garden.ContainerSpec
	}{arg1, arg2})
	fake.recordInvocation("Create", []interface{}{arg1, arg2})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolumizer) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeVolumizer) CreateCalls(stub func(lager.Logger, garden.ContainerSpec) (specs.Spec, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeVolumizer) CreateArgsForCall(i int) (lager.Logger, garden.ContainerSpec) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVolumizer) CreateReturns(result1 specs.Spec, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 specs.Spec
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumizer) CreateReturnsOnCall(i int, result1 specs.Spec, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 specs.Spec
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 specs.Spec
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumizer) Destroy(arg1 lager.Logger, arg2 string) error {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Destroy", []interface{}{arg1, arg2})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.destroyReturns
	return fakeReturns.result1
}

func (fake *FakeVolumizer) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeVolumizer) DestroyCalls(stub func(lager.Logger, string) error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = stub
}

func (fake *FakeVolumizer) DestroyArgsForCall(i int) (lager.Logger, string) {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	argsForCall := fake.destroyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVolumizer) DestroyReturns(result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolumizer) DestroyReturnsOnCall(i int, result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolumizer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVolumizer) recordInvocation(key string, args []interface{}) {
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

var _ peas.Volumizer = new(FakeVolumizer)
