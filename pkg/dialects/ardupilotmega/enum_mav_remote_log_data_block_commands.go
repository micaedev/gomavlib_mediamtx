//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

// Special ACK block numbers control activation of dataflash log streaming.
type MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS int

const (
	// UAV to stop sending DataFlash blocks.
	MAV_REMOTE_LOG_DATA_BLOCK_STOP MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS = 2147483645
	// UAV to start sending DataFlash blocks.
	MAV_REMOTE_LOG_DATA_BLOCK_START MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS = 2147483646
)

var labels_MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS = map[MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS = map[string]MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS) String() string {
	if l, ok := labels_MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS[e]; ok {
		return l
	}
	return "invalid value"
}
