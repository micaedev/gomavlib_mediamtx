//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

type MAV_ODID_ID_TYPE int

const (
	// No type defined.
	MAV_ODID_ID_TYPE_NONE MAV_ODID_ID_TYPE = 0
	// Manufacturer Serial Number (ANSI/CTA-2063 format).
	MAV_ODID_ID_TYPE_SERIAL_NUMBER MAV_ODID_ID_TYPE = 1
	// CAA (Civil Aviation Authority) registered ID. Format: [ICAO Country Code].[CAA Assigned ID].
	MAV_ODID_ID_TYPE_CAA_REGISTRATION_ID MAV_ODID_ID_TYPE = 2
	// UTM (Unmanned Traffic Management) assigned UUID (RFC4122).
	MAV_ODID_ID_TYPE_UTM_ASSIGNED_UUID MAV_ODID_ID_TYPE = 3
	// A 20 byte ID for a specific flight/session. The exact ID type is indicated by the first byte of uas_id and these type values are managed by ICAO.
	MAV_ODID_ID_TYPE_SPECIFIC_SESSION_ID MAV_ODID_ID_TYPE = 4
)

var labels_MAV_ODID_ID_TYPE = map[MAV_ODID_ID_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_ID_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ODID_ID_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ODID_ID_TYPE = map[string]MAV_ODID_ID_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_ID_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ODID_ID_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_ID_TYPE) String() string {
	if l, ok := labels_MAV_ODID_ID_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
