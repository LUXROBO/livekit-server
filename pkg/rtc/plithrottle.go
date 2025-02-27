package rtc

import (
	"sync"
	"time"

	"github.com/livekit/livekit-server/pkg/config"
	"github.com/livekit/livekit-server/pkg/sfu"
)

type pliThrottle struct {
	config   config.PLIThrottleConfig
	mu       sync.RWMutex
	periods  map[uint32]int64
	lastSent map[uint32]int64
}

func newPLIThrottle(conf config.PLIThrottleConfig) *pliThrottle {
	return &pliThrottle{
		config:   conf,
		periods:  make(map[uint32]int64),
		lastSent: make(map[uint32]int64),
	}
}

func (t *pliThrottle) addTrack(ssrc uint32, rid string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	var duration time.Duration
	switch rid {
	case sfu.FullResolution:
		duration = t.config.HighQuality
	case sfu.HalfResolution:
		duration = t.config.MidQuality
	case sfu.QuarterResolution:
		duration = t.config.LowQuality
	default:
		duration = t.config.MidQuality
	}

	t.periods[ssrc] = duration.Nanoseconds()
}

func (t *pliThrottle) canSend(ssrc uint32) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	if period, ok := t.periods[ssrc]; ok {
		if n := time.Now().UnixNano(); n-t.lastSent[ssrc] > period {
			t.lastSent[ssrc] = n
			return true
		} else {
			return false
		}
	}
	return true
}
