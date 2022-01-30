//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package development

import (
	"errors"
)

type MAV_ODID_OPERATOR_ID_TYPE uint32

const (
	// CAA (Civil Aviation Authority) registered operator ID.
	MAV_ODID_OPERATOR_ID_TYPE_CAA MAV_ODID_OPERATOR_ID_TYPE = 0
)

var labels_MAV_ODID_OPERATOR_ID_TYPE = map[MAV_ODID_OPERATOR_ID_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_OPERATOR_ID_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ODID_OPERATOR_ID_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ODID_OPERATOR_ID_TYPE = map[string]MAV_ODID_OPERATOR_ID_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_OPERATOR_ID_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ODID_OPERATOR_ID_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_OPERATOR_ID_TYPE) String() string {
	if l, ok := labels_MAV_ODID_OPERATOR_ID_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
