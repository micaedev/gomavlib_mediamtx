//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ualberta

import (
	"errors"
)

// A data stream is not a fixed set of messages, but rather a
// recommendation to the autopilot software. Individual autopilots may or may not obey
// the recommended messages.
type MAV_DATA_STREAM uint32

const (
	// Enable all data streams
	MAV_DATA_STREAM_ALL MAV_DATA_STREAM = 0
	// Enable IMU_RAW, GPS_RAW, GPS_STATUS packets.
	MAV_DATA_STREAM_RAW_SENSORS MAV_DATA_STREAM = 1
	// Enable GPS_STATUS, CONTROL_STATUS, AUX_STATUS
	MAV_DATA_STREAM_EXTENDED_STATUS MAV_DATA_STREAM = 2
	// Enable RC_CHANNELS_SCALED, RC_CHANNELS_RAW, SERVO_OUTPUT_RAW
	MAV_DATA_STREAM_RC_CHANNELS MAV_DATA_STREAM = 3
	// Enable ATTITUDE_CONTROLLER_OUTPUT, POSITION_CONTROLLER_OUTPUT, NAV_CONTROLLER_OUTPUT.
	MAV_DATA_STREAM_RAW_CONTROLLER MAV_DATA_STREAM = 4
	// Enable LOCAL_POSITION, GLOBAL_POSITION_INT messages.
	MAV_DATA_STREAM_POSITION MAV_DATA_STREAM = 6
	// Dependent on the autopilot
	MAV_DATA_STREAM_EXTRA1 MAV_DATA_STREAM = 10
	// Dependent on the autopilot
	MAV_DATA_STREAM_EXTRA2 MAV_DATA_STREAM = 11
	// Dependent on the autopilot
	MAV_DATA_STREAM_EXTRA3 MAV_DATA_STREAM = 12
)

var labels_MAV_DATA_STREAM = map[MAV_DATA_STREAM]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_DATA_STREAM) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_DATA_STREAM[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_DATA_STREAM = map[string]MAV_DATA_STREAM{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_DATA_STREAM) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_DATA_STREAM[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_DATA_STREAM) String() string {
	if l, ok := labels_MAV_DATA_STREAM[e]; ok {
		return l
	}
	return "invalid value"
}
