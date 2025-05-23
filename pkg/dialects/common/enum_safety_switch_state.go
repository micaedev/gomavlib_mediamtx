//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strconv"
)

// Possible safety switch states.
type SAFETY_SWITCH_STATE uint64

const (
	// Safety switch is engaged and vehicle should be safe to approach.
	SAFETY_SWITCH_STATE_SAFE SAFETY_SWITCH_STATE = 0
	// Safety switch is NOT engaged and motors, propellers and other actuators should be considered active.
	SAFETY_SWITCH_STATE_DANGEROUS SAFETY_SWITCH_STATE = 1
)

var value_to_label_SAFETY_SWITCH_STATE = map[SAFETY_SWITCH_STATE]string{
	SAFETY_SWITCH_STATE_SAFE:      "SAFETY_SWITCH_STATE_SAFE",
	SAFETY_SWITCH_STATE_DANGEROUS: "SAFETY_SWITCH_STATE_DANGEROUS",
}

var label_to_value_SAFETY_SWITCH_STATE = map[string]SAFETY_SWITCH_STATE{
	"SAFETY_SWITCH_STATE_SAFE":      SAFETY_SWITCH_STATE_SAFE,
	"SAFETY_SWITCH_STATE_DANGEROUS": SAFETY_SWITCH_STATE_DANGEROUS,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e SAFETY_SWITCH_STATE) MarshalText() ([]byte, error) {
	if name, ok := value_to_label_SAFETY_SWITCH_STATE[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *SAFETY_SWITCH_STATE) UnmarshalText(text []byte) error {
	if value, ok := label_to_value_SAFETY_SWITCH_STATE[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = SAFETY_SWITCH_STATE(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e SAFETY_SWITCH_STATE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
