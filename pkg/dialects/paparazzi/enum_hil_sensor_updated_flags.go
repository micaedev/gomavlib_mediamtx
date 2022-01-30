//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package paparazzi

import (
	"errors"
)

// Flags in the HIL_SENSOR message indicate which fields have updated since the last message
type HIL_SENSOR_UPDATED_FLAGS uint32

const (
	// None of the fields in HIL_SENSOR have been updated
	HIL_SENSOR_UPDATED_NONE HIL_SENSOR_UPDATED_FLAGS = 0
	// The value in the xacc field has been updated
	HIL_SENSOR_UPDATED_XACC HIL_SENSOR_UPDATED_FLAGS = 1
	// The value in the yacc field has been updated
	HIL_SENSOR_UPDATED_YACC HIL_SENSOR_UPDATED_FLAGS = 2
	// The value in the zacc field has been updated
	HIL_SENSOR_UPDATED_ZACC HIL_SENSOR_UPDATED_FLAGS = 4
	// The value in the xgyro field has been updated
	HIL_SENSOR_UPDATED_XGYRO HIL_SENSOR_UPDATED_FLAGS = 8
	// The value in the ygyro field has been updated
	HIL_SENSOR_UPDATED_YGYRO HIL_SENSOR_UPDATED_FLAGS = 16
	// The value in the zgyro field has been updated
	HIL_SENSOR_UPDATED_ZGYRO HIL_SENSOR_UPDATED_FLAGS = 32
	// The value in the xmag field has been updated
	HIL_SENSOR_UPDATED_XMAG HIL_SENSOR_UPDATED_FLAGS = 64
	// The value in the ymag field has been updated
	HIL_SENSOR_UPDATED_YMAG HIL_SENSOR_UPDATED_FLAGS = 128
	// The value in the zmag field has been updated
	HIL_SENSOR_UPDATED_ZMAG HIL_SENSOR_UPDATED_FLAGS = 256
	// The value in the abs_pressure field has been updated
	HIL_SENSOR_UPDATED_ABS_PRESSURE HIL_SENSOR_UPDATED_FLAGS = 512
	// The value in the diff_pressure field has been updated
	HIL_SENSOR_UPDATED_DIFF_PRESSURE HIL_SENSOR_UPDATED_FLAGS = 1024
	// The value in the pressure_alt field has been updated
	HIL_SENSOR_UPDATED_PRESSURE_ALT HIL_SENSOR_UPDATED_FLAGS = 2048
	// The value in the temperature field has been updated
	HIL_SENSOR_UPDATED_TEMPERATURE HIL_SENSOR_UPDATED_FLAGS = 4096
	// Full reset of attitude/position/velocities/etc was performed in sim (Bit 31).
	HIL_SENSOR_UPDATED_RESET HIL_SENSOR_UPDATED_FLAGS = 2147483648
)

var labels_HIL_SENSOR_UPDATED_FLAGS = map[HIL_SENSOR_UPDATED_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e HIL_SENSOR_UPDATED_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_HIL_SENSOR_UPDATED_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_HIL_SENSOR_UPDATED_FLAGS = map[string]HIL_SENSOR_UPDATED_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *HIL_SENSOR_UPDATED_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_HIL_SENSOR_UPDATED_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e HIL_SENSOR_UPDATED_FLAGS) String() string {
	if l, ok := labels_HIL_SENSOR_UPDATED_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
