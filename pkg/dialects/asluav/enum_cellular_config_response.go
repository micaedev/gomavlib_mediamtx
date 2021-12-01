//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package asluav

import (
	"errors"
)

// Possible responses from a CELLULAR_CONFIG message.
type CELLULAR_CONFIG_RESPONSE int

const (
	// Changes accepted.
	CELLULAR_CONFIG_RESPONSE_ACCEPTED CELLULAR_CONFIG_RESPONSE = 0
	// Invalid APN.
	CELLULAR_CONFIG_RESPONSE_APN_ERROR CELLULAR_CONFIG_RESPONSE = 1
	// Invalid PIN.
	CELLULAR_CONFIG_RESPONSE_PIN_ERROR CELLULAR_CONFIG_RESPONSE = 2
	// Changes rejected.
	CELLULAR_CONFIG_RESPONSE_REJECTED CELLULAR_CONFIG_RESPONSE = 3
	// PUK is required to unblock SIM card.
	CELLULAR_CONFIG_BLOCKED_PUK_REQUIRED CELLULAR_CONFIG_RESPONSE = 4
)

var labels_CELLULAR_CONFIG_RESPONSE = map[CELLULAR_CONFIG_RESPONSE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e CELLULAR_CONFIG_RESPONSE) MarshalText() ([]byte, error) {
	if l, ok := labels_CELLULAR_CONFIG_RESPONSE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_CELLULAR_CONFIG_RESPONSE = map[string]CELLULAR_CONFIG_RESPONSE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *CELLULAR_CONFIG_RESPONSE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_CELLULAR_CONFIG_RESPONSE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e CELLULAR_CONFIG_RESPONSE) String() string {
	if l, ok := labels_CELLULAR_CONFIG_RESPONSE[e]; ok {
		return l
	}
	return "invalid value"
}
