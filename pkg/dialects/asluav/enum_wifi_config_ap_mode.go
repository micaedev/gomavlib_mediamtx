//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package asluav

import (
	"errors"
)

// WiFi Mode.
type WIFI_CONFIG_AP_MODE uint32

const (
	// WiFi mode is undefined.
	WIFI_CONFIG_AP_MODE_UNDEFINED WIFI_CONFIG_AP_MODE = 0
	// WiFi configured as an access point.
	WIFI_CONFIG_AP_MODE_AP WIFI_CONFIG_AP_MODE = 1
	// WiFi configured as a station connected to an existing local WiFi network.
	WIFI_CONFIG_AP_MODE_STATION WIFI_CONFIG_AP_MODE = 2
	// WiFi disabled.
	WIFI_CONFIG_AP_MODE_DISABLED WIFI_CONFIG_AP_MODE = 3
)

var labels_WIFI_CONFIG_AP_MODE = map[WIFI_CONFIG_AP_MODE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e WIFI_CONFIG_AP_MODE) MarshalText() ([]byte, error) {
	if l, ok := labels_WIFI_CONFIG_AP_MODE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_WIFI_CONFIG_AP_MODE = map[string]WIFI_CONFIG_AP_MODE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *WIFI_CONFIG_AP_MODE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_WIFI_CONFIG_AP_MODE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e WIFI_CONFIG_AP_MODE) String() string {
	if l, ok := labels_WIFI_CONFIG_AP_MODE[e]; ok {
		return l
	}
	return "invalid value"
}
