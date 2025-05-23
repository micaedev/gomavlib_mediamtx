//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package development

import (
	"fmt"
	"strconv"
	"strings"
)

// These flags indicate the sensor reporting capabilities for TARGET_ABSOLUTE.
type TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS uint64

const (
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_POSITION     TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS = 1
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_VELOCITY     TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS = 2
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_ACCELERATION TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS = 4
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_ATTITUDE     TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS = 8
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_RATES        TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS = 16
)

var values_TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS = []TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS{
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_POSITION,
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_VELOCITY,
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_ACCELERATION,
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_ATTITUDE,
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_RATES,
}

var value_to_label_TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS = map[TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS]string{
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_POSITION:     "TARGET_ABSOLUTE_SENSOR_CAPABILITY_POSITION",
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_VELOCITY:     "TARGET_ABSOLUTE_SENSOR_CAPABILITY_VELOCITY",
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_ACCELERATION: "TARGET_ABSOLUTE_SENSOR_CAPABILITY_ACCELERATION",
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_ATTITUDE:     "TARGET_ABSOLUTE_SENSOR_CAPABILITY_ATTITUDE",
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_RATES:        "TARGET_ABSOLUTE_SENSOR_CAPABILITY_RATES",
}

var label_to_value_TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS = map[string]TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS{
	"TARGET_ABSOLUTE_SENSOR_CAPABILITY_POSITION":     TARGET_ABSOLUTE_SENSOR_CAPABILITY_POSITION,
	"TARGET_ABSOLUTE_SENSOR_CAPABILITY_VELOCITY":     TARGET_ABSOLUTE_SENSOR_CAPABILITY_VELOCITY,
	"TARGET_ABSOLUTE_SENSOR_CAPABILITY_ACCELERATION": TARGET_ABSOLUTE_SENSOR_CAPABILITY_ACCELERATION,
	"TARGET_ABSOLUTE_SENSOR_CAPABILITY_ATTITUDE":     TARGET_ABSOLUTE_SENSOR_CAPABILITY_ATTITUDE,
	"TARGET_ABSOLUTE_SENSOR_CAPABILITY_RATES":        TARGET_ABSOLUTE_SENSOR_CAPABILITY_RATES,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS) MarshalText() ([]byte, error) {
	if e == 0 {
		return []byte("0"), nil
	}
	var names []string
	for _, val := range values_TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS {
		if e&val == val {
			names = append(names, value_to_label_TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS[val])
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS
	for _, label := range labels {
		if value, ok := label_to_value_TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS[label]; ok {
			mask |= value
		} else if value, err := strconv.Atoi(label); err == nil {
			mask |= TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS(value)
		} else {
			return fmt.Errorf("invalid label '%s'", label)
		}
	}
	*e = mask
	return nil
}

// String implements the fmt.Stringer interface.
func (e TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
