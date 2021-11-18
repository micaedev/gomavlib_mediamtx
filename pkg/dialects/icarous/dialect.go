//autogenerated:yes
//nolint:golint,misspell,govet,lll
// Package icarous contains the icarous dialect.
package icarous

import (
	"errors"

	"github.com/aler9/gomavlib/pkg/dialect"
	"github.com/aler9/gomavlib/pkg/msg"
)

// Dialect contains the dialect object that can be passed to the library.
var Dialect = dial

// dialect is not exposed directly such that it is not displayed in godoc.
var dial = &dialect.Dialect{0, []msg.Message{
	// icarous.xml
	&MessageIcarousHeartbeat{},
	&MessageIcarousKinematicBands{},
}}

type ICAROUS_FMS_STATE int

const (
	ICAROUS_FMS_STATE_IDLE     ICAROUS_FMS_STATE = 0
	ICAROUS_FMS_STATE_TAKEOFF  ICAROUS_FMS_STATE = 1
	ICAROUS_FMS_STATE_CLIMB    ICAROUS_FMS_STATE = 2
	ICAROUS_FMS_STATE_CRUISE   ICAROUS_FMS_STATE = 3
	ICAROUS_FMS_STATE_APPROACH ICAROUS_FMS_STATE = 4
	ICAROUS_FMS_STATE_LAND     ICAROUS_FMS_STATE = 5
)

var labels_ICAROUS_FMS_STATE = map[ICAROUS_FMS_STATE]string{
	ICAROUS_FMS_STATE_IDLE:     "ICAROUS_FMS_STATE_IDLE",
	ICAROUS_FMS_STATE_TAKEOFF:  "ICAROUS_FMS_STATE_TAKEOFF",
	ICAROUS_FMS_STATE_CLIMB:    "ICAROUS_FMS_STATE_CLIMB",
	ICAROUS_FMS_STATE_CRUISE:   "ICAROUS_FMS_STATE_CRUISE",
	ICAROUS_FMS_STATE_APPROACH: "ICAROUS_FMS_STATE_APPROACH",
	ICAROUS_FMS_STATE_LAND:     "ICAROUS_FMS_STATE_LAND",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ICAROUS_FMS_STATE) MarshalText() ([]byte, error) {
	if l, ok := labels_ICAROUS_FMS_STATE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_ICAROUS_FMS_STATE = map[string]ICAROUS_FMS_STATE{
	"ICAROUS_FMS_STATE_IDLE":     ICAROUS_FMS_STATE_IDLE,
	"ICAROUS_FMS_STATE_TAKEOFF":  ICAROUS_FMS_STATE_TAKEOFF,
	"ICAROUS_FMS_STATE_CLIMB":    ICAROUS_FMS_STATE_CLIMB,
	"ICAROUS_FMS_STATE_CRUISE":   ICAROUS_FMS_STATE_CRUISE,
	"ICAROUS_FMS_STATE_APPROACH": ICAROUS_FMS_STATE_APPROACH,
	"ICAROUS_FMS_STATE_LAND":     ICAROUS_FMS_STATE_LAND,
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ICAROUS_FMS_STATE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_ICAROUS_FMS_STATE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e ICAROUS_FMS_STATE) String() string {
	if l, ok := labels_ICAROUS_FMS_STATE[e]; ok {
		return l
	}
	return "invalid value"
}

type ICAROUS_TRACK_BAND_TYPES int

const (
	ICAROUS_TRACK_BAND_TYPE_NONE     ICAROUS_TRACK_BAND_TYPES = 0
	ICAROUS_TRACK_BAND_TYPE_NEAR     ICAROUS_TRACK_BAND_TYPES = 1
	ICAROUS_TRACK_BAND_TYPE_RECOVERY ICAROUS_TRACK_BAND_TYPES = 2
)

var labels_ICAROUS_TRACK_BAND_TYPES = map[ICAROUS_TRACK_BAND_TYPES]string{
	ICAROUS_TRACK_BAND_TYPE_NONE:     "ICAROUS_TRACK_BAND_TYPE_NONE",
	ICAROUS_TRACK_BAND_TYPE_NEAR:     "ICAROUS_TRACK_BAND_TYPE_NEAR",
	ICAROUS_TRACK_BAND_TYPE_RECOVERY: "ICAROUS_TRACK_BAND_TYPE_RECOVERY",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ICAROUS_TRACK_BAND_TYPES) MarshalText() ([]byte, error) {
	if l, ok := labels_ICAROUS_TRACK_BAND_TYPES[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_ICAROUS_TRACK_BAND_TYPES = map[string]ICAROUS_TRACK_BAND_TYPES{
	"ICAROUS_TRACK_BAND_TYPE_NONE":     ICAROUS_TRACK_BAND_TYPE_NONE,
	"ICAROUS_TRACK_BAND_TYPE_NEAR":     ICAROUS_TRACK_BAND_TYPE_NEAR,
	"ICAROUS_TRACK_BAND_TYPE_RECOVERY": ICAROUS_TRACK_BAND_TYPE_RECOVERY,
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ICAROUS_TRACK_BAND_TYPES) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_ICAROUS_TRACK_BAND_TYPES[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e ICAROUS_TRACK_BAND_TYPES) String() string {
	if l, ok := labels_ICAROUS_TRACK_BAND_TYPES[e]; ok {
		return l
	}
	return "invalid value"
}

// icarous.xml

// ICAROUS heartbeat
type MessageIcarousHeartbeat struct {
	// See the FMS_STATE enum.
	Status ICAROUS_FMS_STATE `mavenum:"uint8"`
}

// GetID implements the msg.Message interface.
func (*MessageIcarousHeartbeat) GetID() uint32 {
	return 42000
}

// Kinematic multi bands (track) output from Daidalus
type MessageIcarousKinematicBands struct {
	// Number of track bands
	Numbands int8 `mavname:"numBands"`
	// See the TRACK_BAND_TYPES enum.
	Type1 ICAROUS_TRACK_BAND_TYPES `mavenum:"uint8"`
	// min angle (degrees)
	Min1 float32
	// max angle (degrees)
	Max1 float32
	// See the TRACK_BAND_TYPES enum.
	Type2 ICAROUS_TRACK_BAND_TYPES `mavenum:"uint8"`
	// min angle (degrees)
	Min2 float32
	// max angle (degrees)
	Max2 float32
	// See the TRACK_BAND_TYPES enum.
	Type3 ICAROUS_TRACK_BAND_TYPES `mavenum:"uint8"`
	// min angle (degrees)
	Min3 float32
	// max angle (degrees)
	Max3 float32
	// See the TRACK_BAND_TYPES enum.
	Type4 ICAROUS_TRACK_BAND_TYPES `mavenum:"uint8"`
	// min angle (degrees)
	Min4 float32
	// max angle (degrees)
	Max4 float32
	// See the TRACK_BAND_TYPES enum.
	Type5 ICAROUS_TRACK_BAND_TYPES `mavenum:"uint8"`
	// min angle (degrees)
	Min5 float32
	// max angle (degrees)
	Max5 float32
}

// GetID implements the msg.Message interface.
func (*MessageIcarousKinematicBands) GetID() uint32 {
	return 42001
}
