//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package standard

import (
	"errors"
)

// These values encode the bit positions of the decode position. These values can be used to read the value of a flag bit by combining the base_mode variable with AND with the flag position value. The result will be either 0 or 1, depending on if the flag is set or not.
type MAV_MODE_FLAG_DECODE_POSITION int

const (
	// First bit:  10000000
	MAV_MODE_FLAG_DECODE_POSITION_SAFETY MAV_MODE_FLAG_DECODE_POSITION = 128
	// Second bit: 01000000
	MAV_MODE_FLAG_DECODE_POSITION_MANUAL MAV_MODE_FLAG_DECODE_POSITION = 64
	// Third bit:  00100000
	MAV_MODE_FLAG_DECODE_POSITION_HIL MAV_MODE_FLAG_DECODE_POSITION = 32
	// Fourth bit: 00010000
	MAV_MODE_FLAG_DECODE_POSITION_STABILIZE MAV_MODE_FLAG_DECODE_POSITION = 16
	// Fifth bit:  00001000
	MAV_MODE_FLAG_DECODE_POSITION_GUIDED MAV_MODE_FLAG_DECODE_POSITION = 8
	// Sixth bit:   00000100
	MAV_MODE_FLAG_DECODE_POSITION_AUTO MAV_MODE_FLAG_DECODE_POSITION = 4
	// Seventh bit: 00000010
	MAV_MODE_FLAG_DECODE_POSITION_TEST MAV_MODE_FLAG_DECODE_POSITION = 2
	// Eighth bit: 00000001
	MAV_MODE_FLAG_DECODE_POSITION_CUSTOM_MODE MAV_MODE_FLAG_DECODE_POSITION = 1
)

var labels_MAV_MODE_FLAG_DECODE_POSITION = map[MAV_MODE_FLAG_DECODE_POSITION]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_MODE_FLAG_DECODE_POSITION) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_MODE_FLAG_DECODE_POSITION[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_MODE_FLAG_DECODE_POSITION = map[string]MAV_MODE_FLAG_DECODE_POSITION{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_MODE_FLAG_DECODE_POSITION) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_MODE_FLAG_DECODE_POSITION[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_MODE_FLAG_DECODE_POSITION) String() string {
	if l, ok := labels_MAV_MODE_FLAG_DECODE_POSITION[e]; ok {
		return l
	}
	return "invalid value"
}
