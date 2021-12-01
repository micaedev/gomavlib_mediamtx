//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package paparazzi

import (
	"errors"
)

// Bitmap to indicate which dimensions should be ignored by the vehicle: a value of 0b0000000000000000 or 0b0000001000000000 indicates that none of the setpoint dimensions should be ignored. If bit 9 is set the floats afx afy afz should be interpreted as force instead of acceleration.
type POSITION_TARGET_TYPEMASK int

const (
	// Ignore position x
	POSITION_TARGET_TYPEMASK_X_IGNORE POSITION_TARGET_TYPEMASK = 1
	// Ignore position y
	POSITION_TARGET_TYPEMASK_Y_IGNORE POSITION_TARGET_TYPEMASK = 2
	// Ignore position z
	POSITION_TARGET_TYPEMASK_Z_IGNORE POSITION_TARGET_TYPEMASK = 4
	// Ignore velocity x
	POSITION_TARGET_TYPEMASK_VX_IGNORE POSITION_TARGET_TYPEMASK = 8
	// Ignore velocity y
	POSITION_TARGET_TYPEMASK_VY_IGNORE POSITION_TARGET_TYPEMASK = 16
	// Ignore velocity z
	POSITION_TARGET_TYPEMASK_VZ_IGNORE POSITION_TARGET_TYPEMASK = 32
	// Ignore acceleration x
	POSITION_TARGET_TYPEMASK_AX_IGNORE POSITION_TARGET_TYPEMASK = 64
	// Ignore acceleration y
	POSITION_TARGET_TYPEMASK_AY_IGNORE POSITION_TARGET_TYPEMASK = 128
	// Ignore acceleration z
	POSITION_TARGET_TYPEMASK_AZ_IGNORE POSITION_TARGET_TYPEMASK = 256
	// Use force instead of acceleration
	POSITION_TARGET_TYPEMASK_FORCE_SET POSITION_TARGET_TYPEMASK = 512
	// Ignore yaw
	POSITION_TARGET_TYPEMASK_YAW_IGNORE POSITION_TARGET_TYPEMASK = 1024
	// Ignore yaw rate
	POSITION_TARGET_TYPEMASK_YAW_RATE_IGNORE POSITION_TARGET_TYPEMASK = 2048
)

var labels_POSITION_TARGET_TYPEMASK = map[POSITION_TARGET_TYPEMASK]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e POSITION_TARGET_TYPEMASK) MarshalText() ([]byte, error) {
	if l, ok := labels_POSITION_TARGET_TYPEMASK[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_POSITION_TARGET_TYPEMASK = map[string]POSITION_TARGET_TYPEMASK{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *POSITION_TARGET_TYPEMASK) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_POSITION_TARGET_TYPEMASK[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e POSITION_TARGET_TYPEMASK) String() string {
	if l, ok := labels_POSITION_TARGET_TYPEMASK[e]; ok {
		return l
	}
	return "invalid value"
}
