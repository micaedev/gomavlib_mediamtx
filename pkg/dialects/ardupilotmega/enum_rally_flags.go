//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

// Flags in RALLY_POINT message.
type RALLY_FLAGS int

const (
	// Flag set when requiring favorable winds for landing.
	FAVORABLE_WIND RALLY_FLAGS = 1
	// Flag set when plane is to immediately descend to break altitude and land without GCS intervention. Flag not set when plane is to loiter at Rally point until commanded to land.
	LAND_IMMEDIATELY RALLY_FLAGS = 2
)

var labels_RALLY_FLAGS = map[RALLY_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e RALLY_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_RALLY_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_RALLY_FLAGS = map[string]RALLY_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *RALLY_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_RALLY_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e RALLY_FLAGS) String() string {
	if l, ok := labels_RALLY_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
