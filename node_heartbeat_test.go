package gomavlib

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/aler9/gomavlib/pkg/dialect"
	"github.com/aler9/gomavlib/pkg/message"
)

func TestNodeHeartbeat(t *testing.T) {
	dialect := &dialect.Dialect{
		Version:  3,
		Messages: []message.Message{&MessageHeartbeat{}},
	}

	node1, err := NewNode(NodeConf{
		Dialect:     dialect,
		OutVersion:  V2,
		OutSystemID: 10,
		Endpoints: []EndpointConf{
			EndpointUDPServer{"127.0.0.1:5600"},
		},
		HeartbeatDisable: true,
	})
	require.NoError(t, err)
	defer node1.Close()

	node2, err := NewNode(NodeConf{
		Dialect:     dialect,
		OutVersion:  V2,
		OutSystemID: 11,
		Endpoints: []EndpointConf{
			EndpointUDPClient{"127.0.0.1:5600"},
		},
		HeartbeatDisable: false,
		HeartbeatPeriod:  500 * time.Millisecond,
	})
	require.NoError(t, err)
	defer node2.Close()

	<-node1.Events()
	evt := <-node1.Events()
	fr, ok := evt.(*EventFrame)
	require.Equal(t, true, ok)
	_, ok = fr.Message().(*MessageHeartbeat)
	require.Equal(t, true, ok)
}
