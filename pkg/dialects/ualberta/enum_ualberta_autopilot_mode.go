//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package ualberta

import (
	"errors"
)

// Available autopilot modes for ualberta uav
type UALBERTA_AUTOPILOT_MODE int

const (
	// Raw input pulse widts sent to output
	MODE_MANUAL_DIRECT UALBERTA_AUTOPILOT_MODE = 1
	// Inputs are normalized using calibration, the converted back to raw pulse widths for output
	MODE_MANUAL_SCALED UALBERTA_AUTOPILOT_MODE = 2
	// dfsdfs
	MODE_AUTO_PID_ATT UALBERTA_AUTOPILOT_MODE = 3
	// dfsfds
	MODE_AUTO_PID_VEL UALBERTA_AUTOPILOT_MODE = 4
	// dfsdfsdfs
	MODE_AUTO_PID_POS UALBERTA_AUTOPILOT_MODE = 5
)

var labels_UALBERTA_AUTOPILOT_MODE = map[UALBERTA_AUTOPILOT_MODE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UALBERTA_AUTOPILOT_MODE) MarshalText() ([]byte, error) {
	if l, ok := labels_UALBERTA_AUTOPILOT_MODE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_UALBERTA_AUTOPILOT_MODE = map[string]UALBERTA_AUTOPILOT_MODE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UALBERTA_AUTOPILOT_MODE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_UALBERTA_AUTOPILOT_MODE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e UALBERTA_AUTOPILOT_MODE) String() string {
	if l, ok := labels_UALBERTA_AUTOPILOT_MODE[e]; ok {
		return l
	}
	return "invalid value"
}
