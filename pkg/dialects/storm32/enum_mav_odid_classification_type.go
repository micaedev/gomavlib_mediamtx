//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package storm32

import (
	"errors"
)

type MAV_ODID_CLASSIFICATION_TYPE uint32

const (
	// The classification type for the UA is undeclared.
	MAV_ODID_CLASSIFICATION_TYPE_UNDECLARED MAV_ODID_CLASSIFICATION_TYPE = 0
	// The classification type for the UA follows EU (European Union) specifications.
	MAV_ODID_CLASSIFICATION_TYPE_EU MAV_ODID_CLASSIFICATION_TYPE = 1
)

var labels_MAV_ODID_CLASSIFICATION_TYPE = map[MAV_ODID_CLASSIFICATION_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_CLASSIFICATION_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ODID_CLASSIFICATION_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ODID_CLASSIFICATION_TYPE = map[string]MAV_ODID_CLASSIFICATION_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_CLASSIFICATION_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ODID_CLASSIFICATION_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_CLASSIFICATION_TYPE) String() string {
	if l, ok := labels_MAV_ODID_CLASSIFICATION_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
