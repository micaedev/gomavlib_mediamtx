//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package common

import (
	"errors"
)

// The ROI (region of interest) for the vehicle. This can be
// be used by the vehicle for camera/vehicle attitude alignment (see
// MAV_CMD_NAV_ROI).
type MAV_ROI int

const (
	// No region of interest.
	MAV_ROI_NONE MAV_ROI = 0
	// Point toward next waypoint, with optional pitch/roll/yaw offset.
	MAV_ROI_WPNEXT MAV_ROI = 1
	// Point toward given waypoint.
	MAV_ROI_WPINDEX MAV_ROI = 2
	// Point toward fixed location.
	MAV_ROI_LOCATION MAV_ROI = 3
	// Point toward of given id.
	MAV_ROI_TARGET MAV_ROI = 4
)

var labels_MAV_ROI = map[MAV_ROI]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ROI) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ROI[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ROI = map[string]MAV_ROI{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ROI) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ROI[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ROI) String() string {
	if l, ok := labels_MAV_ROI[e]; ok {
		return l
	}
	return "invalid value"
}
