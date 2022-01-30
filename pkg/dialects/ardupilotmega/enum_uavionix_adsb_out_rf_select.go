//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

// Transceiver RF control flags for ADS-B transponder dynamic reports
type UAVIONIX_ADSB_OUT_RF_SELECT uint32

const (
	UAVIONIX_ADSB_OUT_RF_SELECT_STANDBY    UAVIONIX_ADSB_OUT_RF_SELECT = 0
	UAVIONIX_ADSB_OUT_RF_SELECT_RX_ENABLED UAVIONIX_ADSB_OUT_RF_SELECT = 1
	UAVIONIX_ADSB_OUT_RF_SELECT_TX_ENABLED UAVIONIX_ADSB_OUT_RF_SELECT = 2
)

var labels_UAVIONIX_ADSB_OUT_RF_SELECT = map[UAVIONIX_ADSB_OUT_RF_SELECT]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UAVIONIX_ADSB_OUT_RF_SELECT) MarshalText() ([]byte, error) {
	if l, ok := labels_UAVIONIX_ADSB_OUT_RF_SELECT[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_UAVIONIX_ADSB_OUT_RF_SELECT = map[string]UAVIONIX_ADSB_OUT_RF_SELECT{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UAVIONIX_ADSB_OUT_RF_SELECT) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_UAVIONIX_ADSB_OUT_RF_SELECT[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e UAVIONIX_ADSB_OUT_RF_SELECT) String() string {
	if l, ok := labels_UAVIONIX_ADSB_OUT_RF_SELECT[e]; ok {
		return l
	}
	return "invalid value"
}
