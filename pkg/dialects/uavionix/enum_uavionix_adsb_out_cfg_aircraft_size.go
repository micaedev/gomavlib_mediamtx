//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"fmt"
	"strconv"
)

// Definitions for aircraft size
type UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE uint32

const (
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_NO_DATA     UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = 0
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L15M_W23M   UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = 1
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L25M_W28P5M UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = 2
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L25_34M     UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = 3
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L35_33M     UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = 4
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L35_38M     UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = 5
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L45_39P5M   UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = 6
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L45_45M     UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = 7
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L55_45M     UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = 8
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L55_52M     UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = 9
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L65_59P5M   UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = 10
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L65_67M     UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = 11
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L75_W72P5M  UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = 12
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L75_W80M    UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = 13
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L85_W80M    UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = 14
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L85_W90M    UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = 15
)

var labels_UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = map[UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE]string{
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_NO_DATA:     "UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_NO_DATA",
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L15M_W23M:   "UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L15M_W23M",
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L25M_W28P5M: "UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L25M_W28P5M",
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L25_34M:     "UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L25_34M",
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L35_33M:     "UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L35_33M",
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L35_38M:     "UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L35_38M",
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L45_39P5M:   "UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L45_39P5M",
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L45_45M:     "UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L45_45M",
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L55_45M:     "UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L55_45M",
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L55_52M:     "UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L55_52M",
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L65_59P5M:   "UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L65_59P5M",
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L65_67M:     "UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L65_67M",
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L75_W72P5M:  "UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L75_W72P5M",
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L75_W80M:    "UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L75_W80M",
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L85_W80M:    "UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L85_W80M",
	UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L85_W90M:    "UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L85_W90M",
}

var values_UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE = map[string]UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE{
	"UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_NO_DATA":     UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_NO_DATA,
	"UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L15M_W23M":   UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L15M_W23M,
	"UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L25M_W28P5M": UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L25M_W28P5M,
	"UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L25_34M":     UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L25_34M,
	"UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L35_33M":     UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L35_33M,
	"UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L35_38M":     UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L35_38M,
	"UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L45_39P5M":   UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L45_39P5M,
	"UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L45_45M":     UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L45_45M,
	"UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L55_45M":     UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L55_45M,
	"UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L55_52M":     UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L55_52M,
	"UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L65_59P5M":   UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L65_59P5M,
	"UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L65_67M":     UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L65_67M,
	"UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L75_W72P5M":  UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L75_W72P5M,
	"UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L75_W80M":    UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L75_W80M,
	"UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L85_W80M":    UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L85_W80M,
	"UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L85_W90M":    UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE_L85_W90M,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE) MarshalText() ([]byte, error) {
	if name, ok := labels_UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE) UnmarshalText(text []byte) error {
	if value, ok := values_UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE) String() string {
	if name, ok := labels_UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE[e]; ok {
		return name
	}
	return strconv.Itoa(int(e))
}
