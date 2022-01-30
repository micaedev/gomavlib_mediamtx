//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package matrixpilot

import (
	"errors"
)

// Supported component metadata types. These are used in the "general" metadata file returned by COMPONENT_INFORMATION to provide information about supported metadata types. The types are not used directly in MAVLink messages.
type COMP_METADATA_TYPE uint32

const (
	// General information about the component. General metadata includes information about other COMP_METADATA_TYPEs supported by the component. This type must be supported and must be downloadable from vehicle.
	COMP_METADATA_TYPE_GENERAL COMP_METADATA_TYPE = 0
	// Parameter meta data.
	COMP_METADATA_TYPE_PARAMETER COMP_METADATA_TYPE = 1
	// Meta data that specifies which commands and command parameters the vehicle supports. (WIP)
	COMP_METADATA_TYPE_COMMANDS COMP_METADATA_TYPE = 2
	// Meta data that specifies external non-MAVLink peripherals.
	COMP_METADATA_TYPE_PERIPHERALS COMP_METADATA_TYPE = 3
	// Meta data for the events interface.
	COMP_METADATA_TYPE_EVENTS COMP_METADATA_TYPE = 4
	// Meta data for actuator configuration (motors, servos and vehicle geometry) and testing.
	COMP_METADATA_TYPE_ACTUATORS COMP_METADATA_TYPE = 5
)

var labels_COMP_METADATA_TYPE = map[COMP_METADATA_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e COMP_METADATA_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_COMP_METADATA_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_COMP_METADATA_TYPE = map[string]COMP_METADATA_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *COMP_METADATA_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_COMP_METADATA_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e COMP_METADATA_TYPE) String() string {
	if l, ok := labels_COMP_METADATA_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
