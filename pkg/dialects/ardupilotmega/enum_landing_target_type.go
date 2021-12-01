//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

// Type of landing target
type LANDING_TARGET_TYPE int

const (
	// Landing target signaled by light beacon (ex: IR-LOCK)
	LANDING_TARGET_TYPE_LIGHT_BEACON LANDING_TARGET_TYPE = 0
	// Landing target signaled by radio beacon (ex: ILS, NDB)
	LANDING_TARGET_TYPE_RADIO_BEACON LANDING_TARGET_TYPE = 1
	// Landing target represented by a fiducial marker (ex: ARTag)
	LANDING_TARGET_TYPE_VISION_FIDUCIAL LANDING_TARGET_TYPE = 2
	// Landing target represented by a pre-defined visual shape/feature (ex: X-marker, H-marker, square)
	LANDING_TARGET_TYPE_VISION_OTHER LANDING_TARGET_TYPE = 3
)

var labels_LANDING_TARGET_TYPE = map[LANDING_TARGET_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e LANDING_TARGET_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_LANDING_TARGET_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_LANDING_TARGET_TYPE = map[string]LANDING_TARGET_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *LANDING_TARGET_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_LANDING_TARGET_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e LANDING_TARGET_TYPE) String() string {
	if l, ok := labels_LANDING_TARGET_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
