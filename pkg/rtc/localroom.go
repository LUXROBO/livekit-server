package rtc

import (
	"math"
	"sort"
	"sync"
	"time"

	"github.com/go-logr/logr"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"google.golang.org/protobuf/proto"

	"github.com/livekit/livekit-server/pkg/config"
	"github.com/livekit/livekit-server/pkg/routing"
	"github.com/livekit/livekit-server/pkg/rtc/types"
	"github.com/livekit/livekit-server/pkg/sfu/buffer"
	"github.com/livekit/livekit-server/pkg/telemetry"
	"github.com/livekit/livekit-server/pkg/telemetry/prometheus"
)

const (
	/* RAJA-TODO
	   DefaultEmptyTimeout       = 5 * 60 // 5m
	   DefaultRoomDepartureGrace = 20
	   RAJA-TODO */
	AudioLevelQuantization = 8 // ideally power of 2 to minimize float decimal
)

type LocalRoom struct {
	lock sync.RWMutex

	Room   *livekit.Room
	Logger logger.Logger

	// RAJA-TODO config      WebRTCConfig
	audioConfig *config.AudioConfig
	telemetry   telemetry.TelemetryService

	// map of identity -> types.Participant
	participants    map[livekit.ParticipantIdentity]types.Participant
	participantOpts map[livekit.ParticipantIdentity]*ParticipantOptions

	bufferFactory *buffer.Factory

	/* RAJA-TODO
	// time the first participant joined the room
	joinedAt atomic.Value
	// time that the last participant left the room
	leftAt    atomic.Value
	closed    chan struct{}
	closeOnce sync.Once

	onParticipantChanged func(p types.Participant)
	onMetadataUpdate     func(metadata string)
	onClose              func()
	RAJA-TODO */
}

type ParticipantOptions struct {
	AutoSubscribe bool
	SyncState     *livekit.SyncState
	MigrateIn     bool
	MigrateOut    bool
}

func NewLocalRoom(
	room *livekit.Room,
	config WebRTCConfig,
	audioConfig *config.AudioConfig,
	telemetry telemetry.TelemetryService,
	logger logger.Logger,
) *LocalRoom {
	r := &LocalRoom{
		Room:            proto.Clone(room).(*livekit.Room),
		Logger:          logger,
		audioConfig:     audioConfig,
		telemetry:       telemetry,
		participants:    make(map[livekit.ParticipantIdentity]types.Participant),
		participantOpts: make(map[livekit.ParticipantIdentity]*ParticipantOptions),
		bufferFactory:   buffer.NewBufferFactory(config.Receiver.PacketBufferSize, logr.Logger{}),
		// RAJA-TODO closed:          make(chan struct{}),
	}
	/* RAJA-TODO
	if r.Room.EmptyTimeout == 0 {
		r.Room.EmptyTimeout = DefaultEmptyTimeout
	}
	if r.Room.CreationTime == 0 {
		r.Room.CreationTime = time.Now().Unix()
	}

	go r.audioUpdateWorker()
	go r.connectionQualityWorker()
	RAJA-TODO */

	return r
}

func (r *LocalRoom) GetParticipant(identity livekit.ParticipantIdentity) types.Participant {
	r.lock.RLock()
	defer r.lock.RUnlock()

	return r.participants[identity]
}

func (r *LocalRoom) GetParticipantByID(participantID livekit.ParticipantID) types.Participant {
	r.lock.RLock()
	defer r.lock.RUnlock()

	for _, p := range r.participants {
		if p.ID() == participantID {
			return p
		}
	}

	return nil
}

func (r *LocalRoom) GetParticipants() []types.Participant {
	r.lock.RLock()
	defer r.lock.RUnlock()
	participants := make([]types.Participant, 0, len(r.participants))
	for _, p := range r.participants {
		participants = append(participants, p)
	}
	return participants
}

func (r *LocalRoom) GetActiveSpeakers() []*livekit.SpeakerInfo {
	participants := r.GetParticipants()
	speakers := make([]*livekit.SpeakerInfo, 0, len(participants))
	for _, p := range participants {
		level, active := p.GetAudioLevel()
		if !active {
			continue
		}
		speakers = append(speakers, &livekit.SpeakerInfo{
			Sid:    string(p.ID()),
			Level:  ConvertAudioLevel(level),
			Active: active,
		})
	}

	sort.Slice(speakers, func(i, j int) bool {
		return speakers[i].Level > speakers[j].Level
	})
	return speakers
}

func (r *LocalRoom) GetBufferFactory() *buffer.Factory {
	return r.bufferFactory
}

/* RAJA-TODO
func (r *LocalRoom) FirstJoinedAt() int64 {
	j := r.joinedAt.Load()
	if t, ok := j.(int64); ok {
		return t
	}
	return 0
}

func (r *Room) LastLeftAt() int64 {
	l := r.leftAt.Load()
	if t, ok := l.(int64); ok {
		return t
	}
	return 0
}
RAJA-TODO */

func (r *LocalRoom) Join(participant types.Participant, opts *ParticipantOptions, iceServers []*livekit.ICEServer) error {
	if r.IsClosed() {
		prometheus.ServiceOperationCounter.WithLabelValues("participant_join", "error", "room_closed").Add(1)
		return ErrRoomClosed
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	if r.participants[participant.Identity()] != nil {
		prometheus.ServiceOperationCounter.WithLabelValues("participant_join", "error", "already_joined").Add(1)
		return ErrAlreadyJoined
	}

	// RAJA-TODO: this needs a callback to check if participant limit is exceeded
	/* RAJA-TODO
	if r.Room.MaxParticipants > 0 && int(r.Room.MaxParticipants) == len(r.participants) {
		prometheus.ServiceOperationCounter.WithLabelValues("participant_join", "error", "max_exceeded").Add(1)
		return ErrMaxParticipantsExceeded
	}
	RAJA-TODO */

	/* RAJA-TODO: This needs to be a callback
	if r.FirstJoinedAt() == 0 {
		r.joinedAt.Store(time.Now().Unix())
	}
	RAJA-TODO */
	/* RAJA-TODO
	if !participant.Hidden() {
		r.Room.NumParticipants++
	}
	RAJA-TODO */

	// it's important to set this before connection, we don't want to miss out on any publishedTracks
	participant.OnTrackPublished(r.onTrackPublished)
	participant.OnStateChange(func(p types.Participant, oldState livekit.ParticipantInfo_State) {
		r.Logger.Debugw("participant state changed",
			"state", p.State(),
			"participant", p.Identity(),
			"pID", p.ID(),
			"oldState", oldState)
		if r.onParticipantChanged != nil {
			r.onParticipantChanged(participant)
		}
		r.broadcastParticipantState(p, true)

		state := p.State()
		if state == livekit.ParticipantInfo_ACTIVE {
			// subscribe participant to existing publishedTracks
			r.subscribeToExistingTracks(p)

			// start the workers once connectivity is established
			p.Start()

		} else if state == livekit.ParticipantInfo_DISCONNECTED {
			// remove participant from room
			go r.RemoveParticipant(p.Identity())
		}
	})
	// RAJA-TODO: all these need context, Join should take a context argument
	participant.OnTrackUpdated(r.onTrackUpdated)
	participant.OnMetadataUpdate(r.onParticipantMetadataUpdate)
	participant.OnDataPacket(r.onDataPacket)
	r.Logger.Infow("new participant joined",
		"pID", participant.ID(),
		"participant", participant.Identity(),
		"protocol", participant.ProtocolVersion(),
		"room", r.Room.Name,
		"roomID", r.Room.Sid)

	if participant.IsRecorder() && !r.Room.ActiveRecording {
		r.Room.ActiveRecording = true
		r.sendRoomUpdateLocked()
	}

	r.participants[participant.Identity()] = participant
	r.participantOpts[participant.Identity()] = opts

	// gather other participants and send join response
	otherParticipants := make([]*livekit.ParticipantInfo, 0, len(r.participants))
	for _, p := range r.participants {
		if p.ID() != participant.ID() && !p.Hidden() {
			otherParticipants = append(otherParticipants, p.ToProto())
		}
	}

	if r.onParticipantChanged != nil {
		r.onParticipantChanged(participant)
	}

	time.AfterFunc(time.Minute, func() {
		state := participant.State()
		if state == livekit.ParticipantInfo_JOINING || state == livekit.ParticipantInfo_JOINED {
			r.RemoveParticipant(participant.Identity())
		}
	})

	if err := participant.SendJoinResponse(r.Room, otherParticipants, iceServers); err != nil {
		prometheus.ServiceOperationCounter.WithLabelValues("participant_join", "error", "send_response").Add(1)
		return err
	}

	participant.SetMigrateState(types.MigrateStateComplete)

	if participant.SubscriberAsPrimary() {
		// initiates sub connection as primary
		participant.Negotiate()
	}

	prometheus.ServiceOperationCounter.WithLabelValues("participant_join", "success", "").Add(1)

	return nil
}

func (r *LocalRoom) verifyParticipant(p types.Participant) {
	/* RAJA-TODO
	ctx, span := tracer.Start(ctx, "room.verifyParticipant", trace.WithNewRoot())
	defer span.End()
	RAJA-TODO */

	state := p.State()
	if state == livekit.ParticipantInfo_JOINING || state == livekit.ParticipantInfo_JOINED {
		r.RemoveParticipant(p.Identity())
	}
}

func (r *LocalRoom) ResumeParticipant(p types.Participant, responseSink routing.MessageSink) error {
	// close previous sink, and link to new one
	if prevSink := p.GetResponseSink(); prevSink != nil {
		prevSink.Close()
	}
	p.SetResponseSink(responseSink)

	updates := ToProtoParticipants(r.GetParticipants())
	if err := p.SendParticipantUpdate(updates, time.Now()); err != nil {
		return err
	}

	if err := p.ICERestart(); err != nil {
		return err
	}
	return nil
}

func (r *LocalRoom) RemoveParticipant(identity livekit.ParticipantIdentity) {
	r.lock.Lock()
	p, ok := r.participants[identity]
	if ok {
		delete(r.participants, identity)
		delete(r.participantOpts, identity)
		if !p.Hidden() {
			r.Room.NumParticipants--
		}
	}

	activeRecording := false
	if (p != nil && p.IsRecorder()) || p == nil && r.Room.ActiveRecording {
		for _, op := range r.participants {
			if op.IsRecorder() {
				activeRecording = true
				break
			}
		}
	}

	if r.Room.ActiveRecording != activeRecording {
		r.Room.ActiveRecording = activeRecording
		r.sendRoomUpdateLocked()
	}
	r.lock.Unlock()

	if !ok {
		return
	}

	// send broadcast only if it's not already closed
	sendUpdates := p.State() != livekit.ParticipantInfo_DISCONNECTED
	// RAJA-TODO !r.LocalParticipantMigrateOut(p)

	p.OnTrackUpdated(nil)
	p.OnTrackPublished(nil)
	p.OnStateChange(nil)
	p.OnMetadataUpdate(nil)
	p.OnDataPacket(nil)

	// close participant as well
	_ = p.Close()

	r.lock.RLock()
	if len(r.participants) == 0 {
		r.leftAt.Store(time.Now().Unix())
	}
	r.lock.RUnlock()

	if sendUpdates {
		if r.onParticipantChanged != nil {
			r.onParticipantChanged(p)
		}
		r.broadcastParticipantState(p, true)
	}

	// RAJA-TODO: visible participants check
}

func (r *LocalRoom) UpdateSubscriptions(
	participant types.Participant,
	trackIDs []livekit.TrackID,
	participantTracks []*livekit.ParticipantTracks,
	subscribe bool,
) error {
	// RAJA-TODO check canSubscribe

	// find all matching tracks
	trackPublishers := make(map[livekit.TrackID]types.Participant)
	// RAJA-TODO getAllParticipants
	participants := r.GetParticipants()
	for _, trackID := range trackIDs {
		for _, p := range participants {
			track := p.GetPublishedTrack(trackID)
			if track != nil {
				trackPublishers[trackID] = p
				break
			}
		}
	}

	for _, pt := range participantTracks {
		p := r.GetParticipantByID(livekit.ParticipantID(pt.ParticipantSid))
		if p == nil {
			continue
		}
		for _, trackID := range livekit.StringsAsTrackIDs(pt.TrackSids) {
			trackPublishers[trackID] = p
		}
	}

	// handle subscription changes
	for trackID, publisher := range trackPublishers {
		if subscribe {
			if _, err := publisher.AddSubscriber(participant, types.AddSubscriberParams{TrackIDs: []livekit.TrackID{trackID}}); err != nil {
				return err
			}
		} else {
			publisher.RemoveSubscriber(participant, trackID)
		}
	}

	return nil
}

/* RAJA-REMOVE
func (r *Room) SyncState(participant types.Participant, state *livekit.SyncState) error {
	return nil
}
*/

func (r *LocalRoom) UpdateSubscriptionPermissions(participant types.Participant, permissions *livekit.UpdateSubscriptionPermissions) error {
	return participant.UpdateSubscriptionPermissions(permissions, r.GetParticipantByID)
}

func (r *LocalRoom) RemoveDisallowedSubscriptions(sub types.Participant, disallowedSubscriptions map[livekit.TrackID]livekit.ParticipantID) {
	// RAJA-TODO: cloud has to call this and also through remote participants
	for trackID, publisherID := range disallowedSubscriptions {
		pub := r.GetParticipantByID(publisherID)
		if pub == nil {
			continue
		}

		pub.RemoveSubscriber(sub, trackID)
	}
}

func (r *LocalRoom) UpdateVideoLayers(participant types.Participant, updateVideoLayers *livekit.UpdateVideoLayers) error {
	return participant.UpdateVideoLayers(updateVideoLayers)
}

/* RAJA_TODO: can potentially re-used
func (r *LocalRoom) IsClosed() bool {
	select {
	case <-r.closed:
		return true
	default:
		return false
	}
}
*/

/* RAJA-TODO
// CloseIfEmpty closes the room if all participants had left, or it's still empty past timeout
func (r *Room) CloseIfEmpty() {
	if r.IsClosed() {
		return
	}

	r.lock.RLock()
	visibleParticipants := 0
	for _, p := range r.participants {
		if !p.Hidden() {
			visibleParticipants++
		}
	}
	r.lock.RUnlock()

	if visibleParticipants > 0 {
		return
	}

	timeout := r.Room.EmptyTimeout
	var elapsed int64
	if r.FirstJoinedAt() > 0 {
		// exit 20s after
		elapsed = time.Now().Unix() - r.LastLeftAt()
		if timeout > DefaultRoomDepartureGrace {
			timeout = DefaultRoomDepartureGrace
		}
	} else {
		elapsed = time.Now().Unix() - r.Room.CreationTime
	}

	if elapsed >= int64(timeout) {
		r.Close()
	}
}
*/

func (r *LocalRoom) Close() {
	// RAJA-TODO: close all participants, cloud has to close remote participants also
	r.closeOnce.Do(func() {
		close(r.closed)
		r.Logger.Infow("closing room", "roomID", r.Room.Sid, "room", r.Room.Name)
		if r.onClose != nil {
			r.onClose()
		}
	})
}

func (r *Room) OnClose(f func()) {
	r.onClose = f
}

func (r *LocalRoom) OnParticipantChanged(f func(participant types.Participant)) {
	r.onParticipantChanged = f
}

func (r *LocalRoom) SendDataPacket(up *livekit.UserPacket, kind livekit.DataPacket_Kind) {
	dp := &livekit.DataPacket{
		Kind: kind,
		Value: &livekit.DataPacket_User{
			User: up,
		},
	}
	r.onDataPacket(nil, dp)
}

func (r *LocalRoom) SetMetadata(metadata string) {
	r.Room.Metadata = metadata

	r.lock.RLock()
	r.sendRoomUpdateLocked()
	r.lock.RUnlock()

	if r.onMetadataUpdate != nil {
		r.onMetadataUpdate(metadata)
	}
}

func (r *LocalRoom) sendRoomUpdateLocked() {
	// Send update to participants
	for _, p := range r.participants {
		if !p.IsReady() {
			continue
		}

		err := p.SendRoomUpdate(r.Room)
		if err != nil {
			r.Logger.Warnw("failed to send room update", err, "room", r.Room.Name, "participant", p.Identity())
		}
	}
}

func (r *LocalRoom) OnMetadataUpdate(f func(metadata string)) {
	r.onMetadataUpdate = f
}

// checks if participant should be autosubscribed to new tracks, assumes lock is already acquired
func (r *LocalRoom) autoSubscribe(participant types.Participant) bool {
	if !participant.CanSubscribe() {
		return false
	}

	opts := r.participantOpts[participant.Identity()]
	// default to true if no options are set
	if opts != nil && !opts.AutoSubscribe {
		return false
	}
	return true
}

// a ParticipantImpl in the room added a new remoteTrack, subscribe other participants to it
func (r *LocalRoom) onTrackPublished(participant types.Participant, track types.PublishedTrack) {
	// publish participant update, since track state is changed
	// RAJA-TODO: this is done only if migrate state is MigrateStateComplete in cloud
	r.broadcastParticipantState(participant, true)

	// RAJA-TODO: cloud need to handle remote participants too
	r.lock.RLock()
	defer r.lock.RUnlock()

	// subscribe all existing participants to this PublishedTrack
	for _, existingParticipant := range r.participants {
		if existingParticipant == participant {
			// skip publishing participant
			continue
		}
		if existingParticipant.State() != livekit.ParticipantInfo_ACTIVE {
			// not fully joined. don't subscribe yet
			continue
		}
		// RAJA-TODO: there is a resume check in cloud code
		if !r.autoSubscribe(existingParticipant) {
			continue
		}

		r.Logger.Debugw("subscribing to new track",
			"participants", []livekit.ParticipantIdentity{participant.Identity(), existingParticipant.Identity()},
			"pIDs", []livekit.ParticipantID{participant.ID(), existingParticipant.ID()},
			"track", track.ID())
		if _, err := participant.AddSubscriber(existingParticipant, types.AddSubscriberParams{TrackIDs: []livekit.TrackID{track.ID()}}); err != nil {
			r.Logger.Errorw("could not subscribe to remoteTrack", err,
				"participants", []livekit.ParticipantIdentity{participant.Identity(), existingParticipant.Identity()},
				"pIDs", []livekit.ParticipantID{participant.ID(), existingParticipant.ID()},
				"track", track.ID())
		}
	}
	// RAJA-TODO - cloud cleans up RemovePublishedTrack and participants if there are no published tracks

	if r.onParticipantChanged != nil {
		r.onParticipantChanged(participant)
	}
}

func (r *LocalRoom) onTrackUpdated(p types.Participant, _ types.PublishedTrack) {
	// send track updates to everyone, especially if track was updated by admin
	// RAJA-TODO: cloud checks for shoudlBroadcastUpdate
	r.broadcastParticipantState(p, false)
	if r.onParticipantChanged != nil {
		r.onParticipantChanged(p)
	}
	// RAJA-TODO cloud has to send to remote
}

func (r *LocalRoom) onParticipantMetadataUpdate(p types.Participant) {
	// RAJA-TODO: cloud does p.ToProto() as the first argument
	r.broadcastParticipantState(p, false)
	if r.onParticipantChanged != nil {
		r.onParticipantChanged(p)
	}
}

func (r *LocalRoom) onDataPacket(source types.Participant, dp *livekit.DataPacket) {
	// don't forward if source isn't allowed to publish data
	if source != nil && !source.CanPublishData() {
		return
	}
	dest := dp.GetUser().GetDestinationSids()

	for _, op := range r.GetParticipants() {
		if op.State() != livekit.ParticipantInfo_ACTIVE {
			continue
		}
		if source != nil && op.ID() == source.ID() {
			continue
		}
		if len(dest) > 0 {
			found := false
			for _, dID := range dest {
				if op.ID() == livekit.ParticipantID(dID) {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}
		_ = op.SendDataPacket(dp)
	}
}

func (r *LocalRoom) subscribeToExistingTracks(p types.Participant) {
	// RAJA-TODO can subscribe check
	r.lock.RLock()
	shouldSubscribe := r.autoSubscribe(p)
	r.lock.RUnlock()
	if !shouldSubscribe {
		return
	}

	tracksAdded := 0
	// RAJA-TODO: getAllParticipants in the cloud
	for _, op := range r.GetParticipants() {
		if p.ID() == op.ID() {
			// don't send to itself
			continue
		}

		// subscribe to all
		n, err := op.AddSubscriber(p, types.AddSubscriberParams{AllTracks: true})
		if err != nil {
			// TODO: log error? or disconnect?
			r.Logger.Errorw("could not subscribe to participant", err,
				"participants", []livekit.ParticipantIdentity{op.Identity(), p.Identity()},
				"pIDs", []livekit.ParticipantID{op.ID(), p.ID()})
		}
		tracksAdded += n
	}
	if tracksAdded > 0 {
		r.Logger.Debugw("subscribed participants to existing tracks", "tracks", tracksAdded)
	}
}

// broadcast an update about participant p
func (r *LocalRoom) broadcastParticipantState(p types.Participant, skipSource bool) {
	//
	// This is a critical section to ensure that participant update time and
	// the corresponding data are paired properly.
	//
	r.lock.Lock()
	updatedAt := time.Now()
	updates := ToProtoParticipants([]types.Participant{p})
	r.lock.Unlock()

	// RAJA-TODO cloud is checking pi.Hidden and doing some looping
	if p.Hidden() {
		if !skipSource {
			// send update only to hidden participant
			err := p.SendParticipantUpdate(updates, updatedAt)
			if err != nil {
				r.Logger.Errorw("could not send update to participant", err,
					"participant", p.Identity(), "pID", p.ID())
			}
		}
		return
	}

	participants := r.GetParticipants()
	for _, op := range participants {
		// skip itself && closed participants
		if (skipSource && p.ID() == op.ID()) || op.State() == livekit.ParticipantInfo_DISCONNECTED {
			continue
		}

		err := op.SendParticipantUpdate(updates, updatedAt)
		if err != nil {
			r.Logger.Errorw("could not send update to participant", err,
				"participant", p.Identity(), "pID", p.ID())
		}
	}
}

// for protocol 2, send all active speakers
func (r *LocalRoom) sendActiveSpeakers(speakers []*livekit.SpeakerInfo) {
	dp := &livekit.DataPacket{
		Kind: livekit.DataPacket_LOSSY,
		Value: &livekit.DataPacket_Speaker{
			Speaker: &livekit.ActiveSpeakerUpdate{
				Speakers: speakers,
			},
		},
	}

	for _, p := range r.GetParticipants() {
		if p.ProtocolVersion().HandlesDataPackets() && !p.ProtocolVersion().SupportsSpeakerChanged() {
			_ = p.SendDataPacket(dp)
		}
	}
}

// for protocol 3, send only changed updates
func (r *LocalRoom) sendSpeakerChanges(changedSpeakers []*livekit.SpeakerInfo) {
	if len(changedSpeakers) == 0 {
		return
	}

	for _, p := range r.GetParticipants() {
		// only protocol 3 or higher supported, send only changed updates
		if p.ProtocolVersion().SupportsSpeakerChanged() {
			_ = p.SendSpeakerUpdate(changeSpeakers)
		}
	}
}

func (r *LocalRoom) audioUpdateWorker() {
	var smoothValues map[livekit.ParticipantID]float32
	var smoothFactor float32
	var activeThreshold float32
	if ss := r.audioConfig.SmoothIntervals; ss > 1 {
		smoothValues = make(map[livekit.ParticipantID]float32)
		// exponential moving average (EMA), same center of mass with simple moving average (SMA)
		smoothFactor = 2 / float32(ss+1)
		activeThreshold = ConvertAudioLevel(r.audioConfig.ActiveLevel)
	}

	lastActiveMap := make(map[livekit.ParticipantID]*livekit.SpeakerInfo)
	for {
		if r.IsClosed() {
			return
		}

		activeSpeakers := r.GetActiveSpeakers()
		if smoothValues != nil {
			for _, speaker := range activeSpeakers {
				sid := livekit.ParticipantID(speaker.Sid)
				level := smoothValues[sid]
				delete(smoothValues, sid)
				// exponential moving average (EMA)
				level += (speaker.Level - level) * smoothFactor
				speaker.Level = level
			}

			// ensure that previous active speakers are also included
			for sid, level := range smoothValues {
				delete(smoothValues, sid)
				level += -level * smoothFactor
				if level > activeThreshold {
					activeSpeakers = append(activeSpeakers, &livekit.SpeakerInfo{
						Sid:    string(sid),
						Level:  level,
						Active: true,
					})
				}
			}

			// smoothValues map is drained, now repopulate it back
			for _, speaker := range activeSpeakers {
				smoothValues[livekit.ParticipantID(speaker.Sid)] = speaker.Level
			}

			sort.Slice(activeSpeakers, func(i, j int) bool {
				return activeSpeakers[i].Level > activeSpeakers[j].Level
			})
		}

		const invAudioLevelQuantization = 1.0 / AudioLevelQuantization
		for _, speaker := range activeSpeakers {
			speaker.Level = float32(math.Ceil(float64(speaker.Level*AudioLevelQuantization)) * invAudioLevelQuantization)
		}

		changedSpeakers := make([]*livekit.SpeakerInfo, 0, len(activeSpeakers))
		nextActiveMap := make(map[livekit.ParticipantID]*livekit.SpeakerInfo, len(activeSpeakers))
		for _, speaker := range activeSpeakers {
			prev := lastActiveMap[livekit.ParticipantID(speaker.Sid)]
			if prev == nil || prev.Level != speaker.Level {
				changedSpeakers = append(changedSpeakers, speaker)
			}
			nextActiveMap[livekit.ParticipantID(speaker.Sid)] = speaker
		}
		// changedSpeakers need to include previous speakers that are no longer speaking
		for sid, speaker := range lastActiveMap {
			if nextActiveMap[sid] == nil {
				speaker.Level = 0
				speaker.Active = false
				changedSpeakers = append(changedSpeakers, speaker)
			}
		}

		// see if an update is needed
		if len(changedSpeakers) > 0 {
			r.sendActiveSpeakers(activeSpeakers)
			r.sendSpeakerChanges(changedSpeakers)
		}

		lastActiveMap = nextActiveMap

		time.Sleep(time.Duration(r.audioConfig.UpdateInterval) * time.Millisecond)
	}
}

func (r *LocalRoom) connectionQualityWorker() {
	// send updates to only users that are subscribed to each other
	for {
		if r.IsClosed() {
			return
		}

		participants := r.GetParticipants()
		// RAJA-TODO: cloud has to send remote updates
		connectionInfos := make(map[livekit.ParticipantID]*livekit.ConnectionQualityInfo, len(participants))

		for _, p := range participants {
			connectionInfos[p.ID()] = p.GetConnectionQuality()
		}

		for _, op := range participants {
			if !op.ProtocolVersion().SupportsConnectionQuality() {
				continue
			}
			update := &livekit.ConnectionQualityUpdate{}

			// send to user itself
			if info, ok := connectionInfos[op.ID()]; ok {
				update.Updates = append(update.Updates, info)
			}

			// add connection quality of other participants its subscribed to
			for _, sid := range op.GetSubscribedParticipants() {
				if info, ok := connectionInfos[sid]; ok {
					update.Updates = append(update.Updates, info)
				}
			}
			if err := op.SendConnectionQualityUpdate(update); err != nil {
				r.Logger.Warnw("could not send connection quality update", err,
					"participant", op.Identity())
			}
		}

		time.Sleep(time.Second * 5)
	}
}

func (r *LocalRoom) DebugInfo() map[string]interface{} {
	info := map[string]interface{}{
		"Name":      r.Room.Name,
		"Sid":       r.Room.Sid,
		"CreatedAt": r.Room.CreationTime,
	}

	participants := r.GetParticipants()
	participantInfo := make(map[string]interface{})
	for _, p := range participants {
		participantInfo[string(p.Identity())] = p.DebugInfo()
	}
	info["Participants"] = participantInfo

	return info
}
