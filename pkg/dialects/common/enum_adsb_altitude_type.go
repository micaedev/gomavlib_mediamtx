//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strconv"
)

// Enumeration of the ADSB altimeter types
type ADSB_ALTITUDE_TYPE uint32

const (
	// Altitude reported from a Baro source using QNH reference
	ADSB_ALTITUDE_TYPE_PRESSURE_QNH ADSB_ALTITUDE_TYPE = 0
	// Altitude reported from a GNSS source
	ADSB_ALTITUDE_TYPE_GEOMETRIC ADSB_ALTITUDE_TYPE = 1
)

var labels_ADSB_ALTITUDE_TYPE = map[ADSB_ALTITUDE_TYPE]string{
	ADSB_ALTITUDE_TYPE_PRESSURE_QNH: "ADSB_ALTITUDE_TYPE_PRESSURE_QNH",
	ADSB_ALTITUDE_TYPE_GEOMETRIC:    "ADSB_ALTITUDE_TYPE_GEOMETRIC",
}

var values_ADSB_ALTITUDE_TYPE = map[string]ADSB_ALTITUDE_TYPE{
	"ADSB_ALTITUDE_TYPE_PRESSURE_QNH": ADSB_ALTITUDE_TYPE_PRESSURE_QNH,
	"ADSB_ALTITUDE_TYPE_GEOMETRIC":    ADSB_ALTITUDE_TYPE_GEOMETRIC,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ADSB_ALTITUDE_TYPE) MarshalText() ([]byte, error) {
	if name, ok := labels_ADSB_ALTITUDE_TYPE[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ADSB_ALTITUDE_TYPE) UnmarshalText(text []byte) error {
	if value, ok := values_ADSB_ALTITUDE_TYPE[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = ADSB_ALTITUDE_TYPE(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e ADSB_ALTITUDE_TYPE) String() string {
	if name, ok := labels_ADSB_ALTITUDE_TYPE[e]; ok {
		return name
	}
	return strconv.Itoa(int(e))
}
