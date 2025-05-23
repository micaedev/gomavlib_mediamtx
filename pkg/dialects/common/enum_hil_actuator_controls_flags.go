//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strconv"
	"strings"
)

// Flags used in HIL_ACTUATOR_CONTROLS message.
type HIL_ACTUATOR_CONTROLS_FLAGS uint64

const (
	// Simulation is using lockstep
	HIL_ACTUATOR_CONTROLS_FLAGS_LOCKSTEP HIL_ACTUATOR_CONTROLS_FLAGS = 1
)

var values_HIL_ACTUATOR_CONTROLS_FLAGS = []HIL_ACTUATOR_CONTROLS_FLAGS{
	HIL_ACTUATOR_CONTROLS_FLAGS_LOCKSTEP,
}

var value_to_label_HIL_ACTUATOR_CONTROLS_FLAGS = map[HIL_ACTUATOR_CONTROLS_FLAGS]string{
	HIL_ACTUATOR_CONTROLS_FLAGS_LOCKSTEP: "HIL_ACTUATOR_CONTROLS_FLAGS_LOCKSTEP",
}

var label_to_value_HIL_ACTUATOR_CONTROLS_FLAGS = map[string]HIL_ACTUATOR_CONTROLS_FLAGS{
	"HIL_ACTUATOR_CONTROLS_FLAGS_LOCKSTEP": HIL_ACTUATOR_CONTROLS_FLAGS_LOCKSTEP,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e HIL_ACTUATOR_CONTROLS_FLAGS) MarshalText() ([]byte, error) {
	if e == 0 {
		return []byte("0"), nil
	}
	var names []string
	for _, val := range values_HIL_ACTUATOR_CONTROLS_FLAGS {
		if e&val == val {
			names = append(names, value_to_label_HIL_ACTUATOR_CONTROLS_FLAGS[val])
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *HIL_ACTUATOR_CONTROLS_FLAGS) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask HIL_ACTUATOR_CONTROLS_FLAGS
	for _, label := range labels {
		if value, ok := label_to_value_HIL_ACTUATOR_CONTROLS_FLAGS[label]; ok {
			mask |= value
		} else if value, err := strconv.Atoi(label); err == nil {
			mask |= HIL_ACTUATOR_CONTROLS_FLAGS(value)
		} else {
			return fmt.Errorf("invalid label '%s'", label)
		}
	}
	*e = mask
	return nil
}

// String implements the fmt.Stringer interface.
func (e HIL_ACTUATOR_CONTROLS_FLAGS) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
