//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package asluav

import (
	"errors"
)

// Winch actions.
type WINCH_ACTIONS uint32

const (
	// Allow motor to freewheel.
	WINCH_RELAXED WINCH_ACTIONS = 0
	// Wind or unwind specified length of line, optionally using specified rate.
	WINCH_RELATIVE_LENGTH_CONTROL WINCH_ACTIONS = 1
	// Wind or unwind line at specified rate.
	WINCH_RATE_CONTROL WINCH_ACTIONS = 2
	// Perform the locking sequence to relieve motor while in the fully retracted position. Only action and instance command parameters are used, others are ignored.
	WINCH_LOCK WINCH_ACTIONS = 3
	// Sequence of drop, slow down, touch down, reel up, lock. Only action and instance command parameters are used, others are ignored.
	WINCH_DELIVER WINCH_ACTIONS = 4
	// Engage motor and hold current position. Only action and instance command parameters are used, others are ignored.
	WINCH_HOLD WINCH_ACTIONS = 5
	// Return the reel to the fully retracted position. Only action and instance command parameters are used, others are ignored.
	WINCH_RETRACT WINCH_ACTIONS = 6
	// Load the reel with line. The winch will calculate the total loaded length and stop when the tension exceeds a threshold. Only action and instance command parameters are used, others are ignored.
	WINCH_LOAD_LINE WINCH_ACTIONS = 7
	// Spool out the entire length of the line. Only action and instance command parameters are used, others are ignored.
	WINCH_ABANDON_LINE WINCH_ACTIONS = 8
)

var labels_WINCH_ACTIONS = map[WINCH_ACTIONS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e WINCH_ACTIONS) MarshalText() ([]byte, error) {
	if l, ok := labels_WINCH_ACTIONS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_WINCH_ACTIONS = map[string]WINCH_ACTIONS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *WINCH_ACTIONS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_WINCH_ACTIONS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e WINCH_ACTIONS) String() string {
	if l, ok := labels_WINCH_ACTIONS[e]; ok {
		return l
	}
	return "invalid value"
}
