//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

// Actions being taken to mitigate/prevent fence breach
type FENCE_MITIGATE uint32

const (
	// Unknown
	FENCE_MITIGATE_UNKNOWN FENCE_MITIGATE = 0
	// No actions being taken
	FENCE_MITIGATE_NONE FENCE_MITIGATE = 1
	// Velocity limiting active to prevent breach
	FENCE_MITIGATE_VEL_LIMIT FENCE_MITIGATE = 2
)

var labels_FENCE_MITIGATE = map[FENCE_MITIGATE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e FENCE_MITIGATE) MarshalText() ([]byte, error) {
	if l, ok := labels_FENCE_MITIGATE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_FENCE_MITIGATE = map[string]FENCE_MITIGATE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *FENCE_MITIGATE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_FENCE_MITIGATE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e FENCE_MITIGATE) String() string {
	if l, ok := labels_FENCE_MITIGATE[e]; ok {
		return l
	}
	return "invalid value"
}
