//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package all

import (
	"errors"
)

type MAV_MODE_GIMBAL uint32

const (
	// Gimbal is powered on but has not started initializing yet.
	MAV_MODE_GIMBAL_UNINITIALIZED MAV_MODE_GIMBAL = 0
	// Gimbal is currently running calibration on the pitch axis.
	MAV_MODE_GIMBAL_CALIBRATING_PITCH MAV_MODE_GIMBAL = 1
	// Gimbal is currently running calibration on the roll axis.
	MAV_MODE_GIMBAL_CALIBRATING_ROLL MAV_MODE_GIMBAL = 2
	// Gimbal is currently running calibration on the yaw axis.
	MAV_MODE_GIMBAL_CALIBRATING_YAW MAV_MODE_GIMBAL = 3
	// Gimbal has finished calibrating and initializing, but is relaxed pending reception of first rate command from copter.
	MAV_MODE_GIMBAL_INITIALIZED MAV_MODE_GIMBAL = 4
	// Gimbal is actively stabilizing.
	MAV_MODE_GIMBAL_ACTIVE MAV_MODE_GIMBAL = 5
	// Gimbal is relaxed because it missed more than 10 expected rate command messages in a row. Gimbal will move back to active mode when it receives a new rate command.
	MAV_MODE_GIMBAL_RATE_CMD_TIMEOUT MAV_MODE_GIMBAL = 6
)

var labels_MAV_MODE_GIMBAL = map[MAV_MODE_GIMBAL]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_MODE_GIMBAL) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_MODE_GIMBAL[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_MODE_GIMBAL = map[string]MAV_MODE_GIMBAL{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_MODE_GIMBAL) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_MODE_GIMBAL[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_MODE_GIMBAL) String() string {
	if l, ok := labels_MAV_MODE_GIMBAL[e]; ok {
		return l
	}
	return "invalid value"
}
