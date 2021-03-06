// This file was generated by counterfeiter
package workerfakes

import (
	"sync"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/worker"
)

type FakeWorkerProvider struct {
	WorkersStub        func() ([]worker.Worker, error)
	workersMutex       sync.RWMutex
	workersArgsForCall []struct{}
	workersReturns     struct {
		result1 []worker.Worker
		result2 error
	}
	GetWorkerStub        func(string) (worker.Worker, bool, error)
	getWorkerMutex       sync.RWMutex
	getWorkerArgsForCall []struct {
		arg1 string
	}
	getWorkerReturns struct {
		result1 worker.Worker
		result2 bool
		result3 error
	}
	FindContainerForIdentifierStub        func(worker.Identifier) (db.SavedContainer, bool, error)
	findContainerForIdentifierMutex       sync.RWMutex
	findContainerForIdentifierArgsForCall []struct {
		arg1 worker.Identifier
	}
	findContainerForIdentifierReturns struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}
	GetContainerStub        func(string) (db.SavedContainer, bool, error)
	getContainerMutex       sync.RWMutex
	getContainerArgsForCall []struct {
		arg1 string
	}
	getContainerReturns struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}
	ReapContainerStub        func(string) error
	reapContainerMutex       sync.RWMutex
	reapContainerArgsForCall []struct {
		arg1 string
	}
	reapContainerReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorkerProvider) Workers() ([]worker.Worker, error) {
	fake.workersMutex.Lock()
	fake.workersArgsForCall = append(fake.workersArgsForCall, struct{}{})
	fake.recordInvocation("Workers", []interface{}{})
	fake.workersMutex.Unlock()
	if fake.WorkersStub != nil {
		return fake.WorkersStub()
	} else {
		return fake.workersReturns.result1, fake.workersReturns.result2
	}
}

func (fake *FakeWorkerProvider) WorkersCallCount() int {
	fake.workersMutex.RLock()
	defer fake.workersMutex.RUnlock()
	return len(fake.workersArgsForCall)
}

func (fake *FakeWorkerProvider) WorkersReturns(result1 []worker.Worker, result2 error) {
	fake.WorkersStub = nil
	fake.workersReturns = struct {
		result1 []worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerProvider) GetWorker(arg1 string) (worker.Worker, bool, error) {
	fake.getWorkerMutex.Lock()
	fake.getWorkerArgsForCall = append(fake.getWorkerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetWorker", []interface{}{arg1})
	fake.getWorkerMutex.Unlock()
	if fake.GetWorkerStub != nil {
		return fake.GetWorkerStub(arg1)
	} else {
		return fake.getWorkerReturns.result1, fake.getWorkerReturns.result2, fake.getWorkerReturns.result3
	}
}

func (fake *FakeWorkerProvider) GetWorkerCallCount() int {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return len(fake.getWorkerArgsForCall)
}

func (fake *FakeWorkerProvider) GetWorkerArgsForCall(i int) string {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return fake.getWorkerArgsForCall[i].arg1
}

func (fake *FakeWorkerProvider) GetWorkerReturns(result1 worker.Worker, result2 bool, result3 error) {
	fake.GetWorkerStub = nil
	fake.getWorkerReturns = struct {
		result1 worker.Worker
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerProvider) FindContainerForIdentifier(arg1 worker.Identifier) (db.SavedContainer, bool, error) {
	fake.findContainerForIdentifierMutex.Lock()
	fake.findContainerForIdentifierArgsForCall = append(fake.findContainerForIdentifierArgsForCall, struct {
		arg1 worker.Identifier
	}{arg1})
	fake.recordInvocation("FindContainerForIdentifier", []interface{}{arg1})
	fake.findContainerForIdentifierMutex.Unlock()
	if fake.FindContainerForIdentifierStub != nil {
		return fake.FindContainerForIdentifierStub(arg1)
	} else {
		return fake.findContainerForIdentifierReturns.result1, fake.findContainerForIdentifierReturns.result2, fake.findContainerForIdentifierReturns.result3
	}
}

func (fake *FakeWorkerProvider) FindContainerForIdentifierCallCount() int {
	fake.findContainerForIdentifierMutex.RLock()
	defer fake.findContainerForIdentifierMutex.RUnlock()
	return len(fake.findContainerForIdentifierArgsForCall)
}

func (fake *FakeWorkerProvider) FindContainerForIdentifierArgsForCall(i int) worker.Identifier {
	fake.findContainerForIdentifierMutex.RLock()
	defer fake.findContainerForIdentifierMutex.RUnlock()
	return fake.findContainerForIdentifierArgsForCall[i].arg1
}

func (fake *FakeWorkerProvider) FindContainerForIdentifierReturns(result1 db.SavedContainer, result2 bool, result3 error) {
	fake.FindContainerForIdentifierStub = nil
	fake.findContainerForIdentifierReturns = struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerProvider) GetContainer(arg1 string) (db.SavedContainer, bool, error) {
	fake.getContainerMutex.Lock()
	fake.getContainerArgsForCall = append(fake.getContainerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetContainer", []interface{}{arg1})
	fake.getContainerMutex.Unlock()
	if fake.GetContainerStub != nil {
		return fake.GetContainerStub(arg1)
	} else {
		return fake.getContainerReturns.result1, fake.getContainerReturns.result2, fake.getContainerReturns.result3
	}
}

func (fake *FakeWorkerProvider) GetContainerCallCount() int {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return len(fake.getContainerArgsForCall)
}

func (fake *FakeWorkerProvider) GetContainerArgsForCall(i int) string {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return fake.getContainerArgsForCall[i].arg1
}

func (fake *FakeWorkerProvider) GetContainerReturns(result1 db.SavedContainer, result2 bool, result3 error) {
	fake.GetContainerStub = nil
	fake.getContainerReturns = struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerProvider) ReapContainer(arg1 string) error {
	fake.reapContainerMutex.Lock()
	fake.reapContainerArgsForCall = append(fake.reapContainerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ReapContainer", []interface{}{arg1})
	fake.reapContainerMutex.Unlock()
	if fake.ReapContainerStub != nil {
		return fake.ReapContainerStub(arg1)
	} else {
		return fake.reapContainerReturns.result1
	}
}

func (fake *FakeWorkerProvider) ReapContainerCallCount() int {
	fake.reapContainerMutex.RLock()
	defer fake.reapContainerMutex.RUnlock()
	return len(fake.reapContainerArgsForCall)
}

func (fake *FakeWorkerProvider) ReapContainerArgsForCall(i int) string {
	fake.reapContainerMutex.RLock()
	defer fake.reapContainerMutex.RUnlock()
	return fake.reapContainerArgsForCall[i].arg1
}

func (fake *FakeWorkerProvider) ReapContainerReturns(result1 error) {
	fake.ReapContainerStub = nil
	fake.reapContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.workersMutex.RLock()
	defer fake.workersMutex.RUnlock()
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	fake.findContainerForIdentifierMutex.RLock()
	defer fake.findContainerForIdentifierMutex.RUnlock()
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	fake.reapContainerMutex.RLock()
	defer fake.reapContainerMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeWorkerProvider) recordInvocation(key string, args []interface{}) {
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

var _ worker.WorkerProvider = new(FakeWorkerProvider)
