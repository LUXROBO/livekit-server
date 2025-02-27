package rtc

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/livekit/livekit-server/pkg/sfu/connectionquality"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"github.com/pion/rtcp"
	"github.com/pion/webrtc/v3"

	"github.com/livekit/livekit-server/pkg/config"
	"github.com/livekit/livekit-server/pkg/sfu"
	"github.com/livekit/livekit-server/pkg/sfu/buffer"
	"github.com/livekit/livekit-server/pkg/sfu/twcc"
	"github.com/livekit/livekit-server/pkg/telemetry"
)

const (
	upLostUpdateDelta               = time.Second
	connectionQualityUpdateInterval = 5 * time.Second
)

// MediaTrack represents a WebRTC track that needs to be forwarded
// Implements MediaTrack and PublishedTrack interface
type MediaTrack struct {
	params      MediaTrackParams
	numUpTracks uint32
	buffer      *buffer.Buffer

	layerSsrcs [livekit.VideoQuality_HIGH + 1]uint32

	audioLevelMu sync.RWMutex
	audioLevel   *AudioLevel

	statsLock         sync.Mutex
	currentUpFracLost uint32
	maxUpFracLost     uint8
	maxUpFracLostTs   time.Time
	connectionStats   *connectionquality.ConnectionStats

	done chan struct{}

	*MediaTrackReceiver

	onMediaLossUpdate func(trackID livekit.TrackID, fractionalLoss uint32)
}

type MediaTrackParams struct {
	TrackInfo           *livekit.TrackInfo
	SignalCid           string
	SdpCid              string
	ParticipantID       livekit.ParticipantID
	ParticipantIdentity livekit.ParticipantIdentity
	// channel to send RTCP packets to the source
	RTCPChan         chan []rtcp.Packet
	BufferFactory    *buffer.Factory
	ReceiverConfig   ReceiverConfig
	SubscriberConfig DirectionConfig
	AudioConfig      config.AudioConfig
	Telemetry        telemetry.TelemetryService
	Logger           logger.Logger
}

func NewMediaTrack(track *webrtc.TrackRemote, params MediaTrackParams) *MediaTrack {
	t := &MediaTrack{
		params:          params,
		connectionStats: connectionquality.NewConnectionStats(),
		done:            make(chan struct{}),
	}

	t.MediaTrackReceiver = NewMediaTrackReceiver(MediaTrackReceiverParams{
		TrackInfo:           params.TrackInfo,
		MediaTrack:          t,
		ParticipantID:       params.ParticipantID,
		ParticipantIdentity: params.ParticipantIdentity,
		BufferFactory:       params.BufferFactory,
		ReceiverConfig:      params.ReceiverConfig,
		SubscriberConfig:    params.SubscriberConfig,
		Telemetry:           params.Telemetry,
		Logger:              params.Logger,
	})
	t.MediaTrackReceiver.OnMediaLossUpdate(func(fractionalLoss uint8) {
		if t.buffer != nil {
			// ok to access buffer since receivers are added before subscribers
			t.buffer.SetLastFractionLostReport(fractionalLoss)
		}
	})

	// on close signal via closing channel to workers
	t.AddOnClose(t.closeChan)
	go t.updateStats()

	return t
}

func (t *MediaTrack) OnMediaLossUpdate(f func(trackID livekit.TrackID, fractionalLoss uint32)) {
	t.onMediaLossUpdate = f
}

func (t *MediaTrack) SignalCid() string {
	return t.params.SignalCid
}

func (t *MediaTrack) SdpCid() string {
	return t.params.SdpCid
}

func (t *MediaTrack) ToProto() *livekit.TrackInfo {
	info := t.MediaTrackReceiver.TrackInfo()
	info.Muted = t.IsMuted()
	info.Simulcast = t.IsSimulcast()
	layers := t.MediaTrackReceiver.GetVideoLayers()
	for _, layer := range layers {
		if int(layer.Quality) < len(t.layerSsrcs) {
			layer.Ssrc = t.layerSsrcs[layer.Quality]
		}
	}
	info.Layers = layers

	return info
}

func (t *MediaTrack) publishLossPercentage() uint32 {
	return FixedPointToPercent(uint8(atomic.LoadUint32(&t.currentUpFracLost)))
}

func (t *MediaTrack) getNumUpTracks() (uint32, uint32) {
	numExpected := atomic.LoadUint32(&t.numUpTracks)

	numSubscribedLayers := t.numSubscribedLayers()
	if numSubscribedLayers < numExpected {
		numExpected = numSubscribedLayers
	}

	numPublishing := uint32(0)
	receiver := t.Receiver()
	if receiver != nil {
		numPublishing = uint32(receiver.(sfu.Receiver).NumAvailableSpatialLayers())
	}

	return numPublishing, numExpected
}

// AddReceiver adds a new RTP receiver to the track
func (t *MediaTrack) AddReceiver(receiver *webrtc.RTPReceiver, track *webrtc.TrackRemote, twcc *twcc.Responder) {
	buff, rtcpReader := t.params.BufferFactory.GetBufferPair(uint32(track.SSRC()))
	if buff == nil || rtcpReader == nil {
		logger.Errorw("could not retrieve buffer pair", nil,
			"track", t.ID())
		return
	}
	buff.OnFeedback(t.handlePublisherFeedback)

	if t.Kind() == livekit.TrackType_AUDIO {
		t.audioLevelMu.Lock()
		t.audioLevel = NewAudioLevel(t.params.AudioConfig.ActiveLevel, t.params.AudioConfig.MinPercentile)
		buff.OnAudioLevel(func(level uint8, duration uint32) {
			t.audioLevelMu.RLock()
			defer t.audioLevelMu.RUnlock()

			t.audioLevel.Observe(level, duration)
		})
		t.audioLevelMu.Unlock()
	} else if t.Kind() == livekit.TrackType_VIDEO {
		if twcc != nil {
			buff.OnTransportWideCC(func(sn uint16, timeNS int64, marker bool) {
				twcc.Push(sn, timeNS, marker)
			})
		}
	}

	rtcpReader.OnPacket(func(bytes []byte) {
		pkts, err := rtcp.Unmarshal(bytes)
		if err != nil {
			t.params.Logger.Errorw("could not unmarshal RTCP", err)
			return
		}

		for _, pkt := range pkts {
			switch pkt := pkt.(type) {
			case *rtcp.SourceDescription:
			// do nothing for now
			case *rtcp.SenderReport:
				buff.SetSenderReportData(pkt.RTPTime, pkt.NTPTime)
			}
		}
	})

	if t.Receiver() == nil {
		wr := sfu.NewWebRTCReceiver(
			receiver,
			track,
			t.PublisherID(),
			sfu.WithPliThrottle(0),
			sfu.WithLoadBalanceThreshold(20),
			sfu.WithStreamTrackers(),
		)
		wr.SetRTCPCh(t.params.RTCPChan)
		wr.OnCloseHandler(func() {
			t.stopMaxQualityTimer()
			t.RemoveAllSubscribers()
			t.MediaTrackReceiver.Close()
			t.params.Telemetry.TrackUnpublished(context.Background(), t.PublisherID(), t.ToProto(), uint32(track.SSRC()))
		})
		t.params.Telemetry.TrackPublished(context.Background(), t.PublisherID(), t.ToProto())
		if t.Kind() == livekit.TrackType_AUDIO {
			t.buffer = buff
		}

		t.MediaTrackReceiver.SetupReceiver(wr)
		t.startMaxQualityTimer()
	}

	t.Receiver().(sfu.Receiver).AddUpTrack(track, buff)
	t.params.Telemetry.AddUpTrack(t.PublisherID(), t.ID(), buff)

	atomic.AddUint32(&t.numUpTracks, 1)
	// LK-TODO: can remove this completely when VideoLayers protocol becomes the default as it has info from client or if we decide to use TrackInfo.Simulcast
	if atomic.LoadUint32(&t.numUpTracks) > 1 || track.RID() != "" {
		// cannot only rely on numUpTracks since we fire metadata events immediately after the first layer
		t.MediaTrackReceiver.SetSimulcast(true)
	}

	if t.IsSimulcast() {
		layer := sfu.RidToLayer(track.RID())
		if int(layer) < len(t.layerSsrcs) {
			t.layerSsrcs[layer] = uint32(track.SSRC())
		}
	}

	buff.Bind(receiver.GetParameters(), track.Codec().RTPCodecCapability, buffer.Options{
		MaxBitRate: t.params.ReceiverConfig.maxBitrate,
	})
}

func (t *MediaTrack) TrySetSimulcastSSRC(layer uint8, ssrc uint32) {
	if int(layer) < len(t.layerSsrcs) && t.layerSsrcs[layer] == 0 {
		t.layerSsrcs[layer] = ssrc
	}
}

func (t *MediaTrack) GetAudioLevel() (level uint8, active bool) {
	t.audioLevelMu.RLock()
	defer t.audioLevelMu.RUnlock()

	if t.audioLevel == nil {
		return SilentAudioLevel, false
	}
	return t.audioLevel.GetLevel()
}

func (t *MediaTrack) handlePublisherFeedback(packets []rtcp.Packet) {
	var maxLost uint8
	var hasReport bool
	var delay uint32
	var jitter uint32
	var totalLost uint32
	var maxSeqNum uint32

	//forward to telemetry
	t.params.Telemetry.HandleRTCP(livekit.StreamType_UPSTREAM, t.params.ParticipantID, t.ID(), packets)

	for _, p := range packets {
		switch pkt := p.(type) {
		// sfu.Buffer generates ReceiverReports for the publisher
		case *rtcp.ReceiverReport:
			for _, rr := range pkt.Reports {
				if rr.FractionLost > maxLost {
					maxLost = rr.FractionLost
				}

				if rr.Delay > delay {
					delay = rr.Delay
				}
				if rr.Jitter > jitter {
					jitter = rr.Jitter
				}
				if rr.LastSequenceNumber > maxSeqNum {
					maxSeqNum = rr.LastSequenceNumber
				}

				totalLost = rr.TotalLost

				hasReport = true
			}
		}
	}

	if hasReport {
		t.statsLock.Lock()
		if maxLost > t.maxUpFracLost {
			t.maxUpFracLost = maxLost
		}

		now := time.Now()
		if now.Sub(t.maxUpFracLostTs) > upLostUpdateDelta {
			atomic.StoreUint32(&t.currentUpFracLost, uint32(t.maxUpFracLost))
			t.maxUpFracLost = 0
			t.maxUpFracLostTs = now
		}
		// update feedback stats
		current := t.connectionStats.Curr
		if jitter > current.Jitter {
			current.Jitter = jitter
		}
		if delay > current.Delay {
			current.Delay = delay
		}
		if maxSeqNum > current.LastSeqNum {
			current.LastSeqNum = maxSeqNum
		}
		current.PacketsLost = totalLost
		t.statsLock.Unlock()
	}

	// also look for sender reports
	// feedback for the source RTCP
	t.params.RTCPChan <- packets
}

func (t *MediaTrack) GetConnectionScore() float64 {
	t.statsLock.Lock()
	defer t.statsLock.Unlock()
	return t.connectionStats.Score
}

func (t *MediaTrack) closeChan() {
	close(t.done)
}

func (t *MediaTrack) updateStats() {
	for {
		select {
		case <-t.done:
			return
		case <-time.After(connectionQualityUpdateInterval):
			t.statsLock.Lock()
			if t.Kind() == livekit.TrackType_AUDIO {
				t.connectionStats.CalculateAudioScore()
			} else {
				t.calculateVideoScore()
			}
			t.statsLock.Unlock()
		}
	}
}

func (t *MediaTrack) calculateVideoScore() {
	var reducedQuality bool
	publishing, expected := t.getNumUpTracks()
	if publishing < expected {
		reducedQuality = true
	}

	loss := t.publishLossPercentage()
	if expected == 0 {
		loss = 0
	}
	t.connectionStats.Score = connectionquality.Loss2Score(loss, reducedQuality)
}
