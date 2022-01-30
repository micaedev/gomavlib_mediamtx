//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

// Camera tracking modes
type CAMERA_TRACKING_MODE uint32

const (
	// Not tracking
	CAMERA_TRACKING_MODE_NONE CAMERA_TRACKING_MODE = 0
	// Target is a point
	CAMERA_TRACKING_MODE_POINT CAMERA_TRACKING_MODE = 1
	// Target is a rectangle
	CAMERA_TRACKING_MODE_RECTANGLE CAMERA_TRACKING_MODE = 2
)

var labels_CAMERA_TRACKING_MODE = map[CAMERA_TRACKING_MODE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e CAMERA_TRACKING_MODE) MarshalText() ([]byte, error) {
	if l, ok := labels_CAMERA_TRACKING_MODE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_CAMERA_TRACKING_MODE = map[string]CAMERA_TRACKING_MODE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *CAMERA_TRACKING_MODE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_CAMERA_TRACKING_MODE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e CAMERA_TRACKING_MODE) String() string {
	if l, ok := labels_CAMERA_TRACKING_MODE[e]; ok {
		return l
	}
	return "invalid value"
}
