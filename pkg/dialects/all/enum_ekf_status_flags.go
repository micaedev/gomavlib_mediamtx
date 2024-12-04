//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/ardupilotmega"
)

// Flags in EKF_STATUS message.
type EKF_STATUS_FLAGS = ardupilotmega.EKF_STATUS_FLAGS

const (
	// Set if EKF's attitude estimate is good.
	EKF_ATTITUDE EKF_STATUS_FLAGS = ardupilotmega.EKF_ATTITUDE
	// Set if EKF's horizontal velocity estimate is good.
	EKF_VELOCITY_HORIZ EKF_STATUS_FLAGS = ardupilotmega.EKF_VELOCITY_HORIZ
	// Set if EKF's vertical velocity estimate is good.
	EKF_VELOCITY_VERT EKF_STATUS_FLAGS = ardupilotmega.EKF_VELOCITY_VERT
	// Set if EKF's horizontal position (relative) estimate is good.
	EKF_POS_HORIZ_REL EKF_STATUS_FLAGS = ardupilotmega.EKF_POS_HORIZ_REL
	// Set if EKF's horizontal position (absolute) estimate is good.
	EKF_POS_HORIZ_ABS EKF_STATUS_FLAGS = ardupilotmega.EKF_POS_HORIZ_ABS
	// Set if EKF's vertical position (absolute) estimate is good.
	EKF_POS_VERT_ABS EKF_STATUS_FLAGS = ardupilotmega.EKF_POS_VERT_ABS
	// Set if EKF's vertical position (above ground) estimate is good.
	EKF_POS_VERT_AGL EKF_STATUS_FLAGS = ardupilotmega.EKF_POS_VERT_AGL
	// EKF is in constant position mode and does not know it's absolute or relative position.
	EKF_CONST_POS_MODE EKF_STATUS_FLAGS = ardupilotmega.EKF_CONST_POS_MODE
	// Set if EKF's predicted horizontal position (relative) estimate is good.
	EKF_PRED_POS_HORIZ_REL EKF_STATUS_FLAGS = ardupilotmega.EKF_PRED_POS_HORIZ_REL
	// Set if EKF's predicted horizontal position (absolute) estimate is good.
	EKF_PRED_POS_HORIZ_ABS EKF_STATUS_FLAGS = ardupilotmega.EKF_PRED_POS_HORIZ_ABS
	// Set if EKF believes the GPS input data is faulty.
	EKF_GPS_GLITCHING EKF_STATUS_FLAGS = ardupilotmega.EKF_GPS_GLITCHING
	// Set if EKF has never been healthy.
	EKF_UNINITIALIZED EKF_STATUS_FLAGS = ardupilotmega.EKF_UNINITIALIZED
)
