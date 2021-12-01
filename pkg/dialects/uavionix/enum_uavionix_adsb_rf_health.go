//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package uavionix

import (
	"errors"
)

// Status flags for ADS-B transponder dynamic output
type UAVIONIX_ADSB_RF_HEALTH int

const (
	UAVIONIX_ADSB_RF_HEALTH_INITIALIZING UAVIONIX_ADSB_RF_HEALTH = 0
	UAVIONIX_ADSB_RF_HEALTH_OK           UAVIONIX_ADSB_RF_HEALTH = 1
	UAVIONIX_ADSB_RF_HEALTH_FAIL_TX      UAVIONIX_ADSB_RF_HEALTH = 2
	UAVIONIX_ADSB_RF_HEALTH_FAIL_RX      UAVIONIX_ADSB_RF_HEALTH = 16
)

var labels_UAVIONIX_ADSB_RF_HEALTH = map[UAVIONIX_ADSB_RF_HEALTH]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UAVIONIX_ADSB_RF_HEALTH) MarshalText() ([]byte, error) {
	if l, ok := labels_UAVIONIX_ADSB_RF_HEALTH[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_UAVIONIX_ADSB_RF_HEALTH = map[string]UAVIONIX_ADSB_RF_HEALTH{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UAVIONIX_ADSB_RF_HEALTH) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_UAVIONIX_ADSB_RF_HEALTH[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e UAVIONIX_ADSB_RF_HEALTH) String() string {
	if l, ok := labels_UAVIONIX_ADSB_RF_HEALTH[e]; ok {
		return l
	}
	return "invalid value"
}
