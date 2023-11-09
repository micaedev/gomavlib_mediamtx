//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strconv"
)

type MAV_ODID_CATEGORY_EU uint32

const (
	// The category for the UA, according to the EU specification, is undeclared.
	MAV_ODID_CATEGORY_EU_UNDECLARED MAV_ODID_CATEGORY_EU = 0
	// The category for the UA, according to the EU specification, is the Open category.
	MAV_ODID_CATEGORY_EU_OPEN MAV_ODID_CATEGORY_EU = 1
	// The category for the UA, according to the EU specification, is the Specific category.
	MAV_ODID_CATEGORY_EU_SPECIFIC MAV_ODID_CATEGORY_EU = 2
	// The category for the UA, according to the EU specification, is the Certified category.
	MAV_ODID_CATEGORY_EU_CERTIFIED MAV_ODID_CATEGORY_EU = 3
)

var labels_MAV_ODID_CATEGORY_EU = map[MAV_ODID_CATEGORY_EU]string{
	MAV_ODID_CATEGORY_EU_UNDECLARED: "MAV_ODID_CATEGORY_EU_UNDECLARED",
	MAV_ODID_CATEGORY_EU_OPEN:       "MAV_ODID_CATEGORY_EU_OPEN",
	MAV_ODID_CATEGORY_EU_SPECIFIC:   "MAV_ODID_CATEGORY_EU_SPECIFIC",
	MAV_ODID_CATEGORY_EU_CERTIFIED:  "MAV_ODID_CATEGORY_EU_CERTIFIED",
}

var values_MAV_ODID_CATEGORY_EU = map[string]MAV_ODID_CATEGORY_EU{
	"MAV_ODID_CATEGORY_EU_UNDECLARED": MAV_ODID_CATEGORY_EU_UNDECLARED,
	"MAV_ODID_CATEGORY_EU_OPEN":       MAV_ODID_CATEGORY_EU_OPEN,
	"MAV_ODID_CATEGORY_EU_SPECIFIC":   MAV_ODID_CATEGORY_EU_SPECIFIC,
	"MAV_ODID_CATEGORY_EU_CERTIFIED":  MAV_ODID_CATEGORY_EU_CERTIFIED,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_CATEGORY_EU) MarshalText() ([]byte, error) {
	if name, ok := labels_MAV_ODID_CATEGORY_EU[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_CATEGORY_EU) UnmarshalText(text []byte) error {
	if value, ok := values_MAV_ODID_CATEGORY_EU[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = MAV_ODID_CATEGORY_EU(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_CATEGORY_EU) String() string {
	if name, ok := labels_MAV_ODID_CATEGORY_EU[e]; ok {
		return name
	}
	return strconv.Itoa(int(e))
}
