//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package pythonarraytest

import (
	"errors"
)

// Bitmap of options for the MAV_CMD_DO_REPOSITION
type MAV_DO_REPOSITION_FLAGS uint32

const (
	// The aircraft should immediately transition into guided. This should not be set for follow me applications
	MAV_DO_REPOSITION_FLAGS_CHANGE_MODE MAV_DO_REPOSITION_FLAGS = 1
)

var labels_MAV_DO_REPOSITION_FLAGS = map[MAV_DO_REPOSITION_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_DO_REPOSITION_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_DO_REPOSITION_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_DO_REPOSITION_FLAGS = map[string]MAV_DO_REPOSITION_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_DO_REPOSITION_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_DO_REPOSITION_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_DO_REPOSITION_FLAGS) String() string {
	if l, ok := labels_MAV_DO_REPOSITION_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
