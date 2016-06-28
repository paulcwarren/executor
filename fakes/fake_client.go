// This file was generated by counterfeiter
package fakes

import (
	"io"
	"sync"

	"code.cloudfoundry.org/executor"
	"github.com/pivotal-golang/lager"
)

type FakeClient struct {
	PingStub        func(logger lager.Logger) error
	pingMutex       sync.RWMutex
	pingArgsForCall []struct {
		logger lager.Logger
	}
	pingReturns struct {
		result1 error
	}
	AllocateContainersStub        func(logger lager.Logger, requests []executor.AllocationRequest) ([]executor.AllocationFailure, error)
	allocateContainersMutex       sync.RWMutex
	allocateContainersArgsForCall []struct {
		logger   lager.Logger
		requests []executor.AllocationRequest
	}
	allocateContainersReturns struct {
		result1 []executor.AllocationFailure
		result2 error
	}
	GetContainerStub        func(logger lager.Logger, guid string) (executor.Container, error)
	getContainerMutex       sync.RWMutex
	getContainerArgsForCall []struct {
		logger lager.Logger
		guid   string
	}
	getContainerReturns struct {
		result1 executor.Container
		result2 error
	}
	RunContainerStub        func(lager.Logger, *executor.RunRequest) error
	runContainerMutex       sync.RWMutex
	runContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 *executor.RunRequest
	}
	runContainerReturns struct {
		result1 error
	}
	StopContainerStub        func(logger lager.Logger, guid string) error
	stopContainerMutex       sync.RWMutex
	stopContainerArgsForCall []struct {
		logger lager.Logger
		guid   string
	}
	stopContainerReturns struct {
		result1 error
	}
	DeleteContainerStub        func(logger lager.Logger, guid string) error
	deleteContainerMutex       sync.RWMutex
	deleteContainerArgsForCall []struct {
		logger lager.Logger
		guid   string
	}
	deleteContainerReturns struct {
		result1 error
	}
	ListContainersStub        func(lager.Logger) ([]executor.Container, error)
	listContainersMutex       sync.RWMutex
	listContainersArgsForCall []struct {
		arg1 lager.Logger
	}
	listContainersReturns struct {
		result1 []executor.Container
		result2 error
	}
	GetBulkMetricsStub        func(lager.Logger) (map[string]executor.Metrics, error)
	getBulkMetricsMutex       sync.RWMutex
	getBulkMetricsArgsForCall []struct {
		arg1 lager.Logger
	}
	getBulkMetricsReturns struct {
		result1 map[string]executor.Metrics
		result2 error
	}
	RemainingResourcesStub        func(lager.Logger) (executor.ExecutorResources, error)
	remainingResourcesMutex       sync.RWMutex
	remainingResourcesArgsForCall []struct {
		arg1 lager.Logger
	}
	remainingResourcesReturns struct {
		result1 executor.ExecutorResources
		result2 error
	}
	TotalResourcesStub        func(lager.Logger) (executor.ExecutorResources, error)
	totalResourcesMutex       sync.RWMutex
	totalResourcesArgsForCall []struct {
		arg1 lager.Logger
	}
	totalResourcesReturns struct {
		result1 executor.ExecutorResources
		result2 error
	}
	GetFilesStub        func(logger lager.Logger, guid string, path string) (io.ReadCloser, error)
	getFilesMutex       sync.RWMutex
	getFilesArgsForCall []struct {
		logger lager.Logger
		guid   string
		path   string
	}
	getFilesReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	VolumeDriversStub        func(logger lager.Logger) ([]string, error)
	volumeDriversMutex       sync.RWMutex
	volumeDriversArgsForCall []struct {
		logger lager.Logger
	}
	volumeDriversReturns struct {
		result1 []string
		result2 error
	}
	SubscribeToEventsStub        func(lager.Logger) (executor.EventSource, error)
	subscribeToEventsMutex       sync.RWMutex
	subscribeToEventsArgsForCall []struct {
		arg1 lager.Logger
	}
	subscribeToEventsReturns struct {
		result1 executor.EventSource
		result2 error
	}
	HealthyStub        func(lager.Logger) bool
	healthyMutex       sync.RWMutex
	healthyArgsForCall []struct {
		arg1 lager.Logger
	}
	healthyReturns struct {
		result1 bool
	}
	SetHealthyStub        func(lager.Logger, bool)
	setHealthyMutex       sync.RWMutex
	setHealthyArgsForCall []struct {
		arg1 lager.Logger
		arg2 bool
	}
	CleanupStub        func(lager.Logger)
	cleanupMutex       sync.RWMutex
	cleanupArgsForCall []struct {
		arg1 lager.Logger
	}
}

func (fake *FakeClient) Ping(logger lager.Logger) error {
	fake.pingMutex.Lock()
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub(logger)
	} else {
		return fake.pingReturns.result1
	}
}

func (fake *FakeClient) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeClient) PingArgsForCall(i int) lager.Logger {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return fake.pingArgsForCall[i].logger
}

func (fake *FakeClient) PingReturns(result1 error) {
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) AllocateContainers(logger lager.Logger, requests []executor.AllocationRequest) ([]executor.AllocationFailure, error) {
	fake.allocateContainersMutex.Lock()
	fake.allocateContainersArgsForCall = append(fake.allocateContainersArgsForCall, struct {
		logger   lager.Logger
		requests []executor.AllocationRequest
	}{logger, requests})
	fake.allocateContainersMutex.Unlock()
	if fake.AllocateContainersStub != nil {
		return fake.AllocateContainersStub(logger, requests)
	} else {
		return fake.allocateContainersReturns.result1, fake.allocateContainersReturns.result2
	}
}

func (fake *FakeClient) AllocateContainersCallCount() int {
	fake.allocateContainersMutex.RLock()
	defer fake.allocateContainersMutex.RUnlock()
	return len(fake.allocateContainersArgsForCall)
}

func (fake *FakeClient) AllocateContainersArgsForCall(i int) (lager.Logger, []executor.AllocationRequest) {
	fake.allocateContainersMutex.RLock()
	defer fake.allocateContainersMutex.RUnlock()
	return fake.allocateContainersArgsForCall[i].logger, fake.allocateContainersArgsForCall[i].requests
}

func (fake *FakeClient) AllocateContainersReturns(result1 []executor.AllocationFailure, result2 error) {
	fake.AllocateContainersStub = nil
	fake.allocateContainersReturns = struct {
		result1 []executor.AllocationFailure
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetContainer(logger lager.Logger, guid string) (executor.Container, error) {
	fake.getContainerMutex.Lock()
	fake.getContainerArgsForCall = append(fake.getContainerArgsForCall, struct {
		logger lager.Logger
		guid   string
	}{logger, guid})
	fake.getContainerMutex.Unlock()
	if fake.GetContainerStub != nil {
		return fake.GetContainerStub(logger, guid)
	} else {
		return fake.getContainerReturns.result1, fake.getContainerReturns.result2
	}
}

func (fake *FakeClient) GetContainerCallCount() int {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return len(fake.getContainerArgsForCall)
}

func (fake *FakeClient) GetContainerArgsForCall(i int) (lager.Logger, string) {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return fake.getContainerArgsForCall[i].logger, fake.getContainerArgsForCall[i].guid
}

func (fake *FakeClient) GetContainerReturns(result1 executor.Container, result2 error) {
	fake.GetContainerStub = nil
	fake.getContainerReturns = struct {
		result1 executor.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) RunContainer(arg1 lager.Logger, arg2 *executor.RunRequest) error {
	fake.runContainerMutex.Lock()
	fake.runContainerArgsForCall = append(fake.runContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 *executor.RunRequest
	}{arg1, arg2})
	fake.runContainerMutex.Unlock()
	if fake.RunContainerStub != nil {
		return fake.RunContainerStub(arg1, arg2)
	} else {
		return fake.runContainerReturns.result1
	}
}

func (fake *FakeClient) RunContainerCallCount() int {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	return len(fake.runContainerArgsForCall)
}

func (fake *FakeClient) RunContainerArgsForCall(i int) (lager.Logger, *executor.RunRequest) {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	return fake.runContainerArgsForCall[i].arg1, fake.runContainerArgsForCall[i].arg2
}

func (fake *FakeClient) RunContainerReturns(result1 error) {
	fake.RunContainerStub = nil
	fake.runContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) StopContainer(logger lager.Logger, guid string) error {
	fake.stopContainerMutex.Lock()
	fake.stopContainerArgsForCall = append(fake.stopContainerArgsForCall, struct {
		logger lager.Logger
		guid   string
	}{logger, guid})
	fake.stopContainerMutex.Unlock()
	if fake.StopContainerStub != nil {
		return fake.StopContainerStub(logger, guid)
	} else {
		return fake.stopContainerReturns.result1
	}
}

func (fake *FakeClient) StopContainerCallCount() int {
	fake.stopContainerMutex.RLock()
	defer fake.stopContainerMutex.RUnlock()
	return len(fake.stopContainerArgsForCall)
}

func (fake *FakeClient) StopContainerArgsForCall(i int) (lager.Logger, string) {
	fake.stopContainerMutex.RLock()
	defer fake.stopContainerMutex.RUnlock()
	return fake.stopContainerArgsForCall[i].logger, fake.stopContainerArgsForCall[i].guid
}

func (fake *FakeClient) StopContainerReturns(result1 error) {
	fake.StopContainerStub = nil
	fake.stopContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DeleteContainer(logger lager.Logger, guid string) error {
	fake.deleteContainerMutex.Lock()
	fake.deleteContainerArgsForCall = append(fake.deleteContainerArgsForCall, struct {
		logger lager.Logger
		guid   string
	}{logger, guid})
	fake.deleteContainerMutex.Unlock()
	if fake.DeleteContainerStub != nil {
		return fake.DeleteContainerStub(logger, guid)
	} else {
		return fake.deleteContainerReturns.result1
	}
}

func (fake *FakeClient) DeleteContainerCallCount() int {
	fake.deleteContainerMutex.RLock()
	defer fake.deleteContainerMutex.RUnlock()
	return len(fake.deleteContainerArgsForCall)
}

func (fake *FakeClient) DeleteContainerArgsForCall(i int) (lager.Logger, string) {
	fake.deleteContainerMutex.RLock()
	defer fake.deleteContainerMutex.RUnlock()
	return fake.deleteContainerArgsForCall[i].logger, fake.deleteContainerArgsForCall[i].guid
}

func (fake *FakeClient) DeleteContainerReturns(result1 error) {
	fake.DeleteContainerStub = nil
	fake.deleteContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) ListContainers(arg1 lager.Logger) ([]executor.Container, error) {
	fake.listContainersMutex.Lock()
	fake.listContainersArgsForCall = append(fake.listContainersArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.listContainersMutex.Unlock()
	if fake.ListContainersStub != nil {
		return fake.ListContainersStub(arg1)
	} else {
		return fake.listContainersReturns.result1, fake.listContainersReturns.result2
	}
}

func (fake *FakeClient) ListContainersCallCount() int {
	fake.listContainersMutex.RLock()
	defer fake.listContainersMutex.RUnlock()
	return len(fake.listContainersArgsForCall)
}

func (fake *FakeClient) ListContainersArgsForCall(i int) lager.Logger {
	fake.listContainersMutex.RLock()
	defer fake.listContainersMutex.RUnlock()
	return fake.listContainersArgsForCall[i].arg1
}

func (fake *FakeClient) ListContainersReturns(result1 []executor.Container, result2 error) {
	fake.ListContainersStub = nil
	fake.listContainersReturns = struct {
		result1 []executor.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetBulkMetrics(arg1 lager.Logger) (map[string]executor.Metrics, error) {
	fake.getBulkMetricsMutex.Lock()
	fake.getBulkMetricsArgsForCall = append(fake.getBulkMetricsArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.getBulkMetricsMutex.Unlock()
	if fake.GetBulkMetricsStub != nil {
		return fake.GetBulkMetricsStub(arg1)
	} else {
		return fake.getBulkMetricsReturns.result1, fake.getBulkMetricsReturns.result2
	}
}

func (fake *FakeClient) GetBulkMetricsCallCount() int {
	fake.getBulkMetricsMutex.RLock()
	defer fake.getBulkMetricsMutex.RUnlock()
	return len(fake.getBulkMetricsArgsForCall)
}

func (fake *FakeClient) GetBulkMetricsArgsForCall(i int) lager.Logger {
	fake.getBulkMetricsMutex.RLock()
	defer fake.getBulkMetricsMutex.RUnlock()
	return fake.getBulkMetricsArgsForCall[i].arg1
}

func (fake *FakeClient) GetBulkMetricsReturns(result1 map[string]executor.Metrics, result2 error) {
	fake.GetBulkMetricsStub = nil
	fake.getBulkMetricsReturns = struct {
		result1 map[string]executor.Metrics
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) RemainingResources(arg1 lager.Logger) (executor.ExecutorResources, error) {
	fake.remainingResourcesMutex.Lock()
	fake.remainingResourcesArgsForCall = append(fake.remainingResourcesArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.remainingResourcesMutex.Unlock()
	if fake.RemainingResourcesStub != nil {
		return fake.RemainingResourcesStub(arg1)
	} else {
		return fake.remainingResourcesReturns.result1, fake.remainingResourcesReturns.result2
	}
}

func (fake *FakeClient) RemainingResourcesCallCount() int {
	fake.remainingResourcesMutex.RLock()
	defer fake.remainingResourcesMutex.RUnlock()
	return len(fake.remainingResourcesArgsForCall)
}

func (fake *FakeClient) RemainingResourcesArgsForCall(i int) lager.Logger {
	fake.remainingResourcesMutex.RLock()
	defer fake.remainingResourcesMutex.RUnlock()
	return fake.remainingResourcesArgsForCall[i].arg1
}

func (fake *FakeClient) RemainingResourcesReturns(result1 executor.ExecutorResources, result2 error) {
	fake.RemainingResourcesStub = nil
	fake.remainingResourcesReturns = struct {
		result1 executor.ExecutorResources
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) TotalResources(arg1 lager.Logger) (executor.ExecutorResources, error) {
	fake.totalResourcesMutex.Lock()
	fake.totalResourcesArgsForCall = append(fake.totalResourcesArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.totalResourcesMutex.Unlock()
	if fake.TotalResourcesStub != nil {
		return fake.TotalResourcesStub(arg1)
	} else {
		return fake.totalResourcesReturns.result1, fake.totalResourcesReturns.result2
	}
}

func (fake *FakeClient) TotalResourcesCallCount() int {
	fake.totalResourcesMutex.RLock()
	defer fake.totalResourcesMutex.RUnlock()
	return len(fake.totalResourcesArgsForCall)
}

func (fake *FakeClient) TotalResourcesArgsForCall(i int) lager.Logger {
	fake.totalResourcesMutex.RLock()
	defer fake.totalResourcesMutex.RUnlock()
	return fake.totalResourcesArgsForCall[i].arg1
}

func (fake *FakeClient) TotalResourcesReturns(result1 executor.ExecutorResources, result2 error) {
	fake.TotalResourcesStub = nil
	fake.totalResourcesReturns = struct {
		result1 executor.ExecutorResources
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetFiles(logger lager.Logger, guid string, path string) (io.ReadCloser, error) {
	fake.getFilesMutex.Lock()
	fake.getFilesArgsForCall = append(fake.getFilesArgsForCall, struct {
		logger lager.Logger
		guid   string
		path   string
	}{logger, guid, path})
	fake.getFilesMutex.Unlock()
	if fake.GetFilesStub != nil {
		return fake.GetFilesStub(logger, guid, path)
	} else {
		return fake.getFilesReturns.result1, fake.getFilesReturns.result2
	}
}

func (fake *FakeClient) GetFilesCallCount() int {
	fake.getFilesMutex.RLock()
	defer fake.getFilesMutex.RUnlock()
	return len(fake.getFilesArgsForCall)
}

func (fake *FakeClient) GetFilesArgsForCall(i int) (lager.Logger, string, string) {
	fake.getFilesMutex.RLock()
	defer fake.getFilesMutex.RUnlock()
	return fake.getFilesArgsForCall[i].logger, fake.getFilesArgsForCall[i].guid, fake.getFilesArgsForCall[i].path
}

func (fake *FakeClient) GetFilesReturns(result1 io.ReadCloser, result2 error) {
	fake.GetFilesStub = nil
	fake.getFilesReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) VolumeDrivers(logger lager.Logger) ([]string, error) {
	fake.volumeDriversMutex.Lock()
	fake.volumeDriversArgsForCall = append(fake.volumeDriversArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.volumeDriversMutex.Unlock()
	if fake.VolumeDriversStub != nil {
		return fake.VolumeDriversStub(logger)
	} else {
		return fake.volumeDriversReturns.result1, fake.volumeDriversReturns.result2
	}
}

func (fake *FakeClient) VolumeDriversCallCount() int {
	fake.volumeDriversMutex.RLock()
	defer fake.volumeDriversMutex.RUnlock()
	return len(fake.volumeDriversArgsForCall)
}

func (fake *FakeClient) VolumeDriversArgsForCall(i int) lager.Logger {
	fake.volumeDriversMutex.RLock()
	defer fake.volumeDriversMutex.RUnlock()
	return fake.volumeDriversArgsForCall[i].logger
}

func (fake *FakeClient) VolumeDriversReturns(result1 []string, result2 error) {
	fake.VolumeDriversStub = nil
	fake.volumeDriversReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) SubscribeToEvents(arg1 lager.Logger) (executor.EventSource, error) {
	fake.subscribeToEventsMutex.Lock()
	fake.subscribeToEventsArgsForCall = append(fake.subscribeToEventsArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.subscribeToEventsMutex.Unlock()
	if fake.SubscribeToEventsStub != nil {
		return fake.SubscribeToEventsStub(arg1)
	} else {
		return fake.subscribeToEventsReturns.result1, fake.subscribeToEventsReturns.result2
	}
}

func (fake *FakeClient) SubscribeToEventsCallCount() int {
	fake.subscribeToEventsMutex.RLock()
	defer fake.subscribeToEventsMutex.RUnlock()
	return len(fake.subscribeToEventsArgsForCall)
}

func (fake *FakeClient) SubscribeToEventsArgsForCall(i int) lager.Logger {
	fake.subscribeToEventsMutex.RLock()
	defer fake.subscribeToEventsMutex.RUnlock()
	return fake.subscribeToEventsArgsForCall[i].arg1
}

func (fake *FakeClient) SubscribeToEventsReturns(result1 executor.EventSource, result2 error) {
	fake.SubscribeToEventsStub = nil
	fake.subscribeToEventsReturns = struct {
		result1 executor.EventSource
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Healthy(arg1 lager.Logger) bool {
	fake.healthyMutex.Lock()
	fake.healthyArgsForCall = append(fake.healthyArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.healthyMutex.Unlock()
	if fake.HealthyStub != nil {
		return fake.HealthyStub(arg1)
	} else {
		return fake.healthyReturns.result1
	}
}

func (fake *FakeClient) HealthyCallCount() int {
	fake.healthyMutex.RLock()
	defer fake.healthyMutex.RUnlock()
	return len(fake.healthyArgsForCall)
}

func (fake *FakeClient) HealthyArgsForCall(i int) lager.Logger {
	fake.healthyMutex.RLock()
	defer fake.healthyMutex.RUnlock()
	return fake.healthyArgsForCall[i].arg1
}

func (fake *FakeClient) HealthyReturns(result1 bool) {
	fake.HealthyStub = nil
	fake.healthyReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeClient) SetHealthy(arg1 lager.Logger, arg2 bool) {
	fake.setHealthyMutex.Lock()
	fake.setHealthyArgsForCall = append(fake.setHealthyArgsForCall, struct {
		arg1 lager.Logger
		arg2 bool
	}{arg1, arg2})
	fake.setHealthyMutex.Unlock()
	if fake.SetHealthyStub != nil {
		fake.SetHealthyStub(arg1, arg2)
	}
}

func (fake *FakeClient) SetHealthyCallCount() int {
	fake.setHealthyMutex.RLock()
	defer fake.setHealthyMutex.RUnlock()
	return len(fake.setHealthyArgsForCall)
}

func (fake *FakeClient) SetHealthyArgsForCall(i int) (lager.Logger, bool) {
	fake.setHealthyMutex.RLock()
	defer fake.setHealthyMutex.RUnlock()
	return fake.setHealthyArgsForCall[i].arg1, fake.setHealthyArgsForCall[i].arg2
}

func (fake *FakeClient) Cleanup(arg1 lager.Logger) {
	fake.cleanupMutex.Lock()
	fake.cleanupArgsForCall = append(fake.cleanupArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.cleanupMutex.Unlock()
	if fake.CleanupStub != nil {
		fake.CleanupStub(arg1)
	}
}

func (fake *FakeClient) CleanupCallCount() int {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return len(fake.cleanupArgsForCall)
}

func (fake *FakeClient) CleanupArgsForCall(i int) lager.Logger {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return fake.cleanupArgsForCall[i].arg1
}

var _ executor.Client = new(FakeClient)
