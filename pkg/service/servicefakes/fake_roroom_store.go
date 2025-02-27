// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	"context"
	"sync"

	"github.com/livekit/livekit-server/pkg/service"
	"github.com/livekit/protocol/livekit"
)

type FakeRORoomStore struct {
	ListParticipantsStub        func(context.Context, livekit.RoomName) ([]*livekit.ParticipantInfo, error)
	listParticipantsMutex       sync.RWMutex
	listParticipantsArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
	}
	listParticipantsReturns struct {
		result1 []*livekit.ParticipantInfo
		result2 error
	}
	listParticipantsReturnsOnCall map[int]struct {
		result1 []*livekit.ParticipantInfo
		result2 error
	}
	ListRoomsStub        func(context.Context, []livekit.RoomName) ([]*livekit.Room, error)
	listRoomsMutex       sync.RWMutex
	listRoomsArgsForCall []struct {
		arg1 context.Context
		arg2 []livekit.RoomName
	}
	listRoomsReturns struct {
		result1 []*livekit.Room
		result2 error
	}
	listRoomsReturnsOnCall map[int]struct {
		result1 []*livekit.Room
		result2 error
	}
	LoadParticipantStub        func(context.Context, livekit.RoomName, livekit.ParticipantIdentity) (*livekit.ParticipantInfo, error)
	loadParticipantMutex       sync.RWMutex
	loadParticipantArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 livekit.ParticipantIdentity
	}
	loadParticipantReturns struct {
		result1 *livekit.ParticipantInfo
		result2 error
	}
	loadParticipantReturnsOnCall map[int]struct {
		result1 *livekit.ParticipantInfo
		result2 error
	}
	LoadRoomStub        func(context.Context, livekit.RoomName) (*livekit.Room, error)
	loadRoomMutex       sync.RWMutex
	loadRoomArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
	}
	loadRoomReturns struct {
		result1 *livekit.Room
		result2 error
	}
	loadRoomReturnsOnCall map[int]struct {
		result1 *livekit.Room
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRORoomStore) ListParticipants(arg1 context.Context, arg2 livekit.RoomName) ([]*livekit.ParticipantInfo, error) {
	fake.listParticipantsMutex.Lock()
	ret, specificReturn := fake.listParticipantsReturnsOnCall[len(fake.listParticipantsArgsForCall)]
	fake.listParticipantsArgsForCall = append(fake.listParticipantsArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
	}{arg1, arg2})
	stub := fake.ListParticipantsStub
	fakeReturns := fake.listParticipantsReturns
	fake.recordInvocation("ListParticipants", []interface{}{arg1, arg2})
	fake.listParticipantsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRORoomStore) ListParticipantsCallCount() int {
	fake.listParticipantsMutex.RLock()
	defer fake.listParticipantsMutex.RUnlock()
	return len(fake.listParticipantsArgsForCall)
}

func (fake *FakeRORoomStore) ListParticipantsCalls(stub func(context.Context, livekit.RoomName) ([]*livekit.ParticipantInfo, error)) {
	fake.listParticipantsMutex.Lock()
	defer fake.listParticipantsMutex.Unlock()
	fake.ListParticipantsStub = stub
}

func (fake *FakeRORoomStore) ListParticipantsArgsForCall(i int) (context.Context, livekit.RoomName) {
	fake.listParticipantsMutex.RLock()
	defer fake.listParticipantsMutex.RUnlock()
	argsForCall := fake.listParticipantsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRORoomStore) ListParticipantsReturns(result1 []*livekit.ParticipantInfo, result2 error) {
	fake.listParticipantsMutex.Lock()
	defer fake.listParticipantsMutex.Unlock()
	fake.ListParticipantsStub = nil
	fake.listParticipantsReturns = struct {
		result1 []*livekit.ParticipantInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeRORoomStore) ListParticipantsReturnsOnCall(i int, result1 []*livekit.ParticipantInfo, result2 error) {
	fake.listParticipantsMutex.Lock()
	defer fake.listParticipantsMutex.Unlock()
	fake.ListParticipantsStub = nil
	if fake.listParticipantsReturnsOnCall == nil {
		fake.listParticipantsReturnsOnCall = make(map[int]struct {
			result1 []*livekit.ParticipantInfo
			result2 error
		})
	}
	fake.listParticipantsReturnsOnCall[i] = struct {
		result1 []*livekit.ParticipantInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeRORoomStore) ListRooms(arg1 context.Context, arg2 []livekit.RoomName) ([]*livekit.Room, error) {
	var arg2Copy []livekit.RoomName
	if arg2 != nil {
		arg2Copy = make([]livekit.RoomName, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.listRoomsMutex.Lock()
	ret, specificReturn := fake.listRoomsReturnsOnCall[len(fake.listRoomsArgsForCall)]
	fake.listRoomsArgsForCall = append(fake.listRoomsArgsForCall, struct {
		arg1 context.Context
		arg2 []livekit.RoomName
	}{arg1, arg2Copy})
	stub := fake.ListRoomsStub
	fakeReturns := fake.listRoomsReturns
	fake.recordInvocation("ListRooms", []interface{}{arg1, arg2Copy})
	fake.listRoomsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRORoomStore) ListRoomsCallCount() int {
	fake.listRoomsMutex.RLock()
	defer fake.listRoomsMutex.RUnlock()
	return len(fake.listRoomsArgsForCall)
}

func (fake *FakeRORoomStore) ListRoomsCalls(stub func(context.Context, []livekit.RoomName) ([]*livekit.Room, error)) {
	fake.listRoomsMutex.Lock()
	defer fake.listRoomsMutex.Unlock()
	fake.ListRoomsStub = stub
}

func (fake *FakeRORoomStore) ListRoomsArgsForCall(i int) (context.Context, []livekit.RoomName) {
	fake.listRoomsMutex.RLock()
	defer fake.listRoomsMutex.RUnlock()
	argsForCall := fake.listRoomsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRORoomStore) ListRoomsReturns(result1 []*livekit.Room, result2 error) {
	fake.listRoomsMutex.Lock()
	defer fake.listRoomsMutex.Unlock()
	fake.ListRoomsStub = nil
	fake.listRoomsReturns = struct {
		result1 []*livekit.Room
		result2 error
	}{result1, result2}
}

func (fake *FakeRORoomStore) ListRoomsReturnsOnCall(i int, result1 []*livekit.Room, result2 error) {
	fake.listRoomsMutex.Lock()
	defer fake.listRoomsMutex.Unlock()
	fake.ListRoomsStub = nil
	if fake.listRoomsReturnsOnCall == nil {
		fake.listRoomsReturnsOnCall = make(map[int]struct {
			result1 []*livekit.Room
			result2 error
		})
	}
	fake.listRoomsReturnsOnCall[i] = struct {
		result1 []*livekit.Room
		result2 error
	}{result1, result2}
}

func (fake *FakeRORoomStore) LoadParticipant(arg1 context.Context, arg2 livekit.RoomName, arg3 livekit.ParticipantIdentity) (*livekit.ParticipantInfo, error) {
	fake.loadParticipantMutex.Lock()
	ret, specificReturn := fake.loadParticipantReturnsOnCall[len(fake.loadParticipantArgsForCall)]
	fake.loadParticipantArgsForCall = append(fake.loadParticipantArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 livekit.ParticipantIdentity
	}{arg1, arg2, arg3})
	stub := fake.LoadParticipantStub
	fakeReturns := fake.loadParticipantReturns
	fake.recordInvocation("LoadParticipant", []interface{}{arg1, arg2, arg3})
	fake.loadParticipantMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRORoomStore) LoadParticipantCallCount() int {
	fake.loadParticipantMutex.RLock()
	defer fake.loadParticipantMutex.RUnlock()
	return len(fake.loadParticipantArgsForCall)
}

func (fake *FakeRORoomStore) LoadParticipantCalls(stub func(context.Context, livekit.RoomName, livekit.ParticipantIdentity) (*livekit.ParticipantInfo, error)) {
	fake.loadParticipantMutex.Lock()
	defer fake.loadParticipantMutex.Unlock()
	fake.LoadParticipantStub = stub
}

func (fake *FakeRORoomStore) LoadParticipantArgsForCall(i int) (context.Context, livekit.RoomName, livekit.ParticipantIdentity) {
	fake.loadParticipantMutex.RLock()
	defer fake.loadParticipantMutex.RUnlock()
	argsForCall := fake.loadParticipantArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRORoomStore) LoadParticipantReturns(result1 *livekit.ParticipantInfo, result2 error) {
	fake.loadParticipantMutex.Lock()
	defer fake.loadParticipantMutex.Unlock()
	fake.LoadParticipantStub = nil
	fake.loadParticipantReturns = struct {
		result1 *livekit.ParticipantInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeRORoomStore) LoadParticipantReturnsOnCall(i int, result1 *livekit.ParticipantInfo, result2 error) {
	fake.loadParticipantMutex.Lock()
	defer fake.loadParticipantMutex.Unlock()
	fake.LoadParticipantStub = nil
	if fake.loadParticipantReturnsOnCall == nil {
		fake.loadParticipantReturnsOnCall = make(map[int]struct {
			result1 *livekit.ParticipantInfo
			result2 error
		})
	}
	fake.loadParticipantReturnsOnCall[i] = struct {
		result1 *livekit.ParticipantInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeRORoomStore) LoadRoom(arg1 context.Context, arg2 livekit.RoomName) (*livekit.Room, error) {
	fake.loadRoomMutex.Lock()
	ret, specificReturn := fake.loadRoomReturnsOnCall[len(fake.loadRoomArgsForCall)]
	fake.loadRoomArgsForCall = append(fake.loadRoomArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
	}{arg1, arg2})
	stub := fake.LoadRoomStub
	fakeReturns := fake.loadRoomReturns
	fake.recordInvocation("LoadRoom", []interface{}{arg1, arg2})
	fake.loadRoomMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRORoomStore) LoadRoomCallCount() int {
	fake.loadRoomMutex.RLock()
	defer fake.loadRoomMutex.RUnlock()
	return len(fake.loadRoomArgsForCall)
}

func (fake *FakeRORoomStore) LoadRoomCalls(stub func(context.Context, livekit.RoomName) (*livekit.Room, error)) {
	fake.loadRoomMutex.Lock()
	defer fake.loadRoomMutex.Unlock()
	fake.LoadRoomStub = stub
}

func (fake *FakeRORoomStore) LoadRoomArgsForCall(i int) (context.Context, livekit.RoomName) {
	fake.loadRoomMutex.RLock()
	defer fake.loadRoomMutex.RUnlock()
	argsForCall := fake.loadRoomArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRORoomStore) LoadRoomReturns(result1 *livekit.Room, result2 error) {
	fake.loadRoomMutex.Lock()
	defer fake.loadRoomMutex.Unlock()
	fake.LoadRoomStub = nil
	fake.loadRoomReturns = struct {
		result1 *livekit.Room
		result2 error
	}{result1, result2}
}

func (fake *FakeRORoomStore) LoadRoomReturnsOnCall(i int, result1 *livekit.Room, result2 error) {
	fake.loadRoomMutex.Lock()
	defer fake.loadRoomMutex.Unlock()
	fake.LoadRoomStub = nil
	if fake.loadRoomReturnsOnCall == nil {
		fake.loadRoomReturnsOnCall = make(map[int]struct {
			result1 *livekit.Room
			result2 error
		})
	}
	fake.loadRoomReturnsOnCall[i] = struct {
		result1 *livekit.Room
		result2 error
	}{result1, result2}
}

func (fake *FakeRORoomStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listParticipantsMutex.RLock()
	defer fake.listParticipantsMutex.RUnlock()
	fake.listRoomsMutex.RLock()
	defer fake.listRoomsMutex.RUnlock()
	fake.loadParticipantMutex.RLock()
	defer fake.loadParticipantMutex.RUnlock()
	fake.loadRoomMutex.RLock()
	defer fake.loadRoomMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRORoomStore) recordInvocation(key string, args []interface{}) {
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

var _ service.RORoomStore = new(FakeRORoomStore)
