//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package avssuas

import (
	"errors"
)

// Enumeration of distance sensor types
type MAV_DISTANCE_SENSOR uint32

const (
	// Laser rangefinder, e.g. LightWare SF02/F or PulsedLight units
	MAV_DISTANCE_SENSOR_LASER MAV_DISTANCE_SENSOR = 0
	// Ultrasound rangefinder, e.g. MaxBotix units
	MAV_DISTANCE_SENSOR_ULTRASOUND MAV_DISTANCE_SENSOR = 1
	// Infrared rangefinder, e.g. Sharp units
	MAV_DISTANCE_SENSOR_INFRARED MAV_DISTANCE_SENSOR = 2
	// Radar type, e.g. uLanding units
	MAV_DISTANCE_SENSOR_RADAR MAV_DISTANCE_SENSOR = 3
	// Broken or unknown type, e.g. analog units
	MAV_DISTANCE_SENSOR_UNKNOWN MAV_DISTANCE_SENSOR = 4
)

var labels_MAV_DISTANCE_SENSOR = map[MAV_DISTANCE_SENSOR]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_DISTANCE_SENSOR) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_DISTANCE_SENSOR[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_DISTANCE_SENSOR = map[string]MAV_DISTANCE_SENSOR{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_DISTANCE_SENSOR) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_DISTANCE_SENSOR[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_DISTANCE_SENSOR) String() string {
	if l, ok := labels_MAV_DISTANCE_SENSOR[e]; ok {
		return l
	}
	return "invalid value"
}
