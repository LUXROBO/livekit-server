package telemetry

import (
	"context"

	"github.com/gammazero/workerpool"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/webhook"
	"github.com/pion/rtcp"

	"github.com/livekit/livekit-server/pkg/sfu/buffer"
	"github.com/livekit/livekit-server/pkg/telemetry/prometheus"
)

type TelemetryServiceInternal interface {
	TelemetryService
	SendAnalytics()
}

type TelemetryReporter interface {
	Report(ctx context.Context, stats []*livekit.AnalyticsStat)
}

type telemetryServiceInternal struct {
	notifier    webhook.Notifier
	webhookPool *workerpool.WorkerPool

	// one worker per participant
	workers map[livekit.ParticipantID]*StatsWorker

	analytics AnalyticsService
}

func NewTelemetryServiceInternal(notifier webhook.Notifier, analytics AnalyticsService) TelemetryServiceInternal {
	return &telemetryServiceInternal{
		notifier:    notifier,
		webhookPool: workerpool.New(1),
		workers:     make(map[livekit.ParticipantID]*StatsWorker),
		analytics:   analytics,
	}
}

func (t *telemetryServiceInternal) AddUpTrack(participantID livekit.ParticipantID, trackID livekit.TrackID, buff *buffer.Buffer) {
	w := t.workers[participantID]
	if w != nil {
		w.AddBuffer(trackID, buff)
	}
}

func (t *telemetryServiceInternal) OnDownstreamPacket(participantID livekit.ParticipantID, trackID livekit.TrackID, bytes int) {
	w := t.workers[participantID]
	if w != nil {
		w.OnDownstreamPacket(trackID, bytes)
	}
}

func (t *telemetryServiceInternal) HandleRTCP(streamType livekit.StreamType, participantID livekit.ParticipantID, trackID livekit.TrackID, pkts []rtcp.Packet) {
	stats := &livekit.AnalyticsStat{}
	for _, pkt := range pkts {
		switch pkt := pkt.(type) {
		case *rtcp.TransportLayerNack:
			stats.NackCount++
		case *rtcp.PictureLossIndication:
			stats.PliCount++
		case *rtcp.FullIntraRequest:
			stats.FirCount++
		case *rtcp.ReceiverReport:
			for _, rr := range pkt.Reports {
				if jitter := float64(rr.Jitter); jitter > stats.Jitter {
					stats.Jitter = jitter
				}
				if totalLost := uint64(rr.TotalLost); totalLost > stats.PacketLost {
					stats.PacketLost = totalLost
				}
			}
			if streamType == livekit.StreamType_DOWNSTREAM {
				rtt := GetRttMs(pkt)
				if rtt >= 0 {
					stats.Rtt = uint32(rtt)
				}
			}
		}
	}

	direction := prometheus.Incoming
	if streamType == livekit.StreamType_DOWNSTREAM {
		direction = prometheus.Outgoing
	}

	prometheus.IncrementRTCP(direction, stats.NackCount, stats.PliCount, stats.FirCount)

	w := t.workers[participantID]
	if w != nil {
		w.OnRTCP(trackID, streamType, stats)
	}
}

func (t *telemetryServiceInternal) Report(ctx context.Context, stats []*livekit.AnalyticsStat) {
	for _, stat := range stats {
		direction := prometheus.Incoming
		if stat.Kind == livekit.StreamType_DOWNSTREAM {
			direction = prometheus.Outgoing
		}

		prometheus.IncrementPackets(direction, stat.TotalPackets)
		prometheus.IncrementBytes(direction, stat.TotalBytes)
	}

	t.analytics.SendStats(ctx, stats)
}

func (t *telemetryServiceInternal) SendAnalytics() {
	for _, worker := range t.workers {
		worker.Update()
	}
}
