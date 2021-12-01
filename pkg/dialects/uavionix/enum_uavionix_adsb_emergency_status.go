//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package uavionix

import (
	"errors"
)

// Emergency status encoding
type UAVIONIX_ADSB_EMERGENCY_STATUS int

const (
	UAVIONIX_ADSB_OUT_NO_EMERGENCY                    UAVIONIX_ADSB_EMERGENCY_STATUS = 0
	UAVIONIX_ADSB_OUT_GENERAL_EMERGENCY               UAVIONIX_ADSB_EMERGENCY_STATUS = 1
	UAVIONIX_ADSB_OUT_LIFEGUARD_EMERGENCY             UAVIONIX_ADSB_EMERGENCY_STATUS = 2
	UAVIONIX_ADSB_OUT_MINIMUM_FUEL_EMERGENCY          UAVIONIX_ADSB_EMERGENCY_STATUS = 3
	UAVIONIX_ADSB_OUT_NO_COMM_EMERGENCY               UAVIONIX_ADSB_EMERGENCY_STATUS = 4
	UAVIONIX_ADSB_OUT_UNLAWFUL_INTERFERANCE_EMERGENCY UAVIONIX_ADSB_EMERGENCY_STATUS = 5
	UAVIONIX_ADSB_OUT_DOWNED_AIRCRAFT_EMERGENCY       UAVIONIX_ADSB_EMERGENCY_STATUS = 6
	UAVIONIX_ADSB_OUT_RESERVED                        UAVIONIX_ADSB_EMERGENCY_STATUS = 7
)

var labels_UAVIONIX_ADSB_EMERGENCY_STATUS = map[UAVIONIX_ADSB_EMERGENCY_STATUS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UAVIONIX_ADSB_EMERGENCY_STATUS) MarshalText() ([]byte, error) {
	if l, ok := labels_UAVIONIX_ADSB_EMERGENCY_STATUS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_UAVIONIX_ADSB_EMERGENCY_STATUS = map[string]UAVIONIX_ADSB_EMERGENCY_STATUS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UAVIONIX_ADSB_EMERGENCY_STATUS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_UAVIONIX_ADSB_EMERGENCY_STATUS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e UAVIONIX_ADSB_EMERGENCY_STATUS) String() string {
	if l, ok := labels_UAVIONIX_ADSB_EMERGENCY_STATUS[e]; ok {
		return l
	}
	return "invalid value"
}
