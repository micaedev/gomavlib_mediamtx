//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package development

import (
	"errors"
)

// Flags to report failure cases over the high latency telemtry.
type HL_FAILURE_FLAG uint32

const (
	// GPS failure.
	HL_FAILURE_FLAG_GPS HL_FAILURE_FLAG = 1
	// Differential pressure sensor failure.
	HL_FAILURE_FLAG_DIFFERENTIAL_PRESSURE HL_FAILURE_FLAG = 2
	// Absolute pressure sensor failure.
	HL_FAILURE_FLAG_ABSOLUTE_PRESSURE HL_FAILURE_FLAG = 4
	// Accelerometer sensor failure.
	HL_FAILURE_FLAG_3D_ACCEL HL_FAILURE_FLAG = 8
	// Gyroscope sensor failure.
	HL_FAILURE_FLAG_3D_GYRO HL_FAILURE_FLAG = 16
	// Magnetometer sensor failure.
	HL_FAILURE_FLAG_3D_MAG HL_FAILURE_FLAG = 32
	// Terrain subsystem failure.
	HL_FAILURE_FLAG_TERRAIN HL_FAILURE_FLAG = 64
	// Battery failure/critical low battery.
	HL_FAILURE_FLAG_BATTERY HL_FAILURE_FLAG = 128
	// RC receiver failure/no rc connection.
	HL_FAILURE_FLAG_RC_RECEIVER HL_FAILURE_FLAG = 256
	// Offboard link failure.
	HL_FAILURE_FLAG_OFFBOARD_LINK HL_FAILURE_FLAG = 512
	// Engine failure.
	HL_FAILURE_FLAG_ENGINE HL_FAILURE_FLAG = 1024
	// Geofence violation.
	HL_FAILURE_FLAG_GEOFENCE HL_FAILURE_FLAG = 2048
	// Estimator failure, for example measurement rejection or large variances.
	HL_FAILURE_FLAG_ESTIMATOR HL_FAILURE_FLAG = 4096
	// Mission failure.
	HL_FAILURE_FLAG_MISSION HL_FAILURE_FLAG = 8192
)

var labels_HL_FAILURE_FLAG = map[HL_FAILURE_FLAG]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e HL_FAILURE_FLAG) MarshalText() ([]byte, error) {
	if l, ok := labels_HL_FAILURE_FLAG[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_HL_FAILURE_FLAG = map[string]HL_FAILURE_FLAG{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *HL_FAILURE_FLAG) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_HL_FAILURE_FLAG[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e HL_FAILURE_FLAG) String() string {
	if l, ok := labels_HL_FAILURE_FLAG[e]; ok {
		return l
	}
	return "invalid value"
}
