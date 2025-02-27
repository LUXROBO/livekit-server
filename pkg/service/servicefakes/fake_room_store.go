// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	"context"
	"sync"
	"time"

	"github.com/livekit/livekit-server/pkg/service"
	"github.com/livekit/protocol/livekit"
)

type FakeRoomStore struct {
	DeleteParticipantStub        func(context.Context, livekit.RoomName, livekit.ParticipantIdentity) error
	deleteParticipantMutex       sync.RWMutex
	deleteParticipantArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 livekit.ParticipantIdentity
	}
	deleteParticipantReturns struct {
		result1 error
	}
	deleteParticipantReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteRoomStub        func(context.Context, livekit.RoomName) error
	deleteRoomMutex       sync.RWMutex
	deleteRoomArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
	}
	deleteRoomReturns struct {
		result1 error
	}
	deleteRoomReturnsOnCall map[int]struct {
		result1 error
	}
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
	LockRoomStub        func(context.Context, livekit.RoomName, time.Duration) (string, error)
	lockRoomMutex       sync.RWMutex
	lockRoomArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 time.Duration
	}
	lockRoomReturns struct {
		result1 string
		result2 error
	}
	lockRoomReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	StoreParticipantStub        func(context.Context, livekit.RoomName, *livekit.ParticipantInfo) error
	storeParticipantMutex       sync.RWMutex
	storeParticipantArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 *livekit.ParticipantInfo
	}
	storeParticipantReturns struct {
		result1 error
	}
	storeParticipantReturnsOnCall map[int]struct {
		result1 error
	}
	StoreRoomStub        func(context.Context, *livekit.Room) error
	storeRoomMutex       sync.RWMutex
	storeRoomArgsForCall []struct {
		arg1 context.Context
		arg2 *livekit.Room
	}
	storeRoomReturns struct {
		result1 error
	}
	storeRoomReturnsOnCall map[int]struct {
		result1 error
	}
	UnlockRoomStub        func(context.Context, livekit.RoomName, string) error
	unlockRoomMutex       sync.RWMutex
	unlockRoomArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 string
	}
	unlockRoomReturns struct {
		result1 error
	}
	unlockRoomReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRoomStore) DeleteParticipant(arg1 context.Context, arg2 livekit.RoomName, arg3 livekit.ParticipantIdentity) error {
	fake.deleteParticipantMutex.Lock()
	ret, specificReturn := fake.deleteParticipantReturnsOnCall[len(fake.deleteParticipantArgsForCall)]
	fake.deleteParticipantArgsForCall = append(fake.deleteParticipantArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 livekit.ParticipantIdentity
	}{arg1, arg2, arg3})
	stub := fake.DeleteParticipantStub
	fakeReturns := fake.deleteParticipantReturns
	fake.recordInvocation("DeleteParticipant", []interface{}{arg1, arg2, arg3})
	fake.deleteParticipantMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRoomStore) DeleteParticipantCallCount() int {
	fake.deleteParticipantMutex.RLock()
	defer fake.deleteParticipantMutex.RUnlock()
	return len(fake.deleteParticipantArgsForCall)
}

func (fake *FakeRoomStore) DeleteParticipantCalls(stub func(context.Context, livekit.RoomName, livekit.ParticipantIdentity) error) {
	fake.deleteParticipantMutex.Lock()
	defer fake.deleteParticipantMutex.Unlock()
	fake.DeleteParticipantStub = stub
}

func (fake *FakeRoomStore) DeleteParticipantArgsForCall(i int) (context.Context, livekit.RoomName, livekit.ParticipantIdentity) {
	fake.deleteParticipantMutex.RLock()
	defer fake.deleteParticipantMutex.RUnlock()
	argsForCall := fake.deleteParticipantArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRoomStore) DeleteParticipantReturns(result1 error) {
	fake.deleteParticipantMutex.Lock()
	defer fake.deleteParticipantMutex.Unlock()
	fake.DeleteParticipantStub = nil
	fake.deleteParticipantReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoomStore) DeleteParticipantReturnsOnCall(i int, result1 error) {
	fake.deleteParticipantMutex.Lock()
	defer fake.deleteParticipantMutex.Unlock()
	fake.DeleteParticipantStub = nil
	if fake.deleteParticipantReturnsOnCall == nil {
		fake.deleteParticipantReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteParticipantReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoomStore) DeleteRoom(arg1 context.Context, arg2 livekit.RoomName) error {
	fake.deleteRoomMutex.Lock()
	ret, specificReturn := fake.deleteRoomReturnsOnCall[len(fake.deleteRoomArgsForCall)]
	fake.deleteRoomArgsForCall = append(fake.deleteRoomArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
	}{arg1, arg2})
	stub := fake.DeleteRoomStub
	fakeReturns := fake.deleteRoomReturns
	fake.recordInvocation("DeleteRoom", []interface{}{arg1, arg2})
	fake.deleteRoomMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRoomStore) DeleteRoomCallCount() int {
	fake.deleteRoomMutex.RLock()
	defer fake.deleteRoomMutex.RUnlock()
	return len(fake.deleteRoomArgsForCall)
}

func (fake *FakeRoomStore) DeleteRoomCalls(stub func(context.Context, livekit.RoomName) error) {
	fake.deleteRoomMutex.Lock()
	defer fake.deleteRoomMutex.Unlock()
	fake.DeleteRoomStub = stub
}

func (fake *FakeRoomStore) DeleteRoomArgsForCall(i int) (context.Context, livekit.RoomName) {
	fake.deleteRoomMutex.RLock()
	defer fake.deleteRoomMutex.RUnlock()
	argsForCall := fake.deleteRoomArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRoomStore) DeleteRoomReturns(result1 error) {
	fake.deleteRoomMutex.Lock()
	defer fake.deleteRoomMutex.Unlock()
	fake.DeleteRoomStub = nil
	fake.deleteRoomReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoomStore) DeleteRoomReturnsOnCall(i int, result1 error) {
	fake.deleteRoomMutex.Lock()
	defer fake.deleteRoomMutex.Unlock()
	fake.DeleteRoomStub = nil
	if fake.deleteRoomReturnsOnCall == nil {
		fake.deleteRoomReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteRoomReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoomStore) ListParticipants(arg1 context.Context, arg2 livekit.RoomName) ([]*livekit.ParticipantInfo, error) {
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

func (fake *FakeRoomStore) ListParticipantsCallCount() int {
	fake.listParticipantsMutex.RLock()
	defer fake.listParticipantsMutex.RUnlock()
	return len(fake.listParticipantsArgsForCall)
}

func (fake *FakeRoomStore) ListParticipantsCalls(stub func(context.Context, livekit.RoomName) ([]*livekit.ParticipantInfo, error)) {
	fake.listParticipantsMutex.Lock()
	defer fake.listParticipantsMutex.Unlock()
	fake.ListParticipantsStub = stub
}

func (fake *FakeRoomStore) ListParticipantsArgsForCall(i int) (context.Context, livekit.RoomName) {
	fake.listParticipantsMutex.RLock()
	defer fake.listParticipantsMutex.RUnlock()
	argsForCall := fake.listParticipantsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRoomStore) ListParticipantsReturns(result1 []*livekit.ParticipantInfo, result2 error) {
	fake.listParticipantsMutex.Lock()
	defer fake.listParticipantsMutex.Unlock()
	fake.ListParticipantsStub = nil
	fake.listParticipantsReturns = struct {
		result1 []*livekit.ParticipantInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeRoomStore) ListParticipantsReturnsOnCall(i int, result1 []*livekit.ParticipantInfo, result2 error) {
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

func (fake *FakeRoomStore) ListRooms(arg1 context.Context, arg2 []livekit.RoomName) ([]*livekit.Room, error) {
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

func (fake *FakeRoomStore) ListRoomsCallCount() int {
	fake.listRoomsMutex.RLock()
	defer fake.listRoomsMutex.RUnlock()
	return len(fake.listRoomsArgsForCall)
}

func (fake *FakeRoomStore) ListRoomsCalls(stub func(context.Context, []livekit.RoomName) ([]*livekit.Room, error)) {
	fake.listRoomsMutex.Lock()
	defer fake.listRoomsMutex.Unlock()
	fake.ListRoomsStub = stub
}

func (fake *FakeRoomStore) ListRoomsArgsForCall(i int) (context.Context, []livekit.RoomName) {
	fake.listRoomsMutex.RLock()
	defer fake.listRoomsMutex.RUnlock()
	argsForCall := fake.listRoomsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRoomStore) ListRoomsReturns(result1 []*livekit.Room, result2 error) {
	fake.listRoomsMutex.Lock()
	defer fake.listRoomsMutex.Unlock()
	fake.ListRoomsStub = nil
	fake.listRoomsReturns = struct {
		result1 []*livekit.Room
		result2 error
	}{result1, result2}
}

func (fake *FakeRoomStore) ListRoomsReturnsOnCall(i int, result1 []*livekit.Room, result2 error) {
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

func (fake *FakeRoomStore) LoadParticipant(arg1 context.Context, arg2 livekit.RoomName, arg3 livekit.ParticipantIdentity) (*livekit.ParticipantInfo, error) {
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

func (fake *FakeRoomStore) LoadParticipantCallCount() int {
	fake.loadParticipantMutex.RLock()
	defer fake.loadParticipantMutex.RUnlock()
	return len(fake.loadParticipantArgsForCall)
}

func (fake *FakeRoomStore) LoadParticipantCalls(stub func(context.Context, livekit.RoomName, livekit.ParticipantIdentity) (*livekit.ParticipantInfo, error)) {
	fake.loadParticipantMutex.Lock()
	defer fake.loadParticipantMutex.Unlock()
	fake.LoadParticipantStub = stub
}

func (fake *FakeRoomStore) LoadParticipantArgsForCall(i int) (context.Context, livekit.RoomName, livekit.ParticipantIdentity) {
	fake.loadParticipantMutex.RLock()
	defer fake.loadParticipantMutex.RUnlock()
	argsForCall := fake.loadParticipantArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRoomStore) LoadParticipantReturns(result1 *livekit.ParticipantInfo, result2 error) {
	fake.loadParticipantMutex.Lock()
	defer fake.loadParticipantMutex.Unlock()
	fake.LoadParticipantStub = nil
	fake.loadParticipantReturns = struct {
		result1 *livekit.ParticipantInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeRoomStore) LoadParticipantReturnsOnCall(i int, result1 *livekit.ParticipantInfo, result2 error) {
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

func (fake *FakeRoomStore) LoadRoom(arg1 context.Context, arg2 livekit.RoomName) (*livekit.Room, error) {
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

func (fake *FakeRoomStore) LoadRoomCallCount() int {
	fake.loadRoomMutex.RLock()
	defer fake.loadRoomMutex.RUnlock()
	return len(fake.loadRoomArgsForCall)
}

func (fake *FakeRoomStore) LoadRoomCalls(stub func(context.Context, livekit.RoomName) (*livekit.Room, error)) {
	fake.loadRoomMutex.Lock()
	defer fake.loadRoomMutex.Unlock()
	fake.LoadRoomStub = stub
}

func (fake *FakeRoomStore) LoadRoomArgsForCall(i int) (context.Context, livekit.RoomName) {
	fake.loadRoomMutex.RLock()
	defer fake.loadRoomMutex.RUnlock()
	argsForCall := fake.loadRoomArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRoomStore) LoadRoomReturns(result1 *livekit.Room, result2 error) {
	fake.loadRoomMutex.Lock()
	defer fake.loadRoomMutex.Unlock()
	fake.LoadRoomStub = nil
	fake.loadRoomReturns = struct {
		result1 *livekit.Room
		result2 error
	}{result1, result2}
}

func (fake *FakeRoomStore) LoadRoomReturnsOnCall(i int, result1 *livekit.Room, result2 error) {
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

func (fake *FakeRoomStore) LockRoom(arg1 context.Context, arg2 livekit.RoomName, arg3 time.Duration) (string, error) {
	fake.lockRoomMutex.Lock()
	ret, specificReturn := fake.lockRoomReturnsOnCall[len(fake.lockRoomArgsForCall)]
	fake.lockRoomArgsForCall = append(fake.lockRoomArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 time.Duration
	}{arg1, arg2, arg3})
	stub := fake.LockRoomStub
	fakeReturns := fake.lockRoomReturns
	fake.recordInvocation("LockRoom", []interface{}{arg1, arg2, arg3})
	fake.lockRoomMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRoomStore) LockRoomCallCount() int {
	fake.lockRoomMutex.RLock()
	defer fake.lockRoomMutex.RUnlock()
	return len(fake.lockRoomArgsForCall)
}

func (fake *FakeRoomStore) LockRoomCalls(stub func(context.Context, livekit.RoomName, time.Duration) (string, error)) {
	fake.lockRoomMutex.Lock()
	defer fake.lockRoomMutex.Unlock()
	fake.LockRoomStub = stub
}

func (fake *FakeRoomStore) LockRoomArgsForCall(i int) (context.Context, livekit.RoomName, time.Duration) {
	fake.lockRoomMutex.RLock()
	defer fake.lockRoomMutex.RUnlock()
	argsForCall := fake.lockRoomArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRoomStore) LockRoomReturns(result1 string, result2 error) {
	fake.lockRoomMutex.Lock()
	defer fake.lockRoomMutex.Unlock()
	fake.LockRoomStub = nil
	fake.lockRoomReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRoomStore) LockRoomReturnsOnCall(i int, result1 string, result2 error) {
	fake.lockRoomMutex.Lock()
	defer fake.lockRoomMutex.Unlock()
	fake.LockRoomStub = nil
	if fake.lockRoomReturnsOnCall == nil {
		fake.lockRoomReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.lockRoomReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRoomStore) StoreParticipant(arg1 context.Context, arg2 livekit.RoomName, arg3 *livekit.ParticipantInfo) error {
	fake.storeParticipantMutex.Lock()
	ret, specificReturn := fake.storeParticipantReturnsOnCall[len(fake.storeParticipantArgsForCall)]
	fake.storeParticipantArgsForCall = append(fake.storeParticipantArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 *livekit.ParticipantInfo
	}{arg1, arg2, arg3})
	stub := fake.StoreParticipantStub
	fakeReturns := fake.storeParticipantReturns
	fake.recordInvocation("StoreParticipant", []interface{}{arg1, arg2, arg3})
	fake.storeParticipantMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRoomStore) StoreParticipantCallCount() int {
	fake.storeParticipantMutex.RLock()
	defer fake.storeParticipantMutex.RUnlock()
	return len(fake.storeParticipantArgsForCall)
}

func (fake *FakeRoomStore) StoreParticipantCalls(stub func(context.Context, livekit.RoomName, *livekit.ParticipantInfo) error) {
	fake.storeParticipantMutex.Lock()
	defer fake.storeParticipantMutex.Unlock()
	fake.StoreParticipantStub = stub
}

func (fake *FakeRoomStore) StoreParticipantArgsForCall(i int) (context.Context, livekit.RoomName, *livekit.ParticipantInfo) {
	fake.storeParticipantMutex.RLock()
	defer fake.storeParticipantMutex.RUnlock()
	argsForCall := fake.storeParticipantArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRoomStore) StoreParticipantReturns(result1 error) {
	fake.storeParticipantMutex.Lock()
	defer fake.storeParticipantMutex.Unlock()
	fake.StoreParticipantStub = nil
	fake.storeParticipantReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoomStore) StoreParticipantReturnsOnCall(i int, result1 error) {
	fake.storeParticipantMutex.Lock()
	defer fake.storeParticipantMutex.Unlock()
	fake.StoreParticipantStub = nil
	if fake.storeParticipantReturnsOnCall == nil {
		fake.storeParticipantReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.storeParticipantReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoomStore) StoreRoom(arg1 context.Context, arg2 *livekit.Room) error {
	fake.storeRoomMutex.Lock()
	ret, specificReturn := fake.storeRoomReturnsOnCall[len(fake.storeRoomArgsForCall)]
	fake.storeRoomArgsForCall = append(fake.storeRoomArgsForCall, struct {
		arg1 context.Context
		arg2 *livekit.Room
	}{arg1, arg2})
	stub := fake.StoreRoomStub
	fakeReturns := fake.storeRoomReturns
	fake.recordInvocation("StoreRoom", []interface{}{arg1, arg2})
	fake.storeRoomMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRoomStore) StoreRoomCallCount() int {
	fake.storeRoomMutex.RLock()
	defer fake.storeRoomMutex.RUnlock()
	return len(fake.storeRoomArgsForCall)
}

func (fake *FakeRoomStore) StoreRoomCalls(stub func(context.Context, *livekit.Room) error) {
	fake.storeRoomMutex.Lock()
	defer fake.storeRoomMutex.Unlock()
	fake.StoreRoomStub = stub
}

func (fake *FakeRoomStore) StoreRoomArgsForCall(i int) (context.Context, *livekit.Room) {
	fake.storeRoomMutex.RLock()
	defer fake.storeRoomMutex.RUnlock()
	argsForCall := fake.storeRoomArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRoomStore) StoreRoomReturns(result1 error) {
	fake.storeRoomMutex.Lock()
	defer fake.storeRoomMutex.Unlock()
	fake.StoreRoomStub = nil
	fake.storeRoomReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoomStore) StoreRoomReturnsOnCall(i int, result1 error) {
	fake.storeRoomMutex.Lock()
	defer fake.storeRoomMutex.Unlock()
	fake.StoreRoomStub = nil
	if fake.storeRoomReturnsOnCall == nil {
		fake.storeRoomReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.storeRoomReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoomStore) UnlockRoom(arg1 context.Context, arg2 livekit.RoomName, arg3 string) error {
	fake.unlockRoomMutex.Lock()
	ret, specificReturn := fake.unlockRoomReturnsOnCall[len(fake.unlockRoomArgsForCall)]
	fake.unlockRoomArgsForCall = append(fake.unlockRoomArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.UnlockRoomStub
	fakeReturns := fake.unlockRoomReturns
	fake.recordInvocation("UnlockRoom", []interface{}{arg1, arg2, arg3})
	fake.unlockRoomMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRoomStore) UnlockRoomCallCount() int {
	fake.unlockRoomMutex.RLock()
	defer fake.unlockRoomMutex.RUnlock()
	return len(fake.unlockRoomArgsForCall)
}

func (fake *FakeRoomStore) UnlockRoomCalls(stub func(context.Context, livekit.RoomName, string) error) {
	fake.unlockRoomMutex.Lock()
	defer fake.unlockRoomMutex.Unlock()
	fake.UnlockRoomStub = stub
}

func (fake *FakeRoomStore) UnlockRoomArgsForCall(i int) (context.Context, livekit.RoomName, string) {
	fake.unlockRoomMutex.RLock()
	defer fake.unlockRoomMutex.RUnlock()
	argsForCall := fake.unlockRoomArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRoomStore) UnlockRoomReturns(result1 error) {
	fake.unlockRoomMutex.Lock()
	defer fake.unlockRoomMutex.Unlock()
	fake.UnlockRoomStub = nil
	fake.unlockRoomReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoomStore) UnlockRoomReturnsOnCall(i int, result1 error) {
	fake.unlockRoomMutex.Lock()
	defer fake.unlockRoomMutex.Unlock()
	fake.UnlockRoomStub = nil
	if fake.unlockRoomReturnsOnCall == nil {
		fake.unlockRoomReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unlockRoomReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoomStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteParticipantMutex.RLock()
	defer fake.deleteParticipantMutex.RUnlock()
	fake.deleteRoomMutex.RLock()
	defer fake.deleteRoomMutex.RUnlock()
	fake.listParticipantsMutex.RLock()
	defer fake.listParticipantsMutex.RUnlock()
	fake.listRoomsMutex.RLock()
	defer fake.listRoomsMutex.RUnlock()
	fake.loadParticipantMutex.RLock()
	defer fake.loadParticipantMutex.RUnlock()
	fake.loadRoomMutex.RLock()
	defer fake.loadRoomMutex.RUnlock()
	fake.lockRoomMutex.RLock()
	defer fake.lockRoomMutex.RUnlock()
	fake.storeParticipantMutex.RLock()
	defer fake.storeParticipantMutex.RUnlock()
	fake.storeRoomMutex.RLock()
	defer fake.storeRoomMutex.RUnlock()
	fake.unlockRoomMutex.RLock()
	defer fake.unlockRoomMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRoomStore) recordInvocation(key string, args []interface{}) {
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

var _ service.RoomStore = new(FakeRoomStore)
