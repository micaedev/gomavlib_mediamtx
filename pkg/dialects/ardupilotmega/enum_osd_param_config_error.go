//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"fmt"
	"strconv"
)

// The error type for the OSD parameter editor.
type OSD_PARAM_CONFIG_ERROR uint64

const (
	OSD_PARAM_SUCCESS                 OSD_PARAM_CONFIG_ERROR = 0
	OSD_PARAM_INVALID_SCREEN          OSD_PARAM_CONFIG_ERROR = 1
	OSD_PARAM_INVALID_PARAMETER_INDEX OSD_PARAM_CONFIG_ERROR = 2
	OSD_PARAM_INVALID_PARAMETER       OSD_PARAM_CONFIG_ERROR = 3
)

var value_to_label_OSD_PARAM_CONFIG_ERROR = map[OSD_PARAM_CONFIG_ERROR]string{
	OSD_PARAM_SUCCESS:                 "OSD_PARAM_SUCCESS",
	OSD_PARAM_INVALID_SCREEN:          "OSD_PARAM_INVALID_SCREEN",
	OSD_PARAM_INVALID_PARAMETER_INDEX: "OSD_PARAM_INVALID_PARAMETER_INDEX",
	OSD_PARAM_INVALID_PARAMETER:       "OSD_PARAM_INVALID_PARAMETER",
}

var label_to_value_OSD_PARAM_CONFIG_ERROR = map[string]OSD_PARAM_CONFIG_ERROR{
	"OSD_PARAM_SUCCESS":                 OSD_PARAM_SUCCESS,
	"OSD_PARAM_INVALID_SCREEN":          OSD_PARAM_INVALID_SCREEN,
	"OSD_PARAM_INVALID_PARAMETER_INDEX": OSD_PARAM_INVALID_PARAMETER_INDEX,
	"OSD_PARAM_INVALID_PARAMETER":       OSD_PARAM_INVALID_PARAMETER,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e OSD_PARAM_CONFIG_ERROR) MarshalText() ([]byte, error) {
	if name, ok := value_to_label_OSD_PARAM_CONFIG_ERROR[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *OSD_PARAM_CONFIG_ERROR) UnmarshalText(text []byte) error {
	if value, ok := label_to_value_OSD_PARAM_CONFIG_ERROR[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = OSD_PARAM_CONFIG_ERROR(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e OSD_PARAM_CONFIG_ERROR) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
