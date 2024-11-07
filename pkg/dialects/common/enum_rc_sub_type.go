//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strconv"
)

// RC sub-type of types defined in RC_TYPE. Used in MAV_CMD_START_RX_PAIR. Ignored if value does not correspond to the set RC_TYPE.
type RC_SUB_TYPE uint64

const (
	// Spektrum DSM2
	RC_SUB_TYPE_SPEKTRUM_DSM2 RC_SUB_TYPE = 0
	// Spektrum DSMX
	RC_SUB_TYPE_SPEKTRUM_DSMX RC_SUB_TYPE = 1
	// Spektrum DSMX8
	RC_SUB_TYPE_SPEKTRUM_DSMX8 RC_SUB_TYPE = 2
)

var labels_RC_SUB_TYPE = map[RC_SUB_TYPE]string{
	RC_SUB_TYPE_SPEKTRUM_DSM2:  "RC_SUB_TYPE_SPEKTRUM_DSM2",
	RC_SUB_TYPE_SPEKTRUM_DSMX:  "RC_SUB_TYPE_SPEKTRUM_DSMX",
	RC_SUB_TYPE_SPEKTRUM_DSMX8: "RC_SUB_TYPE_SPEKTRUM_DSMX8",
}

var values_RC_SUB_TYPE = map[string]RC_SUB_TYPE{
	"RC_SUB_TYPE_SPEKTRUM_DSM2":  RC_SUB_TYPE_SPEKTRUM_DSM2,
	"RC_SUB_TYPE_SPEKTRUM_DSMX":  RC_SUB_TYPE_SPEKTRUM_DSMX,
	"RC_SUB_TYPE_SPEKTRUM_DSMX8": RC_SUB_TYPE_SPEKTRUM_DSMX8,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e RC_SUB_TYPE) MarshalText() ([]byte, error) {
	if name, ok := labels_RC_SUB_TYPE[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *RC_SUB_TYPE) UnmarshalText(text []byte) error {
	if value, ok := values_RC_SUB_TYPE[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = RC_SUB_TYPE(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e RC_SUB_TYPE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
