//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ualberta

import (
	"errors"
)

// Tune formats (used for vehicle buzzer/tone generation).
type TUNE_FORMAT uint32

const (
	// Format is QBasic 1.1 Play: https://www.qbasic.net/en/reference/qb11/Statement/PLAY-006.htm.
	TUNE_FORMAT_QBASIC1_1 TUNE_FORMAT = 1
	// Format is Modern Music Markup Language (MML): https://en.wikipedia.org/wiki/Music_Macro_Language#Modern_MML.
	TUNE_FORMAT_MML_MODERN TUNE_FORMAT = 2
)

var labels_TUNE_FORMAT = map[TUNE_FORMAT]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e TUNE_FORMAT) MarshalText() ([]byte, error) {
	if l, ok := labels_TUNE_FORMAT[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_TUNE_FORMAT = map[string]TUNE_FORMAT{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *TUNE_FORMAT) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_TUNE_FORMAT[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e TUNE_FORMAT) String() string {
	if l, ok := labels_TUNE_FORMAT[e]; ok {
		return l
	}
	return "invalid value"
}
