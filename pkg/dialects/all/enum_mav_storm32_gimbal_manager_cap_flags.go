//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package all

import (
	"errors"
)

// Gimbal manager capability flags.
type MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS uint32

const (
	// The gimbal manager supports several profiles.
	MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS_HAS_PROFILES MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS = 1
	// The gimbal manager supports changing the gimbal manager during run time, i.e. can be enabled/disabled.
	MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_CHANGE MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS = 2
)

var labels_MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS = map[MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS = map[string]MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS) String() string {
	if l, ok := labels_MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
