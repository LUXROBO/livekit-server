package selector_test

import (
	"testing"
	"time"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/utils"
	"github.com/stretchr/testify/require"

	"github.com/livekit/livekit-server/pkg/config"
	"github.com/livekit/livekit-server/pkg/routing/selector"
)

const (
	loadLimit     = 0.5
	regionWest    = "us-west"
	regionEast    = "us-east"
	regionSeattle = "seattle"
)

func TestRegionAwareRouting(t *testing.T) {
	rc := []config.RegionConfig{
		{
			Name: regionWest,
			Lat:  37.64046607830567,
			Lon:  -120.88026233189062,
		},
		{
			Name: regionEast,
			Lat:  40.68914362140307,
			Lon:  -74.04445748616385,
		},
		{
			Name: regionSeattle,
			Lat:  47.620426730945454,
			Lon:  -122.34938468973702,
		},
	}
	t.Run("works without region config", func(t *testing.T) {
		nodes := []*livekit.Node{
			newTestNodeInRegion("", false),
		}
		s, err := selector.NewRegionAwareSelector(regionEast, nil)
		require.NoError(t, err)

		node, err := s.SelectNode(nodes)
		require.NoError(t, err)
		require.NotNil(t, node)
	})

	t.Run("picks available nodes in same region", func(t *testing.T) {
		expectedNode := newTestNodeInRegion(regionEast, true)
		nodes := []*livekit.Node{
			newTestNodeInRegion(regionSeattle, true),
			newTestNodeInRegion(regionWest, true),
			expectedNode,
			newTestNodeInRegion(regionEast, false),
		}
		s, err := selector.NewRegionAwareSelector(regionEast, rc)
		require.NoError(t, err)
		s.SysloadLimit = loadLimit

		node, err := s.SelectNode(nodes)
		require.NoError(t, err)
		require.Equal(t, expectedNode, node)
	})

	t.Run("picks available nodes in same region when current node is first in the list", func(t *testing.T) {
		expectedNode := newTestNodeInRegion(regionEast, true)
		nodes := []*livekit.Node{
			expectedNode,
			newTestNodeInRegion(regionSeattle, true),
			newTestNodeInRegion(regionWest, true),
			newTestNodeInRegion(regionEast, false),
		}
		s, err := selector.NewRegionAwareSelector(regionEast, rc)
		require.NoError(t, err)
		s.SysloadLimit = loadLimit

		node, err := s.SelectNode(nodes)
		require.NoError(t, err)
		require.Equal(t, expectedNode, node)
	})

	t.Run("picks closest node in a diff region", func(t *testing.T) {
		expectedNode := newTestNodeInRegion(regionWest, true)
		nodes := []*livekit.Node{
			newTestNodeInRegion(regionSeattle, false),
			expectedNode,
			newTestNodeInRegion(regionEast, true),
		}
		s, err := selector.NewRegionAwareSelector(regionSeattle, rc)
		require.NoError(t, err)
		s.SysloadLimit = loadLimit

		node, err := s.SelectNode(nodes)
		require.NoError(t, err)
		require.Equal(t, expectedNode, node)
	})

	t.Run("handles multiple nodes in same region", func(t *testing.T) {
		expectedNode := newTestNodeInRegion(regionWest, true)
		nodes := []*livekit.Node{
			newTestNodeInRegion(regionSeattle, false),
			newTestNodeInRegion(regionEast, true),
			newTestNodeInRegion(regionEast, true),
			expectedNode,
			expectedNode,
		}
		s, err := selector.NewRegionAwareSelector(regionSeattle, rc)
		require.NoError(t, err)
		s.SysloadLimit = loadLimit

		node, err := s.SelectNode(nodes)
		require.NoError(t, err)
		require.Equal(t, expectedNode, node)
	})

	t.Run("functions when current region is full", func(t *testing.T) {
		nodes := []*livekit.Node{
			newTestNodeInRegion(regionWest, true),
		}
		s, err := selector.NewRegionAwareSelector(regionEast, rc)
		require.NoError(t, err)

		node, err := s.SelectNode(nodes)
		require.NoError(t, err)
		require.NotNil(t, node)
	})
}

func newTestNodeInRegion(region string, available bool) *livekit.Node {
	load := float32(0.4)
	if !available {
		load = 1.0
	}
	return &livekit.Node{
		Id:     utils.NewGuid(utils.NodePrefix),
		Region: region,
		State:  livekit.NodeState_SERVING,
		Stats: &livekit.NodeStats{
			UpdatedAt:       time.Now().Unix(),
			NumCpus:         1,
			LoadAvgLast1Min: load,
		},
	}
}
