//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"fmt"
	"strconv"
)

type LED_CONTROL_PATTERN uint64

const (
	// LED patterns off (return control to regular vehicle control).
	LED_CONTROL_PATTERN_OFF LED_CONTROL_PATTERN = 0
	// LEDs show pattern during firmware update.
	LED_CONTROL_PATTERN_FIRMWAREUPDATE LED_CONTROL_PATTERN = 1
	// Custom Pattern using custom bytes fields.
	LED_CONTROL_PATTERN_CUSTOM LED_CONTROL_PATTERN = 255
)

var value_to_label_LED_CONTROL_PATTERN = map[LED_CONTROL_PATTERN]string{
	LED_CONTROL_PATTERN_OFF:            "LED_CONTROL_PATTERN_OFF",
	LED_CONTROL_PATTERN_FIRMWAREUPDATE: "LED_CONTROL_PATTERN_FIRMWAREUPDATE",
	LED_CONTROL_PATTERN_CUSTOM:         "LED_CONTROL_PATTERN_CUSTOM",
}

var label_to_value_LED_CONTROL_PATTERN = map[string]LED_CONTROL_PATTERN{
	"LED_CONTROL_PATTERN_OFF":            LED_CONTROL_PATTERN_OFF,
	"LED_CONTROL_PATTERN_FIRMWAREUPDATE": LED_CONTROL_PATTERN_FIRMWAREUPDATE,
	"LED_CONTROL_PATTERN_CUSTOM":         LED_CONTROL_PATTERN_CUSTOM,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e LED_CONTROL_PATTERN) MarshalText() ([]byte, error) {
	if name, ok := value_to_label_LED_CONTROL_PATTERN[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *LED_CONTROL_PATTERN) UnmarshalText(text []byte) error {
	if value, ok := label_to_value_LED_CONTROL_PATTERN[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = LED_CONTROL_PATTERN(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e LED_CONTROL_PATTERN) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
