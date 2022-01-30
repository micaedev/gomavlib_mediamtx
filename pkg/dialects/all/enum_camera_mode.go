//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package all

import (
	"errors"
)

// Camera Modes.
type CAMERA_MODE uint32

const (
	// Camera is in image/photo capture mode.
	CAMERA_MODE_IMAGE CAMERA_MODE = 0
	// Camera is in video capture mode.
	CAMERA_MODE_VIDEO CAMERA_MODE = 1
	// Camera is in image survey capture mode. It allows for camera controller to do specific settings for surveys.
	CAMERA_MODE_IMAGE_SURVEY CAMERA_MODE = 2
)

var labels_CAMERA_MODE = map[CAMERA_MODE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e CAMERA_MODE) MarshalText() ([]byte, error) {
	if l, ok := labels_CAMERA_MODE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_CAMERA_MODE = map[string]CAMERA_MODE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *CAMERA_MODE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_CAMERA_MODE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e CAMERA_MODE) String() string {
	if l, ok := labels_CAMERA_MODE[e]; ok {
		return l
	}
	return "invalid value"
}
