//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package common

import (
	"errors"
)

// Flags in the HIGHRES_IMU message indicate which fields have updated since the last message
type HIGHRES_IMU_UPDATED_FLAGS uint32

const (
	// None of the fields in HIGHRES_IMU have been updated
	HIGHRES_IMU_UPDATED_NONE HIGHRES_IMU_UPDATED_FLAGS = 0
	// The value in the xacc field has been updated
	HIGHRES_IMU_UPDATED_XACC HIGHRES_IMU_UPDATED_FLAGS = 1
	// The value in the yacc field has been updated
	HIGHRES_IMU_UPDATED_YACC HIGHRES_IMU_UPDATED_FLAGS = 2
	// The value in the zacc field has been updated since
	HIGHRES_IMU_UPDATED_ZACC HIGHRES_IMU_UPDATED_FLAGS = 4
	// The value in the xgyro field has been updated
	HIGHRES_IMU_UPDATED_XGYRO HIGHRES_IMU_UPDATED_FLAGS = 8
	// The value in the ygyro field has been updated
	HIGHRES_IMU_UPDATED_YGYRO HIGHRES_IMU_UPDATED_FLAGS = 16
	// The value in the zgyro field has been updated
	HIGHRES_IMU_UPDATED_ZGYRO HIGHRES_IMU_UPDATED_FLAGS = 32
	// The value in the xmag field has been updated
	HIGHRES_IMU_UPDATED_XMAG HIGHRES_IMU_UPDATED_FLAGS = 64
	// The value in the ymag field has been updated
	HIGHRES_IMU_UPDATED_YMAG HIGHRES_IMU_UPDATED_FLAGS = 128
	// The value in the zmag field has been updated
	HIGHRES_IMU_UPDATED_ZMAG HIGHRES_IMU_UPDATED_FLAGS = 256
	// The value in the abs_pressure field has been updated
	HIGHRES_IMU_UPDATED_ABS_PRESSURE HIGHRES_IMU_UPDATED_FLAGS = 512
	// The value in the diff_pressure field has been updated
	HIGHRES_IMU_UPDATED_DIFF_PRESSURE HIGHRES_IMU_UPDATED_FLAGS = 1024
	// The value in the pressure_alt field has been updated
	HIGHRES_IMU_UPDATED_PRESSURE_ALT HIGHRES_IMU_UPDATED_FLAGS = 2048
	// The value in the temperature field has been updated
	HIGHRES_IMU_UPDATED_TEMPERATURE HIGHRES_IMU_UPDATED_FLAGS = 4096
	// All fields in HIGHRES_IMU have been updated.
	HIGHRES_IMU_UPDATED_ALL HIGHRES_IMU_UPDATED_FLAGS = 65535
)

var labels_HIGHRES_IMU_UPDATED_FLAGS = map[HIGHRES_IMU_UPDATED_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e HIGHRES_IMU_UPDATED_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_HIGHRES_IMU_UPDATED_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_HIGHRES_IMU_UPDATED_FLAGS = map[string]HIGHRES_IMU_UPDATED_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *HIGHRES_IMU_UPDATED_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_HIGHRES_IMU_UPDATED_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e HIGHRES_IMU_UPDATED_FLAGS) String() string {
	if l, ok := labels_HIGHRES_IMU_UPDATED_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
