//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package ualberta

import (
	"errors"
)

// SERIAL_CONTROL flags (bitmask)
type SERIAL_CONTROL_FLAG int

const (
	// Set if this is a reply
	SERIAL_CONTROL_FLAG_REPLY SERIAL_CONTROL_FLAG = 1
	// Set if the sender wants the receiver to send a response as another SERIAL_CONTROL message
	SERIAL_CONTROL_FLAG_RESPOND SERIAL_CONTROL_FLAG = 2
	// Set if access to the serial port should be removed from whatever driver is currently using it, giving exclusive access to the SERIAL_CONTROL protocol. The port can be handed back by sending a request without this flag set
	SERIAL_CONTROL_FLAG_EXCLUSIVE SERIAL_CONTROL_FLAG = 4
	// Block on writes to the serial port
	SERIAL_CONTROL_FLAG_BLOCKING SERIAL_CONTROL_FLAG = 8
	// Send multiple replies until port is drained
	SERIAL_CONTROL_FLAG_MULTI SERIAL_CONTROL_FLAG = 16
)

var labels_SERIAL_CONTROL_FLAG = map[SERIAL_CONTROL_FLAG]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e SERIAL_CONTROL_FLAG) MarshalText() ([]byte, error) {
	if l, ok := labels_SERIAL_CONTROL_FLAG[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_SERIAL_CONTROL_FLAG = map[string]SERIAL_CONTROL_FLAG{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *SERIAL_CONTROL_FLAG) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_SERIAL_CONTROL_FLAG[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e SERIAL_CONTROL_FLAG) String() string {
	if l, ok := labels_SERIAL_CONTROL_FLAG[e]; ok {
		return l
	}
	return "invalid value"
}
