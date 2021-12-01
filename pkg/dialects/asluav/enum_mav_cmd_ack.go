//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package asluav

import (
	"errors"
)

// ACK / NACK / ERROR values as a result of MAV_CMDs and for mission item transmission.
type MAV_CMD_ACK int

const (
	// Command / mission item is ok.
	MAV_CMD_ACK_OK MAV_CMD_ACK = 0
	// Generic error message if none of the other reasons fails or if no detailed error reporting is implemented.
	MAV_CMD_ACK_ERR_FAIL MAV_CMD_ACK = 1
	// The system is refusing to accept this command from this source / communication partner.
	MAV_CMD_ACK_ERR_ACCESS_DENIED MAV_CMD_ACK = 2
	// Command or mission item is not supported, other commands would be accepted.
	MAV_CMD_ACK_ERR_NOT_SUPPORTED MAV_CMD_ACK = 3
	// The coordinate frame of this command / mission item is not supported.
	MAV_CMD_ACK_ERR_COORDINATE_FRAME_NOT_SUPPORTED MAV_CMD_ACK = 4
	// The coordinate frame of this command is ok, but he coordinate values exceed the safety limits of this system. This is a generic error, please use the more specific error messages below if possible.
	MAV_CMD_ACK_ERR_COORDINATES_OUT_OF_RANGE MAV_CMD_ACK = 5
	// The X or latitude value is out of range.
	MAV_CMD_ACK_ERR_X_LAT_OUT_OF_RANGE MAV_CMD_ACK = 6
	// The Y or longitude value is out of range.
	MAV_CMD_ACK_ERR_Y_LON_OUT_OF_RANGE MAV_CMD_ACK = 7
	// The Z or altitude value is out of range.
	MAV_CMD_ACK_ERR_Z_ALT_OUT_OF_RANGE MAV_CMD_ACK = 8
)

var labels_MAV_CMD_ACK = map[MAV_CMD_ACK]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_CMD_ACK) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_CMD_ACK[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_CMD_ACK = map[string]MAV_CMD_ACK{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_CMD_ACK) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_CMD_ACK[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_CMD_ACK) String() string {
	if l, ok := labels_MAV_CMD_ACK[e]; ok {
		return l
	}
	return "invalid value"
}
