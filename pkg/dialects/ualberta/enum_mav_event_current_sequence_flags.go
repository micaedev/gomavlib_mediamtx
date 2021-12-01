//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package ualberta

import (
	"errors"
)

// Flags for CURRENT_EVENT_SEQUENCE.
type MAV_EVENT_CURRENT_SEQUENCE_FLAGS int

const (
	// A sequence reset has happened (e.g. vehicle reboot).
	MAV_EVENT_CURRENT_SEQUENCE_FLAGS_RESET MAV_EVENT_CURRENT_SEQUENCE_FLAGS = 1
)

var labels_MAV_EVENT_CURRENT_SEQUENCE_FLAGS = map[MAV_EVENT_CURRENT_SEQUENCE_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_EVENT_CURRENT_SEQUENCE_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_EVENT_CURRENT_SEQUENCE_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_EVENT_CURRENT_SEQUENCE_FLAGS = map[string]MAV_EVENT_CURRENT_SEQUENCE_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_EVENT_CURRENT_SEQUENCE_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_EVENT_CURRENT_SEQUENCE_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_EVENT_CURRENT_SEQUENCE_FLAGS) String() string {
	if l, ok := labels_MAV_EVENT_CURRENT_SEQUENCE_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
