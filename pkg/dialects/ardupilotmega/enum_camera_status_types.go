//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

type CAMERA_STATUS_TYPES uint32

const (
	// Camera heartbeat, announce camera component ID at 1Hz.
	CAMERA_STATUS_TYPE_HEARTBEAT CAMERA_STATUS_TYPES = 0
	// Camera image triggered.
	CAMERA_STATUS_TYPE_TRIGGER CAMERA_STATUS_TYPES = 1
	// Camera connection lost.
	CAMERA_STATUS_TYPE_DISCONNECT CAMERA_STATUS_TYPES = 2
	// Camera unknown error.
	CAMERA_STATUS_TYPE_ERROR CAMERA_STATUS_TYPES = 3
	// Camera battery low. Parameter p1 shows reported voltage.
	CAMERA_STATUS_TYPE_LOWBATT CAMERA_STATUS_TYPES = 4
	// Camera storage low. Parameter p1 shows reported shots remaining.
	CAMERA_STATUS_TYPE_LOWSTORE CAMERA_STATUS_TYPES = 5
	// Camera storage low. Parameter p1 shows reported video minutes remaining.
	CAMERA_STATUS_TYPE_LOWSTOREV CAMERA_STATUS_TYPES = 6
)

var labels_CAMERA_STATUS_TYPES = map[CAMERA_STATUS_TYPES]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e CAMERA_STATUS_TYPES) MarshalText() ([]byte, error) {
	if l, ok := labels_CAMERA_STATUS_TYPES[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_CAMERA_STATUS_TYPES = map[string]CAMERA_STATUS_TYPES{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *CAMERA_STATUS_TYPES) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_CAMERA_STATUS_TYPES[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e CAMERA_STATUS_TYPES) String() string {
	if l, ok := labels_CAMERA_STATUS_TYPES[e]; ok {
		return l
	}
	return "invalid value"
}
