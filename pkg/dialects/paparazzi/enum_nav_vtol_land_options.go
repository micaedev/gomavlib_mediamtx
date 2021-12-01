//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package paparazzi

import (
	"errors"
)

type NAV_VTOL_LAND_OPTIONS int

const (
	// Default autopilot landing behaviour.
	NAV_VTOL_LAND_OPTIONS_DEFAULT NAV_VTOL_LAND_OPTIONS = 0
	// Descend in fixed wing mode, transitioning to multicopter mode for vertical landing when close to the ground.
	// The fixed wing descent pattern is at the discretion of the vehicle (e.g. transition altitude, loiter direction, radius, and speed, etc.).
	NAV_VTOL_LAND_OPTIONS_FW_DESCENT NAV_VTOL_LAND_OPTIONS = 1
	// Land in multicopter mode on reaching the landing co-ordinates (the whole landing is by "hover descent").
	NAV_VTOL_LAND_OPTIONS_HOVER_DESCENT NAV_VTOL_LAND_OPTIONS = 2
)

var labels_NAV_VTOL_LAND_OPTIONS = map[NAV_VTOL_LAND_OPTIONS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e NAV_VTOL_LAND_OPTIONS) MarshalText() ([]byte, error) {
	if l, ok := labels_NAV_VTOL_LAND_OPTIONS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_NAV_VTOL_LAND_OPTIONS = map[string]NAV_VTOL_LAND_OPTIONS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *NAV_VTOL_LAND_OPTIONS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_NAV_VTOL_LAND_OPTIONS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e NAV_VTOL_LAND_OPTIONS) String() string {
	if l, ok := labels_NAV_VTOL_LAND_OPTIONS[e]; ok {
		return l
	}
	return "invalid value"
}
