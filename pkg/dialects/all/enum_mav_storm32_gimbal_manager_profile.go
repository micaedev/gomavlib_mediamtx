//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package all

import (
	"errors"
)

// Gimbal manager profiles. Only standard profiles are defined. Any implementation can define it's own profile in addition, and should use enum values > 16.
type MAV_STORM32_GIMBAL_MANAGER_PROFILE uint32

const (
	// Default profile. Implementation specific.
	MAV_STORM32_GIMBAL_MANAGER_PROFILE_DEFAULT MAV_STORM32_GIMBAL_MANAGER_PROFILE = 0
	// Custom profile. Configurable profile according to the STorM32 definition. Is configured with STORM32_GIMBAL_MANAGER_PROFIL.
	MAV_STORM32_GIMBAL_MANAGER_PROFILE_CUSTOM MAV_STORM32_GIMBAL_MANAGER_PROFILE = 1
	// Default cooperative profile. Uses STorM32 custom profile with default settings to achieve cooperative behavior.
	MAV_STORM32_GIMBAL_MANAGER_PROFILE_COOPERATIVE MAV_STORM32_GIMBAL_MANAGER_PROFILE = 2
	// Default exclusive profile. Uses STorM32 custom profile with default settings to achieve exclusive behavior.
	MAV_STORM32_GIMBAL_MANAGER_PROFILE_EXCLUSIVE MAV_STORM32_GIMBAL_MANAGER_PROFILE = 3
	// Default priority profile with cooperative behavior for equal priority. Uses STorM32 custom profile with default settings to achieve priority-based behavior.
	MAV_STORM32_GIMBAL_MANAGER_PROFILE_PRIORITY_COOPERATIVE MAV_STORM32_GIMBAL_MANAGER_PROFILE = 4
	// Default priority profile with exclusive behavior for equal priority. Uses STorM32 custom profile with default settings to achieve priority-based behavior.
	MAV_STORM32_GIMBAL_MANAGER_PROFILE_PRIORITY_EXCLUSIVE MAV_STORM32_GIMBAL_MANAGER_PROFILE = 5
)

var labels_MAV_STORM32_GIMBAL_MANAGER_PROFILE = map[MAV_STORM32_GIMBAL_MANAGER_PROFILE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_STORM32_GIMBAL_MANAGER_PROFILE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_STORM32_GIMBAL_MANAGER_PROFILE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_STORM32_GIMBAL_MANAGER_PROFILE = map[string]MAV_STORM32_GIMBAL_MANAGER_PROFILE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_STORM32_GIMBAL_MANAGER_PROFILE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_STORM32_GIMBAL_MANAGER_PROFILE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_STORM32_GIMBAL_MANAGER_PROFILE) String() string {
	if l, ok := labels_MAV_STORM32_GIMBAL_MANAGER_PROFILE[e]; ok {
		return l
	}
	return "invalid value"
}
