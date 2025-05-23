//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"fmt"
	"strconv"
)

type GOPRO_PROTUNE_GAIN uint64

const (
	// ISO 400.
	GOPRO_PROTUNE_GAIN_400 GOPRO_PROTUNE_GAIN = 0
	// ISO 800 (Only Hero 4).
	GOPRO_PROTUNE_GAIN_800 GOPRO_PROTUNE_GAIN = 1
	// ISO 1600.
	GOPRO_PROTUNE_GAIN_1600 GOPRO_PROTUNE_GAIN = 2
	// ISO 3200 (Only Hero 4).
	GOPRO_PROTUNE_GAIN_3200 GOPRO_PROTUNE_GAIN = 3
	// ISO 6400.
	GOPRO_PROTUNE_GAIN_6400 GOPRO_PROTUNE_GAIN = 4
)

var value_to_label_GOPRO_PROTUNE_GAIN = map[GOPRO_PROTUNE_GAIN]string{
	GOPRO_PROTUNE_GAIN_400:  "GOPRO_PROTUNE_GAIN_400",
	GOPRO_PROTUNE_GAIN_800:  "GOPRO_PROTUNE_GAIN_800",
	GOPRO_PROTUNE_GAIN_1600: "GOPRO_PROTUNE_GAIN_1600",
	GOPRO_PROTUNE_GAIN_3200: "GOPRO_PROTUNE_GAIN_3200",
	GOPRO_PROTUNE_GAIN_6400: "GOPRO_PROTUNE_GAIN_6400",
}

var label_to_value_GOPRO_PROTUNE_GAIN = map[string]GOPRO_PROTUNE_GAIN{
	"GOPRO_PROTUNE_GAIN_400":  GOPRO_PROTUNE_GAIN_400,
	"GOPRO_PROTUNE_GAIN_800":  GOPRO_PROTUNE_GAIN_800,
	"GOPRO_PROTUNE_GAIN_1600": GOPRO_PROTUNE_GAIN_1600,
	"GOPRO_PROTUNE_GAIN_3200": GOPRO_PROTUNE_GAIN_3200,
	"GOPRO_PROTUNE_GAIN_6400": GOPRO_PROTUNE_GAIN_6400,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GOPRO_PROTUNE_GAIN) MarshalText() ([]byte, error) {
	if name, ok := value_to_label_GOPRO_PROTUNE_GAIN[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GOPRO_PROTUNE_GAIN) UnmarshalText(text []byte) error {
	if value, ok := label_to_value_GOPRO_PROTUNE_GAIN[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = GOPRO_PROTUNE_GAIN(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e GOPRO_PROTUNE_GAIN) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
