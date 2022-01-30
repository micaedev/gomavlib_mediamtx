//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

// GPS longitudinal offset encoding
type UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LON uint32

const (
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LON_NO_DATA           UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LON = 0
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LON_APPLIED_BY_SENSOR UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LON = 1
)

var labels_UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LON = map[UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LON]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LON) MarshalText() ([]byte, error) {
	if l, ok := labels_UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LON[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LON = map[string]UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LON{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LON) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LON[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LON) String() string {
	if l, ok := labels_UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LON[e]; ok {
		return l
	}
	return "invalid value"
}
