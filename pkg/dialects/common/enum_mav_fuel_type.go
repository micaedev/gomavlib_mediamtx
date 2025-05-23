//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strconv"
)

// Fuel types for use in FUEL_TYPE. Fuel types specify the units for the maximum, available and consumed fuel, and for the flow rates.
type MAV_FUEL_TYPE uint64

const (
	// Not specified. Fuel levels are normalized (i.e. maximum is 1, and other levels are relative to 1).
	MAV_FUEL_TYPE_UNKNOWN MAV_FUEL_TYPE = 0
	// A generic liquid fuel. Fuel levels are in millilitres (ml). Fuel rates are in millilitres/second.
	MAV_FUEL_TYPE_LIQUID MAV_FUEL_TYPE = 1
	// A gas tank. Fuel levels are in kilo-Pascal (kPa), and flow rates are in milliliters per second (ml/s).
	MAV_FUEL_TYPE_GAS MAV_FUEL_TYPE = 2
)

var value_to_label_MAV_FUEL_TYPE = map[MAV_FUEL_TYPE]string{
	MAV_FUEL_TYPE_UNKNOWN: "MAV_FUEL_TYPE_UNKNOWN",
	MAV_FUEL_TYPE_LIQUID:  "MAV_FUEL_TYPE_LIQUID",
	MAV_FUEL_TYPE_GAS:     "MAV_FUEL_TYPE_GAS",
}

var label_to_value_MAV_FUEL_TYPE = map[string]MAV_FUEL_TYPE{
	"MAV_FUEL_TYPE_UNKNOWN": MAV_FUEL_TYPE_UNKNOWN,
	"MAV_FUEL_TYPE_LIQUID":  MAV_FUEL_TYPE_LIQUID,
	"MAV_FUEL_TYPE_GAS":     MAV_FUEL_TYPE_GAS,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_FUEL_TYPE) MarshalText() ([]byte, error) {
	if name, ok := value_to_label_MAV_FUEL_TYPE[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_FUEL_TYPE) UnmarshalText(text []byte) error {
	if value, ok := label_to_value_MAV_FUEL_TYPE[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = MAV_FUEL_TYPE(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e MAV_FUEL_TYPE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
