//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package development

import (
	"errors"
)

// Possible transport layers to set and get parameters via mavlink during a parameter transaction.
type PARAM_TRANSACTION_TRANSPORT uint32

const (
	// Transaction over param transport.
	PARAM_TRANSACTION_TRANSPORT_PARAM PARAM_TRANSACTION_TRANSPORT = 0
	// Transaction over param_ext transport.
	PARAM_TRANSACTION_TRANSPORT_PARAM_EXT PARAM_TRANSACTION_TRANSPORT = 1
)

var labels_PARAM_TRANSACTION_TRANSPORT = map[PARAM_TRANSACTION_TRANSPORT]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e PARAM_TRANSACTION_TRANSPORT) MarshalText() ([]byte, error) {
	if l, ok := labels_PARAM_TRANSACTION_TRANSPORT[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_PARAM_TRANSACTION_TRANSPORT = map[string]PARAM_TRANSACTION_TRANSPORT{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *PARAM_TRANSACTION_TRANSPORT) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_PARAM_TRANSACTION_TRANSPORT[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e PARAM_TRANSACTION_TRANSPORT) String() string {
	if l, ok := labels_PARAM_TRANSACTION_TRANSPORT[e]; ok {
		return l
	}
	return "invalid value"
}
