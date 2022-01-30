//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package storm32

import (
	"errors"
)

type GIMBAL_AXIS uint32

const (
	// Gimbal yaw axis.
	GIMBAL_AXIS_YAW GIMBAL_AXIS = 0
	// Gimbal pitch axis.
	GIMBAL_AXIS_PITCH GIMBAL_AXIS = 1
	// Gimbal roll axis.
	GIMBAL_AXIS_ROLL GIMBAL_AXIS = 2
)

var labels_GIMBAL_AXIS = map[GIMBAL_AXIS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GIMBAL_AXIS) MarshalText() ([]byte, error) {
	if l, ok := labels_GIMBAL_AXIS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GIMBAL_AXIS = map[string]GIMBAL_AXIS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GIMBAL_AXIS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GIMBAL_AXIS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GIMBAL_AXIS) String() string {
	if l, ok := labels_GIMBAL_AXIS[e]; ok {
		return l
	}
	return "invalid value"
}
