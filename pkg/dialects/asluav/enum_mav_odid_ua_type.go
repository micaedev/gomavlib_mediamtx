//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package asluav

import (
	"errors"
)

type MAV_ODID_UA_TYPE uint32

const (
	// No UA (Unmanned Aircraft) type defined.
	MAV_ODID_UA_TYPE_NONE MAV_ODID_UA_TYPE = 0
	// Aeroplane/Airplane. Fixed wing.
	MAV_ODID_UA_TYPE_AEROPLANE MAV_ODID_UA_TYPE = 1
	// Helicopter or multirotor.
	MAV_ODID_UA_TYPE_HELICOPTER_OR_MULTIROTOR MAV_ODID_UA_TYPE = 2
	// Gyroplane.
	MAV_ODID_UA_TYPE_GYROPLANE MAV_ODID_UA_TYPE = 3
	// VTOL (Vertical Take-Off and Landing). Fixed wing aircraft that can take off vertically.
	MAV_ODID_UA_TYPE_HYBRID_LIFT MAV_ODID_UA_TYPE = 4
	// Ornithopter.
	MAV_ODID_UA_TYPE_ORNITHOPTER MAV_ODID_UA_TYPE = 5
	// Glider.
	MAV_ODID_UA_TYPE_GLIDER MAV_ODID_UA_TYPE = 6
	// Kite.
	MAV_ODID_UA_TYPE_KITE MAV_ODID_UA_TYPE = 7
	// Free Balloon.
	MAV_ODID_UA_TYPE_FREE_BALLOON MAV_ODID_UA_TYPE = 8
	// Captive Balloon.
	MAV_ODID_UA_TYPE_CAPTIVE_BALLOON MAV_ODID_UA_TYPE = 9
	// Airship. E.g. a blimp.
	MAV_ODID_UA_TYPE_AIRSHIP MAV_ODID_UA_TYPE = 10
	// Free Fall/Parachute (unpowered).
	MAV_ODID_UA_TYPE_FREE_FALL_PARACHUTE MAV_ODID_UA_TYPE = 11
	// Rocket.
	MAV_ODID_UA_TYPE_ROCKET MAV_ODID_UA_TYPE = 12
	// Tethered powered aircraft.
	MAV_ODID_UA_TYPE_TETHERED_POWERED_AIRCRAFT MAV_ODID_UA_TYPE = 13
	// Ground Obstacle.
	MAV_ODID_UA_TYPE_GROUND_OBSTACLE MAV_ODID_UA_TYPE = 14
	// Other type of aircraft not listed earlier.
	MAV_ODID_UA_TYPE_OTHER MAV_ODID_UA_TYPE = 15
)

var labels_MAV_ODID_UA_TYPE = map[MAV_ODID_UA_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_UA_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ODID_UA_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ODID_UA_TYPE = map[string]MAV_ODID_UA_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_UA_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ODID_UA_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_UA_TYPE) String() string {
	if l, ok := labels_MAV_ODID_UA_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
