//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package avssuas

import (
	"errors"
)

type MAG_CAL_STATUS uint32

const (
	MAG_CAL_NOT_STARTED      MAG_CAL_STATUS = 0
	MAG_CAL_WAITING_TO_START MAG_CAL_STATUS = 1
	MAG_CAL_RUNNING_STEP_ONE MAG_CAL_STATUS = 2
	MAG_CAL_RUNNING_STEP_TWO MAG_CAL_STATUS = 3
	MAG_CAL_SUCCESS          MAG_CAL_STATUS = 4
	MAG_CAL_FAILED           MAG_CAL_STATUS = 5
	MAG_CAL_BAD_ORIENTATION  MAG_CAL_STATUS = 6
	MAG_CAL_BAD_RADIUS       MAG_CAL_STATUS = 7
)

var labels_MAG_CAL_STATUS = map[MAG_CAL_STATUS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAG_CAL_STATUS) MarshalText() ([]byte, error) {
	if l, ok := labels_MAG_CAL_STATUS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAG_CAL_STATUS = map[string]MAG_CAL_STATUS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAG_CAL_STATUS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAG_CAL_STATUS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAG_CAL_STATUS) String() string {
	if l, ok := labels_MAG_CAL_STATUS[e]; ok {
		return l
	}
	return "invalid value"
}
