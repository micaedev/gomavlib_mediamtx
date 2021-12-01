//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

// A mapping of sub flight modes for custom_mode field of heartbeat.
type SUB_MODE int

const (
	SUB_MODE_STABILIZE SUB_MODE = 0
	SUB_MODE_ACRO      SUB_MODE = 1
	SUB_MODE_ALT_HOLD  SUB_MODE = 2
	SUB_MODE_AUTO      SUB_MODE = 3
	SUB_MODE_GUIDED    SUB_MODE = 4
	SUB_MODE_CIRCLE    SUB_MODE = 7
	SUB_MODE_SURFACE   SUB_MODE = 9
	SUB_MODE_POSHOLD   SUB_MODE = 16
	SUB_MODE_MANUAL    SUB_MODE = 19
)

var labels_SUB_MODE = map[SUB_MODE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e SUB_MODE) MarshalText() ([]byte, error) {
	if l, ok := labels_SUB_MODE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_SUB_MODE = map[string]SUB_MODE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *SUB_MODE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_SUB_MODE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e SUB_MODE) String() string {
	if l, ok := labels_SUB_MODE[e]; ok {
		return l
	}
	return "invalid value"
}
