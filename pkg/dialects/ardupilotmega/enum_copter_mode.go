//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"fmt"
	"strconv"
)

// A mapping of copter flight modes for custom_mode field of heartbeat.
type COPTER_MODE uint32

const (
	COPTER_MODE_STABILIZE    COPTER_MODE = 0
	COPTER_MODE_ACRO         COPTER_MODE = 1
	COPTER_MODE_ALT_HOLD     COPTER_MODE = 2
	COPTER_MODE_AUTO         COPTER_MODE = 3
	COPTER_MODE_GUIDED       COPTER_MODE = 4
	COPTER_MODE_LOITER       COPTER_MODE = 5
	COPTER_MODE_RTL          COPTER_MODE = 6
	COPTER_MODE_CIRCLE       COPTER_MODE = 7
	COPTER_MODE_LAND         COPTER_MODE = 9
	COPTER_MODE_DRIFT        COPTER_MODE = 11
	COPTER_MODE_SPORT        COPTER_MODE = 13
	COPTER_MODE_FLIP         COPTER_MODE = 14
	COPTER_MODE_AUTOTUNE     COPTER_MODE = 15
	COPTER_MODE_POSHOLD      COPTER_MODE = 16
	COPTER_MODE_BRAKE        COPTER_MODE = 17
	COPTER_MODE_THROW        COPTER_MODE = 18
	COPTER_MODE_AVOID_ADSB   COPTER_MODE = 19
	COPTER_MODE_GUIDED_NOGPS COPTER_MODE = 20
	COPTER_MODE_SMART_RTL    COPTER_MODE = 21
	COPTER_MODE_FLOWHOLD     COPTER_MODE = 22
	COPTER_MODE_FOLLOW       COPTER_MODE = 23
	COPTER_MODE_ZIGZAG       COPTER_MODE = 24
	COPTER_MODE_SYSTEMID     COPTER_MODE = 25
	COPTER_MODE_AUTOROTATE   COPTER_MODE = 26
	COPTER_MODE_AUTO_RTL     COPTER_MODE = 27
)

var labels_COPTER_MODE = map[COPTER_MODE]string{
	COPTER_MODE_STABILIZE:    "COPTER_MODE_STABILIZE",
	COPTER_MODE_ACRO:         "COPTER_MODE_ACRO",
	COPTER_MODE_ALT_HOLD:     "COPTER_MODE_ALT_HOLD",
	COPTER_MODE_AUTO:         "COPTER_MODE_AUTO",
	COPTER_MODE_GUIDED:       "COPTER_MODE_GUIDED",
	COPTER_MODE_LOITER:       "COPTER_MODE_LOITER",
	COPTER_MODE_RTL:          "COPTER_MODE_RTL",
	COPTER_MODE_CIRCLE:       "COPTER_MODE_CIRCLE",
	COPTER_MODE_LAND:         "COPTER_MODE_LAND",
	COPTER_MODE_DRIFT:        "COPTER_MODE_DRIFT",
	COPTER_MODE_SPORT:        "COPTER_MODE_SPORT",
	COPTER_MODE_FLIP:         "COPTER_MODE_FLIP",
	COPTER_MODE_AUTOTUNE:     "COPTER_MODE_AUTOTUNE",
	COPTER_MODE_POSHOLD:      "COPTER_MODE_POSHOLD",
	COPTER_MODE_BRAKE:        "COPTER_MODE_BRAKE",
	COPTER_MODE_THROW:        "COPTER_MODE_THROW",
	COPTER_MODE_AVOID_ADSB:   "COPTER_MODE_AVOID_ADSB",
	COPTER_MODE_GUIDED_NOGPS: "COPTER_MODE_GUIDED_NOGPS",
	COPTER_MODE_SMART_RTL:    "COPTER_MODE_SMART_RTL",
	COPTER_MODE_FLOWHOLD:     "COPTER_MODE_FLOWHOLD",
	COPTER_MODE_FOLLOW:       "COPTER_MODE_FOLLOW",
	COPTER_MODE_ZIGZAG:       "COPTER_MODE_ZIGZAG",
	COPTER_MODE_SYSTEMID:     "COPTER_MODE_SYSTEMID",
	COPTER_MODE_AUTOROTATE:   "COPTER_MODE_AUTOROTATE",
	COPTER_MODE_AUTO_RTL:     "COPTER_MODE_AUTO_RTL",
}

var values_COPTER_MODE = map[string]COPTER_MODE{
	"COPTER_MODE_STABILIZE":    COPTER_MODE_STABILIZE,
	"COPTER_MODE_ACRO":         COPTER_MODE_ACRO,
	"COPTER_MODE_ALT_HOLD":     COPTER_MODE_ALT_HOLD,
	"COPTER_MODE_AUTO":         COPTER_MODE_AUTO,
	"COPTER_MODE_GUIDED":       COPTER_MODE_GUIDED,
	"COPTER_MODE_LOITER":       COPTER_MODE_LOITER,
	"COPTER_MODE_RTL":          COPTER_MODE_RTL,
	"COPTER_MODE_CIRCLE":       COPTER_MODE_CIRCLE,
	"COPTER_MODE_LAND":         COPTER_MODE_LAND,
	"COPTER_MODE_DRIFT":        COPTER_MODE_DRIFT,
	"COPTER_MODE_SPORT":        COPTER_MODE_SPORT,
	"COPTER_MODE_FLIP":         COPTER_MODE_FLIP,
	"COPTER_MODE_AUTOTUNE":     COPTER_MODE_AUTOTUNE,
	"COPTER_MODE_POSHOLD":      COPTER_MODE_POSHOLD,
	"COPTER_MODE_BRAKE":        COPTER_MODE_BRAKE,
	"COPTER_MODE_THROW":        COPTER_MODE_THROW,
	"COPTER_MODE_AVOID_ADSB":   COPTER_MODE_AVOID_ADSB,
	"COPTER_MODE_GUIDED_NOGPS": COPTER_MODE_GUIDED_NOGPS,
	"COPTER_MODE_SMART_RTL":    COPTER_MODE_SMART_RTL,
	"COPTER_MODE_FLOWHOLD":     COPTER_MODE_FLOWHOLD,
	"COPTER_MODE_FOLLOW":       COPTER_MODE_FOLLOW,
	"COPTER_MODE_ZIGZAG":       COPTER_MODE_ZIGZAG,
	"COPTER_MODE_SYSTEMID":     COPTER_MODE_SYSTEMID,
	"COPTER_MODE_AUTOROTATE":   COPTER_MODE_AUTOROTATE,
	"COPTER_MODE_AUTO_RTL":     COPTER_MODE_AUTO_RTL,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e COPTER_MODE) MarshalText() ([]byte, error) {
	if name, ok := labels_COPTER_MODE[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *COPTER_MODE) UnmarshalText(text []byte) error {
	if value, ok := values_COPTER_MODE[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = COPTER_MODE(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e COPTER_MODE) String() string {
	if name, ok := labels_COPTER_MODE[e]; ok {
		return name
	}
	return strconv.Itoa(int(e))
}
