//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package all

import (
	"errors"
)

// Enumeration of estimator types
type MAV_ESTIMATOR_TYPE uint32

const (
	// Unknown type of the estimator.
	MAV_ESTIMATOR_TYPE_UNKNOWN MAV_ESTIMATOR_TYPE = 0
	// This is a naive estimator without any real covariance feedback.
	MAV_ESTIMATOR_TYPE_NAIVE MAV_ESTIMATOR_TYPE = 1
	// Computer vision based estimate. Might be up to scale.
	MAV_ESTIMATOR_TYPE_VISION MAV_ESTIMATOR_TYPE = 2
	// Visual-inertial estimate.
	MAV_ESTIMATOR_TYPE_VIO MAV_ESTIMATOR_TYPE = 3
	// Plain GPS estimate.
	MAV_ESTIMATOR_TYPE_GPS MAV_ESTIMATOR_TYPE = 4
	// Estimator integrating GPS and inertial sensing.
	MAV_ESTIMATOR_TYPE_GPS_INS MAV_ESTIMATOR_TYPE = 5
	// Estimate from external motion capturing system.
	MAV_ESTIMATOR_TYPE_MOCAP MAV_ESTIMATOR_TYPE = 6
	// Estimator based on lidar sensor input.
	MAV_ESTIMATOR_TYPE_LIDAR MAV_ESTIMATOR_TYPE = 7
	// Estimator on autopilot.
	MAV_ESTIMATOR_TYPE_AUTOPILOT MAV_ESTIMATOR_TYPE = 8
)

var labels_MAV_ESTIMATOR_TYPE = map[MAV_ESTIMATOR_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ESTIMATOR_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ESTIMATOR_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ESTIMATOR_TYPE = map[string]MAV_ESTIMATOR_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ESTIMATOR_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ESTIMATOR_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ESTIMATOR_TYPE) String() string {
	if l, ok := labels_MAV_ESTIMATOR_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
