//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

// Flags to report ESC failures.
type ESC_FAILURE_FLAGS uint32

const (
	// No ESC failure.
	ESC_FAILURE_NONE ESC_FAILURE_FLAGS = 0
	// Over current failure.
	ESC_FAILURE_OVER_CURRENT ESC_FAILURE_FLAGS = 1
	// Over voltage failure.
	ESC_FAILURE_OVER_VOLTAGE ESC_FAILURE_FLAGS = 2
	// Over temperature failure.
	ESC_FAILURE_OVER_TEMPERATURE ESC_FAILURE_FLAGS = 4
	// Over RPM failure.
	ESC_FAILURE_OVER_RPM ESC_FAILURE_FLAGS = 8
	// Inconsistent command failure i.e. out of bounds.
	ESC_FAILURE_INCONSISTENT_CMD ESC_FAILURE_FLAGS = 16
	// Motor stuck failure.
	ESC_FAILURE_MOTOR_STUCK ESC_FAILURE_FLAGS = 32
	// Generic ESC failure.
	ESC_FAILURE_GENERIC ESC_FAILURE_FLAGS = 64
)

var labels_ESC_FAILURE_FLAGS = map[ESC_FAILURE_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ESC_FAILURE_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_ESC_FAILURE_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_ESC_FAILURE_FLAGS = map[string]ESC_FAILURE_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ESC_FAILURE_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_ESC_FAILURE_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e ESC_FAILURE_FLAGS) String() string {
	if l, ok := labels_ESC_FAILURE_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
