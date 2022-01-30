//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package matrixpilot

import (
	"errors"
)

// Zoom types for MAV_CMD_SET_CAMERA_ZOOM
type CAMERA_ZOOM_TYPE uint32

const (
	// Zoom one step increment (-1 for wide, 1 for tele)
	ZOOM_TYPE_STEP CAMERA_ZOOM_TYPE = 0
	// Continuous zoom up/down until stopped (-1 for wide, 1 for tele, 0 to stop zooming)
	ZOOM_TYPE_CONTINUOUS CAMERA_ZOOM_TYPE = 1
	// Zoom value as proportion of full camera range (a value between 0.0 and 100.0)
	ZOOM_TYPE_RANGE CAMERA_ZOOM_TYPE = 2
	// Zoom value/variable focal length in milimetres. Note that there is no message to get the valid zoom range of the camera, so this can type can only be used for cameras where the zoom range is known (implying that this cannot reliably be used in a GCS for an arbitrary camera)
	ZOOM_TYPE_FOCAL_LENGTH CAMERA_ZOOM_TYPE = 3
)

var labels_CAMERA_ZOOM_TYPE = map[CAMERA_ZOOM_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e CAMERA_ZOOM_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_CAMERA_ZOOM_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_CAMERA_ZOOM_TYPE = map[string]CAMERA_ZOOM_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *CAMERA_ZOOM_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_CAMERA_ZOOM_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e CAMERA_ZOOM_TYPE) String() string {
	if l, ok := labels_CAMERA_ZOOM_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
