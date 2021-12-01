//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package uavionix

import (
	"errors"
)

// These flags encode the MAV mode.
type MAV_MODE_FLAG int

const (
	// 0b10000000 MAV safety set to armed. Motors are enabled / running / can start. Ready to fly. Additional note: this flag is to be ignore when sent in the command MAV_CMD_DO_SET_MODE and MAV_CMD_COMPONENT_ARM_DISARM shall be used instead. The flag can still be used to report the armed state.
	MAV_MODE_FLAG_SAFETY_ARMED MAV_MODE_FLAG = 128
	// 0b01000000 remote control input is enabled.
	MAV_MODE_FLAG_MANUAL_INPUT_ENABLED MAV_MODE_FLAG = 64
	// 0b00100000 hardware in the loop simulation. All motors / actuators are blocked, but internal software is full operational.
	MAV_MODE_FLAG_HIL_ENABLED MAV_MODE_FLAG = 32
	// 0b00010000 system stabilizes electronically its attitude (and optionally position). It needs however further control inputs to move around.
	MAV_MODE_FLAG_STABILIZE_ENABLED MAV_MODE_FLAG = 16
	// 0b00001000 guided mode enabled, system flies waypoints / mission items.
	MAV_MODE_FLAG_GUIDED_ENABLED MAV_MODE_FLAG = 8
	// 0b00000100 autonomous mode enabled, system finds its own goal positions. Guided flag can be set or not, depends on the actual implementation.
	MAV_MODE_FLAG_AUTO_ENABLED MAV_MODE_FLAG = 4
	// 0b00000010 system has a test mode enabled. This flag is intended for temporary system tests and should not be used for stable implementations.
	MAV_MODE_FLAG_TEST_ENABLED MAV_MODE_FLAG = 2
	// 0b00000001 Reserved for future use.
	MAV_MODE_FLAG_CUSTOM_MODE_ENABLED MAV_MODE_FLAG = 1
)

var labels_MAV_MODE_FLAG = map[MAV_MODE_FLAG]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_MODE_FLAG) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_MODE_FLAG[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_MODE_FLAG = map[string]MAV_MODE_FLAG{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_MODE_FLAG) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_MODE_FLAG[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_MODE_FLAG) String() string {
	if l, ok := labels_MAV_MODE_FLAG[e]; ok {
		return l
	}
	return "invalid value"
}
