//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

// Enumeration of battery types
type MAV_BATTERY_TYPE uint32

const (
	// Not specified.
	MAV_BATTERY_TYPE_UNKNOWN MAV_BATTERY_TYPE = 0
	// Lithium polymer battery
	MAV_BATTERY_TYPE_LIPO MAV_BATTERY_TYPE = 1
	// Lithium-iron-phosphate battery
	MAV_BATTERY_TYPE_LIFE MAV_BATTERY_TYPE = 2
	// Lithium-ION battery
	MAV_BATTERY_TYPE_LION MAV_BATTERY_TYPE = 3
	// Nickel metal hydride battery
	MAV_BATTERY_TYPE_NIMH MAV_BATTERY_TYPE = 4
)

var labels_MAV_BATTERY_TYPE = map[MAV_BATTERY_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_BATTERY_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_BATTERY_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_BATTERY_TYPE = map[string]MAV_BATTERY_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_BATTERY_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_BATTERY_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_BATTERY_TYPE) String() string {
	if l, ok := labels_MAV_BATTERY_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
