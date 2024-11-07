//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package cubepilot

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
)

// RC type. Used in MAV_CMD_START_RX_PAIR.
type RC_TYPE = common.RC_TYPE

const (
	// Spektrum
	RC_TYPE_SPEKTRUM RC_TYPE = common.RC_TYPE_SPEKTRUM
	// CRSF
	RC_TYPE_CRSF RC_TYPE = common.RC_TYPE_CRSF
)
