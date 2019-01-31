// Code generated by counterfeiter. DO NOT EDIT.
package runcontainerdfakes

import (
	"io"
	"sync"

	"code.cloudfoundry.org/guardian/rundmc/event"
	"code.cloudfoundry.org/guardian/rundmc/runcontainerd"
	"code.cloudfoundry.org/lager"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type FakeContainerManager struct {
	CreateStub        func(lager.Logger, string, *specs.Spec, func() (io.Reader, io.Writer, io.Writer)) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 *specs.Spec
		arg4 func() (io.Reader, io.Writer, io.Writer)
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(lager.Logger, string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	EventsStub        func(lager.Logger) <-chan event.Event
	eventsMutex       sync.RWMutex
	eventsArgsForCall []struct {
		arg1 lager.Logger
	}
	eventsReturns struct {
		result1 <-chan event.Event
	}
	eventsReturnsOnCall map[int]struct {
		result1 <-chan event.Event
	}
	ExecStub        func(lager.Logger, string, string, *specs.Process, func() (io.Reader, io.Writer, io.Writer)) error
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 *specs.Process
		arg5 func() (io.Reader, io.Writer, io.Writer)
	}
	execReturns struct {
		result1 error
	}
	execReturnsOnCall map[int]struct {
		result1 error
	}
	GetContainerPIDStub        func(lager.Logger, string) (uint32, error)
	getContainerPIDMutex       sync.RWMutex
	getContainerPIDArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	getContainerPIDReturns struct {
		result1 uint32
		result2 error
	}
	getContainerPIDReturnsOnCall map[int]struct {
		result1 uint32
		result2 error
	}
	StateStub        func(lager.Logger, string) (int, string, error)
	stateMutex       sync.RWMutex
	stateArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	stateReturns struct {
		result1 int
		result2 string
		result3 error
	}
	stateReturnsOnCall map[int]struct {
		result1 int
		result2 string
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerManager) Create(arg1 lager.Logger, arg2 string, arg3 *specs.Spec, arg4 func() (io.Reader, io.Writer, io.Writer)) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 *specs.Spec
		arg4 func() (io.Reader, io.Writer, io.Writer)
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3, arg4})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1
}

func (fake *FakeContainerManager) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeContainerManager) CreateCalls(stub func(lager.Logger, string, *specs.Spec, func() (io.Reader, io.Writer, io.Writer)) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeContainerManager) CreateArgsForCall(i int) (lager.Logger, string, *specs.Spec, func() (io.Reader, io.Writer, io.Writer)) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeContainerManager) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerManager) CreateReturnsOnCall(i int, result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerManager) Delete(arg1 lager.Logger, arg2 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakeContainerManager) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeContainerManager) DeleteCalls(stub func(lager.Logger, string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeContainerManager) DeleteArgsForCall(i int) (lager.Logger, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeContainerManager) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerManager) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerManager) Events(arg1 lager.Logger) <-chan event.Event {
	fake.eventsMutex.Lock()
	ret, specificReturn := fake.eventsReturnsOnCall[len(fake.eventsArgsForCall)]
	fake.eventsArgsForCall = append(fake.eventsArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Events", []interface{}{arg1})
	fake.eventsMutex.Unlock()
	if fake.EventsStub != nil {
		return fake.EventsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.eventsReturns
	return fakeReturns.result1
}

func (fake *FakeContainerManager) EventsCallCount() int {
	fake.eventsMutex.RLock()
	defer fake.eventsMutex.RUnlock()
	return len(fake.eventsArgsForCall)
}

func (fake *FakeContainerManager) EventsCalls(stub func(lager.Logger) <-chan event.Event) {
	fake.eventsMutex.Lock()
	defer fake.eventsMutex.Unlock()
	fake.EventsStub = stub
}

func (fake *FakeContainerManager) EventsArgsForCall(i int) lager.Logger {
	fake.eventsMutex.RLock()
	defer fake.eventsMutex.RUnlock()
	argsForCall := fake.eventsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeContainerManager) EventsReturns(result1 <-chan event.Event) {
	fake.eventsMutex.Lock()
	defer fake.eventsMutex.Unlock()
	fake.EventsStub = nil
	fake.eventsReturns = struct {
		result1 <-chan event.Event
	}{result1}
}

func (fake *FakeContainerManager) EventsReturnsOnCall(i int, result1 <-chan event.Event) {
	fake.eventsMutex.Lock()
	defer fake.eventsMutex.Unlock()
	fake.EventsStub = nil
	if fake.eventsReturnsOnCall == nil {
		fake.eventsReturnsOnCall = make(map[int]struct {
			result1 <-chan event.Event
		})
	}
	fake.eventsReturnsOnCall[i] = struct {
		result1 <-chan event.Event
	}{result1}
}

func (fake *FakeContainerManager) Exec(arg1 lager.Logger, arg2 string, arg3 string, arg4 *specs.Process, arg5 func() (io.Reader, io.Writer, io.Writer)) error {
	fake.execMutex.Lock()
	ret, specificReturn := fake.execReturnsOnCall[len(fake.execArgsForCall)]
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 *specs.Process
		arg5 func() (io.Reader, io.Writer, io.Writer)
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("Exec", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.execReturns
	return fakeReturns.result1
}

func (fake *FakeContainerManager) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeContainerManager) ExecCalls(stub func(lager.Logger, string, string, *specs.Process, func() (io.Reader, io.Writer, io.Writer)) error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = stub
}

func (fake *FakeContainerManager) ExecArgsForCall(i int) (lager.Logger, string, string, *specs.Process, func() (io.Reader, io.Writer, io.Writer)) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	argsForCall := fake.execArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeContainerManager) ExecReturns(result1 error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerManager) ExecReturnsOnCall(i int, result1 error) {
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

func (fake *FakeContainerManager) GetContainerPID(arg1 lager.Logger, arg2 string) (uint32, error) {
	fake.getContainerPIDMutex.Lock()
	ret, specificReturn := fake.getContainerPIDReturnsOnCall[len(fake.getContainerPIDArgsForCall)]
	fake.getContainerPIDArgsForCall = append(fake.getContainerPIDArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetContainerPID", []interface{}{arg1, arg2})
	fake.getContainerPIDMutex.Unlock()
	if fake.GetContainerPIDStub != nil {
		return fake.GetContainerPIDStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getContainerPIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeContainerManager) GetContainerPIDCallCount() int {
	fake.getContainerPIDMutex.RLock()
	defer fake.getContainerPIDMutex.RUnlock()
	return len(fake.getContainerPIDArgsForCall)
}

func (fake *FakeContainerManager) GetContainerPIDCalls(stub func(lager.Logger, string) (uint32, error)) {
	fake.getContainerPIDMutex.Lock()
	defer fake.getContainerPIDMutex.Unlock()
	fake.GetContainerPIDStub = stub
}

func (fake *FakeContainerManager) GetContainerPIDArgsForCall(i int) (lager.Logger, string) {
	fake.getContainerPIDMutex.RLock()
	defer fake.getContainerPIDMutex.RUnlock()
	argsForCall := fake.getContainerPIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeContainerManager) GetContainerPIDReturns(result1 uint32, result2 error) {
	fake.getContainerPIDMutex.Lock()
	defer fake.getContainerPIDMutex.Unlock()
	fake.GetContainerPIDStub = nil
	fake.getContainerPIDReturns = struct {
		result1 uint32
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerManager) GetContainerPIDReturnsOnCall(i int, result1 uint32, result2 error) {
	fake.getContainerPIDMutex.Lock()
	defer fake.getContainerPIDMutex.Unlock()
	fake.GetContainerPIDStub = nil
	if fake.getContainerPIDReturnsOnCall == nil {
		fake.getContainerPIDReturnsOnCall = make(map[int]struct {
			result1 uint32
			result2 error
		})
	}
	fake.getContainerPIDReturnsOnCall[i] = struct {
		result1 uint32
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerManager) State(arg1 lager.Logger, arg2 string) (int, string, error) {
	fake.stateMutex.Lock()
	ret, specificReturn := fake.stateReturnsOnCall[len(fake.stateArgsForCall)]
	fake.stateArgsForCall = append(fake.stateArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("State", []interface{}{arg1, arg2})
	fake.stateMutex.Unlock()
	if fake.StateStub != nil {
		return fake.StateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.stateReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeContainerManager) StateCallCount() int {
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	return len(fake.stateArgsForCall)
}

func (fake *FakeContainerManager) StateCalls(stub func(lager.Logger, string) (int, string, error)) {
	fake.stateMutex.Lock()
	defer fake.stateMutex.Unlock()
	fake.StateStub = stub
}

func (fake *FakeContainerManager) StateArgsForCall(i int) (lager.Logger, string) {
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	argsForCall := fake.stateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeContainerManager) StateReturns(result1 int, result2 string, result3 error) {
	fake.stateMutex.Lock()
	defer fake.stateMutex.Unlock()
	fake.StateStub = nil
	fake.stateReturns = struct {
		result1 int
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeContainerManager) StateReturnsOnCall(i int, result1 int, result2 string, result3 error) {
	fake.stateMutex.Lock()
	defer fake.stateMutex.Unlock()
	fake.StateStub = nil
	if fake.stateReturnsOnCall == nil {
		fake.stateReturnsOnCall = make(map[int]struct {
			result1 int
			result2 string
			result3 error
		})
	}
	fake.stateReturnsOnCall[i] = struct {
		result1 int
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeContainerManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.eventsMutex.RLock()
	defer fake.eventsMutex.RUnlock()
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	fake.getContainerPIDMutex.RLock()
	defer fake.getContainerPIDMutex.RUnlock()
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContainerManager) recordInvocation(key string, args []interface{}) {
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

var _ runcontainerd.ContainerManager = new(FakeContainerManager)
