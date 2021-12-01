//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package matrixpilot

import (
	"errors"
)

type MAV_ODID_DESC_TYPE int

const (
	// Free-form text description of the purpose of the flight.
	MAV_ODID_DESC_TYPE_TEXT MAV_ODID_DESC_TYPE = 0
)

var labels_MAV_ODID_DESC_TYPE = map[MAV_ODID_DESC_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_DESC_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ODID_DESC_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ODID_DESC_TYPE = map[string]MAV_ODID_DESC_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_DESC_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ODID_DESC_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_DESC_TYPE) String() string {
	if l, ok := labels_MAV_ODID_DESC_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
