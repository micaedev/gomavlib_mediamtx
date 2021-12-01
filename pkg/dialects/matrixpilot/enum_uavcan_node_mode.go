//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package matrixpilot

import (
	"errors"
)

// Generalized UAVCAN node mode
type UAVCAN_NODE_MODE int

const (
	// The node is performing its primary functions.
	UAVCAN_NODE_MODE_OPERATIONAL UAVCAN_NODE_MODE = 0
	// The node is initializing; this mode is entered immediately after startup.
	UAVCAN_NODE_MODE_INITIALIZATION UAVCAN_NODE_MODE = 1
	// The node is under maintenance.
	UAVCAN_NODE_MODE_MAINTENANCE UAVCAN_NODE_MODE = 2
	// The node is in the process of updating its software.
	UAVCAN_NODE_MODE_SOFTWARE_UPDATE UAVCAN_NODE_MODE = 3
	// The node is no longer available online.
	UAVCAN_NODE_MODE_OFFLINE UAVCAN_NODE_MODE = 7
)

var labels_UAVCAN_NODE_MODE = map[UAVCAN_NODE_MODE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UAVCAN_NODE_MODE) MarshalText() ([]byte, error) {
	if l, ok := labels_UAVCAN_NODE_MODE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_UAVCAN_NODE_MODE = map[string]UAVCAN_NODE_MODE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UAVCAN_NODE_MODE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_UAVCAN_NODE_MODE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e UAVCAN_NODE_MODE) String() string {
	if l, ok := labels_UAVCAN_NODE_MODE[e]; ok {
		return l
	}
	return "invalid value"
}
