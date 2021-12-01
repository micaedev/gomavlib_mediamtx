//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package matrixpilot

import (
	"errors"
)

// Enumeration of landed detector states
type MAV_LANDED_STATE int

const (
	// MAV landed state is unknown
	MAV_LANDED_STATE_UNDEFINED MAV_LANDED_STATE = 0
	// MAV is landed (on ground)
	MAV_LANDED_STATE_ON_GROUND MAV_LANDED_STATE = 1
	// MAV is in air
	MAV_LANDED_STATE_IN_AIR MAV_LANDED_STATE = 2
	// MAV currently taking off
	MAV_LANDED_STATE_TAKEOFF MAV_LANDED_STATE = 3
	// MAV currently landing
	MAV_LANDED_STATE_LANDING MAV_LANDED_STATE = 4
)

var labels_MAV_LANDED_STATE = map[MAV_LANDED_STATE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_LANDED_STATE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_LANDED_STATE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_LANDED_STATE = map[string]MAV_LANDED_STATE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_LANDED_STATE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_LANDED_STATE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_LANDED_STATE) String() string {
	if l, ok := labels_MAV_LANDED_STATE[e]; ok {
		return l
	}
	return "invalid value"
}
