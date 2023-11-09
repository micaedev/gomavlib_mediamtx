//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package icarous

import (
	"fmt"
	"strconv"
)

type ICAROUS_FMS_STATE uint32

const (
	ICAROUS_FMS_STATE_IDLE     ICAROUS_FMS_STATE = 0
	ICAROUS_FMS_STATE_TAKEOFF  ICAROUS_FMS_STATE = 1
	ICAROUS_FMS_STATE_CLIMB    ICAROUS_FMS_STATE = 2
	ICAROUS_FMS_STATE_CRUISE   ICAROUS_FMS_STATE = 3
	ICAROUS_FMS_STATE_APPROACH ICAROUS_FMS_STATE = 4
	ICAROUS_FMS_STATE_LAND     ICAROUS_FMS_STATE = 5
)

var labels_ICAROUS_FMS_STATE = map[ICAROUS_FMS_STATE]string{
	ICAROUS_FMS_STATE_IDLE:     "ICAROUS_FMS_STATE_IDLE",
	ICAROUS_FMS_STATE_TAKEOFF:  "ICAROUS_FMS_STATE_TAKEOFF",
	ICAROUS_FMS_STATE_CLIMB:    "ICAROUS_FMS_STATE_CLIMB",
	ICAROUS_FMS_STATE_CRUISE:   "ICAROUS_FMS_STATE_CRUISE",
	ICAROUS_FMS_STATE_APPROACH: "ICAROUS_FMS_STATE_APPROACH",
	ICAROUS_FMS_STATE_LAND:     "ICAROUS_FMS_STATE_LAND",
}

var values_ICAROUS_FMS_STATE = map[string]ICAROUS_FMS_STATE{
	"ICAROUS_FMS_STATE_IDLE":     ICAROUS_FMS_STATE_IDLE,
	"ICAROUS_FMS_STATE_TAKEOFF":  ICAROUS_FMS_STATE_TAKEOFF,
	"ICAROUS_FMS_STATE_CLIMB":    ICAROUS_FMS_STATE_CLIMB,
	"ICAROUS_FMS_STATE_CRUISE":   ICAROUS_FMS_STATE_CRUISE,
	"ICAROUS_FMS_STATE_APPROACH": ICAROUS_FMS_STATE_APPROACH,
	"ICAROUS_FMS_STATE_LAND":     ICAROUS_FMS_STATE_LAND,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ICAROUS_FMS_STATE) MarshalText() ([]byte, error) {
	if name, ok := labels_ICAROUS_FMS_STATE[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ICAROUS_FMS_STATE) UnmarshalText(text []byte) error {
	if value, ok := values_ICAROUS_FMS_STATE[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = ICAROUS_FMS_STATE(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e ICAROUS_FMS_STATE) String() string {
	if name, ok := labels_ICAROUS_FMS_STATE[e]; ok {
		return name
	}
	return strconv.Itoa(int(e))
}
