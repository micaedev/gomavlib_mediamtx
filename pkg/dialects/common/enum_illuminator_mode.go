//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strconv"
)

// Modes of illuminator
type ILLUMINATOR_MODE uint64

const (
	// Illuminator mode is not specified/unknown
	ILLUMINATOR_MODE_UNKNOWN ILLUMINATOR_MODE = 0
	// Illuminator behavior is controlled by MAV_CMD_DO_ILLUMINATOR_CONFIGURE settings
	ILLUMINATOR_MODE_INTERNAL_CONTROL ILLUMINATOR_MODE = 1
	// Illuminator behavior is controlled by external factors: e.g. an external hardware signal
	ILLUMINATOR_MODE_EXTERNAL_SYNC ILLUMINATOR_MODE = 2
)

var labels_ILLUMINATOR_MODE = map[ILLUMINATOR_MODE]string{
	ILLUMINATOR_MODE_UNKNOWN:          "ILLUMINATOR_MODE_UNKNOWN",
	ILLUMINATOR_MODE_INTERNAL_CONTROL: "ILLUMINATOR_MODE_INTERNAL_CONTROL",
	ILLUMINATOR_MODE_EXTERNAL_SYNC:    "ILLUMINATOR_MODE_EXTERNAL_SYNC",
}

var values_ILLUMINATOR_MODE = map[string]ILLUMINATOR_MODE{
	"ILLUMINATOR_MODE_UNKNOWN":          ILLUMINATOR_MODE_UNKNOWN,
	"ILLUMINATOR_MODE_INTERNAL_CONTROL": ILLUMINATOR_MODE_INTERNAL_CONTROL,
	"ILLUMINATOR_MODE_EXTERNAL_SYNC":    ILLUMINATOR_MODE_EXTERNAL_SYNC,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ILLUMINATOR_MODE) MarshalText() ([]byte, error) {
	if name, ok := labels_ILLUMINATOR_MODE[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ILLUMINATOR_MODE) UnmarshalText(text []byte) error {
	if value, ok := values_ILLUMINATOR_MODE[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = ILLUMINATOR_MODE(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e ILLUMINATOR_MODE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
