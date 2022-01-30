//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ualberta

import (
	"errors"
)

// These values define the type of firmware release.  These values indicate the first version or release of this type.  For example the first alpha release would be 64, the second would be 65.
type FIRMWARE_VERSION_TYPE uint32

const (
	// development release
	FIRMWARE_VERSION_TYPE_DEV FIRMWARE_VERSION_TYPE = 0
	// alpha release
	FIRMWARE_VERSION_TYPE_ALPHA FIRMWARE_VERSION_TYPE = 64
	// beta release
	FIRMWARE_VERSION_TYPE_BETA FIRMWARE_VERSION_TYPE = 128
	// release candidate
	FIRMWARE_VERSION_TYPE_RC FIRMWARE_VERSION_TYPE = 192
	// official stable release
	FIRMWARE_VERSION_TYPE_OFFICIAL FIRMWARE_VERSION_TYPE = 255
)

var labels_FIRMWARE_VERSION_TYPE = map[FIRMWARE_VERSION_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e FIRMWARE_VERSION_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_FIRMWARE_VERSION_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_FIRMWARE_VERSION_TYPE = map[string]FIRMWARE_VERSION_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *FIRMWARE_VERSION_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_FIRMWARE_VERSION_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e FIRMWARE_VERSION_TYPE) String() string {
	if l, ok := labels_FIRMWARE_VERSION_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
