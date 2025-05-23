//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"fmt"
	"strconv"
	"strings"
)

// Transceiver RF control flags for ADS-B transponder dynamic reports
type UAVIONIX_ADSB_OUT_RF_SELECT uint64

const (
	UAVIONIX_ADSB_OUT_RF_SELECT_RX_ENABLED UAVIONIX_ADSB_OUT_RF_SELECT = 1
	UAVIONIX_ADSB_OUT_RF_SELECT_TX_ENABLED UAVIONIX_ADSB_OUT_RF_SELECT = 2
)

var values_UAVIONIX_ADSB_OUT_RF_SELECT = []UAVIONIX_ADSB_OUT_RF_SELECT{
	UAVIONIX_ADSB_OUT_RF_SELECT_RX_ENABLED,
	UAVIONIX_ADSB_OUT_RF_SELECT_TX_ENABLED,
}

var value_to_label_UAVIONIX_ADSB_OUT_RF_SELECT = map[UAVIONIX_ADSB_OUT_RF_SELECT]string{
	UAVIONIX_ADSB_OUT_RF_SELECT_RX_ENABLED: "UAVIONIX_ADSB_OUT_RF_SELECT_RX_ENABLED",
	UAVIONIX_ADSB_OUT_RF_SELECT_TX_ENABLED: "UAVIONIX_ADSB_OUT_RF_SELECT_TX_ENABLED",
}

var label_to_value_UAVIONIX_ADSB_OUT_RF_SELECT = map[string]UAVIONIX_ADSB_OUT_RF_SELECT{
	"UAVIONIX_ADSB_OUT_RF_SELECT_RX_ENABLED": UAVIONIX_ADSB_OUT_RF_SELECT_RX_ENABLED,
	"UAVIONIX_ADSB_OUT_RF_SELECT_TX_ENABLED": UAVIONIX_ADSB_OUT_RF_SELECT_TX_ENABLED,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UAVIONIX_ADSB_OUT_RF_SELECT) MarshalText() ([]byte, error) {
	if e == 0 {
		return []byte("0"), nil
	}
	var names []string
	for _, val := range values_UAVIONIX_ADSB_OUT_RF_SELECT {
		if e&val == val {
			names = append(names, value_to_label_UAVIONIX_ADSB_OUT_RF_SELECT[val])
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UAVIONIX_ADSB_OUT_RF_SELECT) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask UAVIONIX_ADSB_OUT_RF_SELECT
	for _, label := range labels {
		if value, ok := label_to_value_UAVIONIX_ADSB_OUT_RF_SELECT[label]; ok {
			mask |= value
		} else if value, err := strconv.Atoi(label); err == nil {
			mask |= UAVIONIX_ADSB_OUT_RF_SELECT(value)
		} else {
			return fmt.Errorf("invalid label '%s'", label)
		}
	}
	*e = mask
	return nil
}

// String implements the fmt.Stringer interface.
func (e UAVIONIX_ADSB_OUT_RF_SELECT) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
