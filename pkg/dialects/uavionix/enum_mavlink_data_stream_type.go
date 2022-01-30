//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package uavionix

import (
	"errors"
)

type MAVLINK_DATA_STREAM_TYPE uint32

const (
	MAVLINK_DATA_STREAM_IMG_JPEG   MAVLINK_DATA_STREAM_TYPE = 0
	MAVLINK_DATA_STREAM_IMG_BMP    MAVLINK_DATA_STREAM_TYPE = 1
	MAVLINK_DATA_STREAM_IMG_RAW8U  MAVLINK_DATA_STREAM_TYPE = 2
	MAVLINK_DATA_STREAM_IMG_RAW32U MAVLINK_DATA_STREAM_TYPE = 3
	MAVLINK_DATA_STREAM_IMG_PGM    MAVLINK_DATA_STREAM_TYPE = 4
	MAVLINK_DATA_STREAM_IMG_PNG    MAVLINK_DATA_STREAM_TYPE = 5
)

var labels_MAVLINK_DATA_STREAM_TYPE = map[MAVLINK_DATA_STREAM_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAVLINK_DATA_STREAM_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAVLINK_DATA_STREAM_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAVLINK_DATA_STREAM_TYPE = map[string]MAVLINK_DATA_STREAM_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAVLINK_DATA_STREAM_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAVLINK_DATA_STREAM_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAVLINK_DATA_STREAM_TYPE) String() string {
	if l, ok := labels_MAVLINK_DATA_STREAM_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
