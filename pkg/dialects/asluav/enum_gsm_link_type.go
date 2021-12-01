//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package asluav

import (
	"errors"
)

type GSM_LINK_TYPE int

const (
	// no service
	GSM_LINK_TYPE_NONE GSM_LINK_TYPE = 0
	// link type unknown
	GSM_LINK_TYPE_UNKNOWN GSM_LINK_TYPE = 1
	// 2G (GSM/GRPS/EDGE) link
	GSM_LINK_TYPE_2G GSM_LINK_TYPE = 2
	// 3G link (WCDMA/HSDPA/HSPA)
	GSM_LINK_TYPE_3G GSM_LINK_TYPE = 3
	// 4G link (LTE)
	GSM_LINK_TYPE_4G GSM_LINK_TYPE = 4
)

var labels_GSM_LINK_TYPE = map[GSM_LINK_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GSM_LINK_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_GSM_LINK_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GSM_LINK_TYPE = map[string]GSM_LINK_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GSM_LINK_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GSM_LINK_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GSM_LINK_TYPE) String() string {
	if l, ok := labels_GSM_LINK_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
