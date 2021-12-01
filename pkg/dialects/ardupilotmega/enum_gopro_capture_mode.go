//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

type GOPRO_CAPTURE_MODE int

const (
	// Video mode.
	GOPRO_CAPTURE_MODE_VIDEO GOPRO_CAPTURE_MODE = 0
	// Photo mode.
	GOPRO_CAPTURE_MODE_PHOTO GOPRO_CAPTURE_MODE = 1
	// Burst mode, Hero 3+ only.
	GOPRO_CAPTURE_MODE_BURST GOPRO_CAPTURE_MODE = 2
	// Time lapse mode, Hero 3+ only.
	GOPRO_CAPTURE_MODE_TIME_LAPSE GOPRO_CAPTURE_MODE = 3
	// Multi shot mode, Hero 4 only.
	GOPRO_CAPTURE_MODE_MULTI_SHOT GOPRO_CAPTURE_MODE = 4
	// Playback mode, Hero 4 only, silver only except when LCD or HDMI is connected to black.
	GOPRO_CAPTURE_MODE_PLAYBACK GOPRO_CAPTURE_MODE = 5
	// Playback mode, Hero 4 only.
	GOPRO_CAPTURE_MODE_SETUP GOPRO_CAPTURE_MODE = 6
	// Mode not yet known.
	GOPRO_CAPTURE_MODE_UNKNOWN GOPRO_CAPTURE_MODE = 255
)

var labels_GOPRO_CAPTURE_MODE = map[GOPRO_CAPTURE_MODE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GOPRO_CAPTURE_MODE) MarshalText() ([]byte, error) {
	if l, ok := labels_GOPRO_CAPTURE_MODE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GOPRO_CAPTURE_MODE = map[string]GOPRO_CAPTURE_MODE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GOPRO_CAPTURE_MODE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GOPRO_CAPTURE_MODE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GOPRO_CAPTURE_MODE) String() string {
	if l, ok := labels_GOPRO_CAPTURE_MODE[e]; ok {
		return l
	}
	return "invalid value"
}
