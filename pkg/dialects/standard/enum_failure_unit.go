//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package standard

import (
	"errors"
)

// List of possible units where failures can be injected.
type FAILURE_UNIT uint32

const (
	FAILURE_UNIT_SENSOR_GYRO            FAILURE_UNIT = 0
	FAILURE_UNIT_SENSOR_ACCEL           FAILURE_UNIT = 1
	FAILURE_UNIT_SENSOR_MAG             FAILURE_UNIT = 2
	FAILURE_UNIT_SENSOR_BARO            FAILURE_UNIT = 3
	FAILURE_UNIT_SENSOR_GPS             FAILURE_UNIT = 4
	FAILURE_UNIT_SENSOR_OPTICAL_FLOW    FAILURE_UNIT = 5
	FAILURE_UNIT_SENSOR_VIO             FAILURE_UNIT = 6
	FAILURE_UNIT_SENSOR_DISTANCE_SENSOR FAILURE_UNIT = 7
	FAILURE_UNIT_SENSOR_AIRSPEED        FAILURE_UNIT = 8
	FAILURE_UNIT_SYSTEM_BATTERY         FAILURE_UNIT = 100
	FAILURE_UNIT_SYSTEM_MOTOR           FAILURE_UNIT = 101
	FAILURE_UNIT_SYSTEM_SERVO           FAILURE_UNIT = 102
	FAILURE_UNIT_SYSTEM_AVOIDANCE       FAILURE_UNIT = 103
	FAILURE_UNIT_SYSTEM_RC_SIGNAL       FAILURE_UNIT = 104
	FAILURE_UNIT_SYSTEM_MAVLINK_SIGNAL  FAILURE_UNIT = 105
)

var labels_FAILURE_UNIT = map[FAILURE_UNIT]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e FAILURE_UNIT) MarshalText() ([]byte, error) {
	if l, ok := labels_FAILURE_UNIT[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_FAILURE_UNIT = map[string]FAILURE_UNIT{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *FAILURE_UNIT) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_FAILURE_UNIT[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e FAILURE_UNIT) String() string {
	if l, ok := labels_FAILURE_UNIT[e]; ok {
		return l
	}
	return "invalid value"
}
