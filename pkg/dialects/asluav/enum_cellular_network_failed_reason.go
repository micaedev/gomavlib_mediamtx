//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package asluav

import (
	"errors"
)

// These flags are used to diagnose the failure state of CELLULAR_STATUS
type CELLULAR_NETWORK_FAILED_REASON int

const (
	// No error
	CELLULAR_NETWORK_FAILED_REASON_NONE CELLULAR_NETWORK_FAILED_REASON = 0
	// Error state is unknown
	CELLULAR_NETWORK_FAILED_REASON_UNKNOWN CELLULAR_NETWORK_FAILED_REASON = 1
	// SIM is required for the modem but missing
	CELLULAR_NETWORK_FAILED_REASON_SIM_MISSING CELLULAR_NETWORK_FAILED_REASON = 2
	// SIM is available, but not usuable for connection
	CELLULAR_NETWORK_FAILED_REASON_SIM_ERROR CELLULAR_NETWORK_FAILED_REASON = 3
)

var labels_CELLULAR_NETWORK_FAILED_REASON = map[CELLULAR_NETWORK_FAILED_REASON]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e CELLULAR_NETWORK_FAILED_REASON) MarshalText() ([]byte, error) {
	if l, ok := labels_CELLULAR_NETWORK_FAILED_REASON[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_CELLULAR_NETWORK_FAILED_REASON = map[string]CELLULAR_NETWORK_FAILED_REASON{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *CELLULAR_NETWORK_FAILED_REASON) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_CELLULAR_NETWORK_FAILED_REASON[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e CELLULAR_NETWORK_FAILED_REASON) String() string {
	if l, ok := labels_CELLULAR_NETWORK_FAILED_REASON[e]; ok {
		return l
	}
	return "invalid value"
}
