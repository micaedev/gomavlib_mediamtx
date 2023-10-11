//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package paparazzi

import (
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/common"
)

// Flags for gimbal device (lower level) operation.
type GIMBAL_DEVICE_FLAGS = common.GIMBAL_DEVICE_FLAGS

const (
	// Set to retracted safe position (no stabilization), takes precedence over all other flags.
	GIMBAL_DEVICE_FLAGS_RETRACT GIMBAL_DEVICE_FLAGS = common.GIMBAL_DEVICE_FLAGS_RETRACT
	// Set to neutral/default position, taking precedence over all other flags except RETRACT. Neutral is commonly forward-facing and horizontal (roll=pitch=yaw=0) but may be any orientation.
	GIMBAL_DEVICE_FLAGS_NEUTRAL GIMBAL_DEVICE_FLAGS = common.GIMBAL_DEVICE_FLAGS_NEUTRAL
	// Lock roll angle to absolute angle relative to horizon (not relative to vehicle). This is generally the default with a stabilizing gimbal.
	GIMBAL_DEVICE_FLAGS_ROLL_LOCK GIMBAL_DEVICE_FLAGS = common.GIMBAL_DEVICE_FLAGS_ROLL_LOCK
	// Lock pitch angle to absolute angle relative to horizon (not relative to vehicle). This is generally the default with a stabilizing gimbal.
	GIMBAL_DEVICE_FLAGS_PITCH_LOCK GIMBAL_DEVICE_FLAGS = common.GIMBAL_DEVICE_FLAGS_PITCH_LOCK
	// Lock yaw angle to absolute angle relative to North (not relative to vehicle). If this flag is set, the yaw angle and z component of angular velocity are relative to North (earth frame, x-axis pointing North), else they are relative to the vehicle heading (vehicle frame, earth frame rotated so that the x-axis is pointing forward).
	GIMBAL_DEVICE_FLAGS_YAW_LOCK GIMBAL_DEVICE_FLAGS = common.GIMBAL_DEVICE_FLAGS_YAW_LOCK
	// Yaw angle and z component of angular velocity are relative to the vehicle heading (vehicle frame, earth frame rotated such that the x-axis is pointing forward).
	GIMBAL_DEVICE_FLAGS_YAW_IN_VEHICLE_FRAME GIMBAL_DEVICE_FLAGS = common.GIMBAL_DEVICE_FLAGS_YAW_IN_VEHICLE_FRAME
	// Yaw angle and z component of angular velocity are relative to North (earth frame, x-axis is pointing North).
	GIMBAL_DEVICE_FLAGS_YAW_IN_EARTH_FRAME GIMBAL_DEVICE_FLAGS = common.GIMBAL_DEVICE_FLAGS_YAW_IN_EARTH_FRAME
	// Gimbal device can accept yaw angle inputs relative to North (earth frame). This flag is only for reporting (attempts to set this flag are ignored).
	GIMBAL_DEVICE_FLAGS_ACCEPTS_YAW_IN_EARTH_FRAME GIMBAL_DEVICE_FLAGS = common.GIMBAL_DEVICE_FLAGS_ACCEPTS_YAW_IN_EARTH_FRAME
	// The gimbal orientation is set exclusively by the RC signals feed to the gimbal's radio control inputs. MAVLink messages for setting the gimbal orientation (GIMBAL_DEVICE_SET_ATTITUDE) are ignored.
	GIMBAL_DEVICE_FLAGS_RC_EXCLUSIVE GIMBAL_DEVICE_FLAGS = common.GIMBAL_DEVICE_FLAGS_RC_EXCLUSIVE
	// The gimbal orientation is determined by combining/mixing the RC signals feed to the gimbal's radio control inputs and the MAVLink messages for setting the gimbal orientation (GIMBAL_DEVICE_SET_ATTITUDE). How these two controls are combined or mixed is not defined by the protocol but is up to the implementation.
	GIMBAL_DEVICE_FLAGS_RC_MIXED GIMBAL_DEVICE_FLAGS = common.GIMBAL_DEVICE_FLAGS_RC_MIXED
)
