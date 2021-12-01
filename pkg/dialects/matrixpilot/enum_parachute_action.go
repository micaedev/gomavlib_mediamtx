//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package matrixpilot

import (
	"errors"
)

// Parachute actions. Trigger release and enable/disable auto-release.
type PARACHUTE_ACTION int

const (
	// Disable auto-release of parachute (i.e. release triggered by crash detectors).
	PARACHUTE_DISABLE PARACHUTE_ACTION = 0
	// Enable auto-release of parachute.
	PARACHUTE_ENABLE PARACHUTE_ACTION = 1
	// Release parachute and kill motors.
	PARACHUTE_RELEASE PARACHUTE_ACTION = 2
)

var labels_PARACHUTE_ACTION = map[PARACHUTE_ACTION]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e PARACHUTE_ACTION) MarshalText() ([]byte, error) {
	if l, ok := labels_PARACHUTE_ACTION[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_PARACHUTE_ACTION = map[string]PARACHUTE_ACTION{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *PARACHUTE_ACTION) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_PARACHUTE_ACTION[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e PARACHUTE_ACTION) String() string {
	if l, ok := labels_PARACHUTE_ACTION[e]; ok {
		return l
	}
	return "invalid value"
}
