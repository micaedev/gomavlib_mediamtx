//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package avssuas

import (
	"errors"
)

// MAVLINK component type reported in HEARTBEAT message. Flight controllers must report the type of the vehicle on which they are mounted (e.g. MAV_TYPE_OCTOROTOR). All other components must report a value appropriate for their type (e.g. a camera must use MAV_TYPE_CAMERA).
type MAV_TYPE uint32

const (
	// Generic micro air vehicle
	MAV_TYPE_GENERIC MAV_TYPE = 0
	// Fixed wing aircraft.
	MAV_TYPE_FIXED_WING MAV_TYPE = 1
	// Quadrotor
	MAV_TYPE_QUADROTOR MAV_TYPE = 2
	// Coaxial helicopter
	MAV_TYPE_COAXIAL MAV_TYPE = 3
	// Normal helicopter with tail rotor.
	MAV_TYPE_HELICOPTER MAV_TYPE = 4
	// Ground installation
	MAV_TYPE_ANTENNA_TRACKER MAV_TYPE = 5
	// Operator control unit / ground control station
	MAV_TYPE_GCS MAV_TYPE = 6
	// Airship, controlled
	MAV_TYPE_AIRSHIP MAV_TYPE = 7
	// Free balloon, uncontrolled
	MAV_TYPE_FREE_BALLOON MAV_TYPE = 8
	// Rocket
	MAV_TYPE_ROCKET MAV_TYPE = 9
	// Ground rover
	MAV_TYPE_GROUND_ROVER MAV_TYPE = 10
	// Surface vessel, boat, ship
	MAV_TYPE_SURFACE_BOAT MAV_TYPE = 11
	// Submarine
	MAV_TYPE_SUBMARINE MAV_TYPE = 12
	// Hexarotor
	MAV_TYPE_HEXAROTOR MAV_TYPE = 13
	// Octorotor
	MAV_TYPE_OCTOROTOR MAV_TYPE = 14
	// Tricopter
	MAV_TYPE_TRICOPTER MAV_TYPE = 15
	// Flapping wing
	MAV_TYPE_FLAPPING_WING MAV_TYPE = 16
	// Kite
	MAV_TYPE_KITE MAV_TYPE = 17
	// Onboard companion controller
	MAV_TYPE_ONBOARD_CONTROLLER MAV_TYPE = 18
	// Two-rotor VTOL using control surfaces in vertical operation in addition. Tailsitter.
	MAV_TYPE_VTOL_DUOROTOR MAV_TYPE = 19
	// Quad-rotor VTOL using a V-shaped quad config in vertical operation. Tailsitter.
	MAV_TYPE_VTOL_QUADROTOR MAV_TYPE = 20
	// Tiltrotor VTOL
	MAV_TYPE_VTOL_TILTROTOR MAV_TYPE = 21
	// VTOL reserved 2
	MAV_TYPE_VTOL_RESERVED2 MAV_TYPE = 22
	// VTOL reserved 3
	MAV_TYPE_VTOL_RESERVED3 MAV_TYPE = 23
	// VTOL reserved 4
	MAV_TYPE_VTOL_RESERVED4 MAV_TYPE = 24
	// VTOL reserved 5
	MAV_TYPE_VTOL_RESERVED5 MAV_TYPE = 25
	// Gimbal
	MAV_TYPE_GIMBAL MAV_TYPE = 26
	// ADSB system
	MAV_TYPE_ADSB MAV_TYPE = 27
	// Steerable, nonrigid airfoil
	MAV_TYPE_PARAFOIL MAV_TYPE = 28
	// Dodecarotor
	MAV_TYPE_DODECAROTOR MAV_TYPE = 29
	// Camera
	MAV_TYPE_CAMERA MAV_TYPE = 30
	// Charging station
	MAV_TYPE_CHARGING_STATION MAV_TYPE = 31
	// FLARM collision avoidance system
	MAV_TYPE_FLARM MAV_TYPE = 32
	// Servo
	MAV_TYPE_SERVO MAV_TYPE = 33
	// Open Drone ID. See https://mavlink.io/en/services/opendroneid.html.
	MAV_TYPE_ODID MAV_TYPE = 34
	// Decarotor
	MAV_TYPE_DECAROTOR MAV_TYPE = 35
	// Battery
	MAV_TYPE_BATTERY MAV_TYPE = 36
	// Parachute
	MAV_TYPE_PARACHUTE MAV_TYPE = 37
	// Log
	MAV_TYPE_LOG MAV_TYPE = 38
	// OSD
	MAV_TYPE_OSD MAV_TYPE = 39
	// IMU
	MAV_TYPE_IMU MAV_TYPE = 40
	// GPS
	MAV_TYPE_GPS MAV_TYPE = 41
)

var labels_MAV_TYPE = map[MAV_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_TYPE = map[string]MAV_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_TYPE) String() string {
	if l, ok := labels_MAV_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
