package routing

import (
	"context"
	"sync"
	"time"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"github.com/livekit/protocol/utils"
	"google.golang.org/protobuf/proto"
)

// a router of messages on the same node, basic implementation for local testing
type LocalRouter struct {
	currentNode LocalNode
	lock        sync.RWMutex
	// channels for each participant
	requestChannels  map[string]*MessageChannel
	responseChannels map[string]*MessageChannel
	isStarted        utils.AtomicFlag

	rtcMessageChan *MessageChannel

	onNewParticipant NewParticipantCallback
	onRTCMessage     RTCMessageCallback
}

func NewLocalRouter(currentNode LocalNode) *LocalRouter {
	return &LocalRouter{
		currentNode:      currentNode,
		requestChannels:  make(map[string]*MessageChannel),
		responseChannels: make(map[string]*MessageChannel),
		rtcMessageChan:   NewMessageChannel(),
	}
}

func (r *LocalRouter) GetNodeForRoom(_ context.Context, _ livekit.RoomName) (*livekit.Node, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	node := proto.Clone((*livekit.Node)(r.currentNode)).(*livekit.Node)
	return node, nil
}

func (r *LocalRouter) SetNodeForRoom(_ context.Context, _ livekit.RoomName, _ string) error {
	return nil
}

func (r *LocalRouter) ClearRoomState(_ context.Context, _ livekit.RoomName) error {
	// do nothing
	return nil
}

func (r *LocalRouter) RegisterNode() error {
	return nil
}

func (r *LocalRouter) UnregisterNode() error {
	return nil
}

func (r *LocalRouter) RemoveDeadNodes() error {
	return nil
}

func (r *LocalRouter) GetNode(nodeID string) (*livekit.Node, error) {
	if nodeID == r.currentNode.Id {
		return r.currentNode, nil
	}
	return nil, ErrNotFound
}

func (r *LocalRouter) ListNodes() ([]*livekit.Node, error) {
	return []*livekit.Node{
		r.currentNode,
	}, nil
}

func (r *LocalRouter) StartParticipantSignal(ctx context.Context, roomName livekit.RoomName, pi ParticipantInit) (connectionID string, reqSink MessageSink, resSource MessageSource, err error) {
	// treat it as a new participant connecting
	if r.onNewParticipant == nil {
		err = ErrHandlerNotDefined
		return
	}

	// index channels by roomName | identity
	key := participantKey(roomName, pi.Identity)

	// close older channels if one already exists
	reqChan := r.getMessageChannel(r.requestChannels, key)
	if reqChan != nil {
		reqChan.Close()
	}
	resChan := r.getMessageChannel(r.responseChannels, key)
	if resChan != nil {
		resChan.Close()
	}
	reqChan = r.getOrCreateMessageChannel(r.requestChannels, key)
	resChan = r.getOrCreateMessageChannel(r.responseChannels, key)

	r.onNewParticipant(
		ctx,
		roomName,
		pi,
		// request source
		reqChan,
		// response sink
		resChan,
	)
	return string(pi.Identity), reqChan, resChan, nil
}

func (r *LocalRouter) WriteParticipantRTC(_ context.Context, roomName livekit.RoomName, identity livekit.ParticipantIdentity, msg *livekit.RTCNodeMessage) error {
	if r.rtcMessageChan.IsClosed() {
		// create a new one
		r.rtcMessageChan = NewMessageChannel()
	}
	msg.ParticipantKey = participantKey(roomName, identity)
	return r.writeRTCMessage(r.rtcMessageChan, msg)
}

func (r *LocalRouter) WriteRoomRTC(ctx context.Context, roomName livekit.RoomName, identity livekit.ParticipantIdentity, msg *livekit.RTCNodeMessage) error {
	msg.ParticipantKey = participantKey(roomName, identity)
	return r.WriteNodeRTC(ctx, r.currentNode.Id, msg)
}

func (r *LocalRouter) WriteNodeRTC(_ context.Context, _ string, msg *livekit.RTCNodeMessage) error {
	if r.rtcMessageChan.IsClosed() {
		// create a new one
		r.rtcMessageChan = NewMessageChannel()
	}
	return r.writeRTCMessage(r.rtcMessageChan, msg)
}

func (r *LocalRouter) writeRTCMessage(sink MessageSink, msg *livekit.RTCNodeMessage) error {
	defer sink.Close()
	msg.SenderTime = time.Now().Unix()
	return sink.WriteMessage(msg)
}

func (r *LocalRouter) OnNewParticipantRTC(callback NewParticipantCallback) {
	r.onNewParticipant = callback
}

func (r *LocalRouter) OnRTCMessage(callback RTCMessageCallback) {
	r.onRTCMessage = callback
}

func (r *LocalRouter) Start() error {
	if !r.isStarted.TrySet(true) {
		return nil
	}
	go r.statsWorker()
	// on local routers, Start doesn't do anything, websocket connections initiate the connections
	go r.rtcMessageWorker()
	return nil
}

func (r *LocalRouter) Drain() {
	r.currentNode.State = livekit.NodeState_SHUTTING_DOWN
}

func (r *LocalRouter) Stop() {
	r.rtcMessageChan.Close()
}

func (r *LocalRouter) statsWorker() {
	for {
		if !r.isStarted.Get() {
			return
		}
		// update every 10 seconds
		<-time.After(statsUpdateInterval)
		r.lock.Lock()
		r.currentNode.Stats.UpdatedAt = time.Now().Unix()
		r.lock.Unlock()
	}
}

func (r *LocalRouter) rtcMessageWorker() {
	// is a new channel available? if so swap to that one
	if !r.isStarted.Get() {
		return
	}

	// start a new worker after this finished
	defer func() {
		go r.rtcMessageWorker()
	}()

	if r.rtcMessageChan.IsClosed() {
		// sleep and retry
		time.Sleep(time.Second)
	}

	// consume messages from
	for msg := range r.rtcMessageChan.ReadChan() {
		if rtcMsg, ok := msg.(*livekit.RTCNodeMessage); ok {
			room, identity, err := parseParticipantKey(rtcMsg.ParticipantKey)
			if err != nil {
				logger.Errorw("could not process RTC message", err)
				continue
			}
			if r.onRTCMessage != nil {
				r.onRTCMessage(context.Background(), room, identity, rtcMsg)
			}
		}
	}
}

func (r *LocalRouter) getMessageChannel(target map[string]*MessageChannel, key string) *MessageChannel {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return target[key]
}

func (r *LocalRouter) getOrCreateMessageChannel(target map[string]*MessageChannel, key string) *MessageChannel {
	r.lock.Lock()
	defer r.lock.Unlock()
	mc := target[key]

	if mc != nil {
		return mc
	}

	mc = NewMessageChannel()
	mc.OnClose(func() {
		r.lock.Lock()
		delete(target, key)
		r.lock.Unlock()
	})
	target[key] = mc

	return mc
}
