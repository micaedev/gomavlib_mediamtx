//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package standard

import (
	"errors"
)

// Power supply status flags (bitmask)
type MAV_POWER_STATUS uint32

const (
	// main brick power supply valid
	MAV_POWER_STATUS_BRICK_VALID MAV_POWER_STATUS = 1
	// main servo power supply valid for FMU
	MAV_POWER_STATUS_SERVO_VALID MAV_POWER_STATUS = 2
	// USB power is connected
	MAV_POWER_STATUS_USB_CONNECTED MAV_POWER_STATUS = 4
	// peripheral supply is in over-current state
	MAV_POWER_STATUS_PERIPH_OVERCURRENT MAV_POWER_STATUS = 8
	// hi-power peripheral supply is in over-current state
	MAV_POWER_STATUS_PERIPH_HIPOWER_OVERCURRENT MAV_POWER_STATUS = 16
	// Power status has changed since boot
	MAV_POWER_STATUS_CHANGED MAV_POWER_STATUS = 32
)

var labels_MAV_POWER_STATUS = map[MAV_POWER_STATUS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_POWER_STATUS) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_POWER_STATUS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_POWER_STATUS = map[string]MAV_POWER_STATUS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_POWER_STATUS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_POWER_STATUS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_POWER_STATUS) String() string {
	if l, ok := labels_MAV_POWER_STATUS[e]; ok {
		return l
	}
	return "invalid value"
}
