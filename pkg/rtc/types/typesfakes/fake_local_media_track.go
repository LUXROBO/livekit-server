// Code generated by counterfeiter. DO NOT EDIT.
package typesfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/rtc/types"
)

type FakeLocalMediaTrack struct {
	NotifySubscriberNodeMediaLossStub        func(string, uint8)
	notifySubscriberNodeMediaLossMutex       sync.RWMutex
	notifySubscriberNodeMediaLossArgsForCall []struct {
		arg1 string
		arg2 uint8
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLocalMediaTrack) NotifySubscriberNodeMediaLoss(arg1 string, arg2 uint8) {
	fake.notifySubscriberNodeMediaLossMutex.Lock()
	fake.notifySubscriberNodeMediaLossArgsForCall = append(fake.notifySubscriberNodeMediaLossArgsForCall, struct {
		arg1 string
		arg2 uint8
	}{arg1, arg2})
	stub := fake.NotifySubscriberNodeMediaLossStub
	fake.recordInvocation("NotifySubscriberNodeMediaLoss", []interface{}{arg1, arg2})
	fake.notifySubscriberNodeMediaLossMutex.Unlock()
	if stub != nil {
		fake.NotifySubscriberNodeMediaLossStub(arg1, arg2)
	}
}

func (fake *FakeLocalMediaTrack) NotifySubscriberNodeMediaLossCallCount() int {
	fake.notifySubscriberNodeMediaLossMutex.RLock()
	defer fake.notifySubscriberNodeMediaLossMutex.RUnlock()
	return len(fake.notifySubscriberNodeMediaLossArgsForCall)
}

func (fake *FakeLocalMediaTrack) NotifySubscriberNodeMediaLossCalls(stub func(string, uint8)) {
	fake.notifySubscriberNodeMediaLossMutex.Lock()
	defer fake.notifySubscriberNodeMediaLossMutex.Unlock()
	fake.NotifySubscriberNodeMediaLossStub = stub
}

func (fake *FakeLocalMediaTrack) NotifySubscriberNodeMediaLossArgsForCall(i int) (string, uint8) {
	fake.notifySubscriberNodeMediaLossMutex.RLock()
	defer fake.notifySubscriberNodeMediaLossMutex.RUnlock()
	argsForCall := fake.notifySubscriberNodeMediaLossArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLocalMediaTrack) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.notifySubscriberNodeMediaLossMutex.RLock()
	defer fake.notifySubscriberNodeMediaLossMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLocalMediaTrack) recordInvocation(key string, args []interface{}) {
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

var _ types.LocalMediaTrack = new(FakeLocalMediaTrack)
