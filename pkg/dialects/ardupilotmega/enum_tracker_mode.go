//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"fmt"
	"strconv"
)

// A mapping of antenna tracker flight modes for custom_mode field of heartbeat.
type TRACKER_MODE uint32

const (
	TRACKER_MODE_MANUAL       TRACKER_MODE = 0
	TRACKER_MODE_STOP         TRACKER_MODE = 1
	TRACKER_MODE_SCAN         TRACKER_MODE = 2
	TRACKER_MODE_SERVO_TEST   TRACKER_MODE = 3
	TRACKER_MODE_AUTO         TRACKER_MODE = 10
	TRACKER_MODE_INITIALIZING TRACKER_MODE = 16
)

var labels_TRACKER_MODE = map[TRACKER_MODE]string{
	TRACKER_MODE_MANUAL:       "TRACKER_MODE_MANUAL",
	TRACKER_MODE_STOP:         "TRACKER_MODE_STOP",
	TRACKER_MODE_SCAN:         "TRACKER_MODE_SCAN",
	TRACKER_MODE_SERVO_TEST:   "TRACKER_MODE_SERVO_TEST",
	TRACKER_MODE_AUTO:         "TRACKER_MODE_AUTO",
	TRACKER_MODE_INITIALIZING: "TRACKER_MODE_INITIALIZING",
}

var values_TRACKER_MODE = map[string]TRACKER_MODE{
	"TRACKER_MODE_MANUAL":       TRACKER_MODE_MANUAL,
	"TRACKER_MODE_STOP":         TRACKER_MODE_STOP,
	"TRACKER_MODE_SCAN":         TRACKER_MODE_SCAN,
	"TRACKER_MODE_SERVO_TEST":   TRACKER_MODE_SERVO_TEST,
	"TRACKER_MODE_AUTO":         TRACKER_MODE_AUTO,
	"TRACKER_MODE_INITIALIZING": TRACKER_MODE_INITIALIZING,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e TRACKER_MODE) MarshalText() ([]byte, error) {
	if name, ok := labels_TRACKER_MODE[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *TRACKER_MODE) UnmarshalText(text []byte) error {
	if value, ok := values_TRACKER_MODE[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = TRACKER_MODE(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e TRACKER_MODE) String() string {
	if name, ok := labels_TRACKER_MODE[e]; ok {
		return name
	}
	return strconv.Itoa(int(e))
}
