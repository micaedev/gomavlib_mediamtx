//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package ualberta

import (
	"errors"
)

// Navigation filter mode
type UALBERTA_NAV_MODE int

const (
	NAV_AHRS_INIT UALBERTA_NAV_MODE = 1
	// AHRS mode
	NAV_AHRS UALBERTA_NAV_MODE = 2
	// INS/GPS initialization mode
	NAV_INS_GPS_INIT UALBERTA_NAV_MODE = 3
	// INS/GPS mode
	NAV_INS_GPS UALBERTA_NAV_MODE = 4
)

var labels_UALBERTA_NAV_MODE = map[UALBERTA_NAV_MODE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UALBERTA_NAV_MODE) MarshalText() ([]byte, error) {
	if l, ok := labels_UALBERTA_NAV_MODE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_UALBERTA_NAV_MODE = map[string]UALBERTA_NAV_MODE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UALBERTA_NAV_MODE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_UALBERTA_NAV_MODE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e UALBERTA_NAV_MODE) String() string {
	if l, ok := labels_UALBERTA_NAV_MODE[e]; ok {
		return l
	}
	return "invalid value"
}
