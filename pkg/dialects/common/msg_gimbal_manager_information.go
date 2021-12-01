//autogenerated:yes
//nolint:golint,misspell,govet,lll
package common

// Information about a high level gimbal manager. This message should be requested by a ground station using MAV_CMD_REQUEST_MESSAGE.
type MessageGimbalManagerInformation struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Bitmap of gimbal capability flags.
	CapFlags GIMBAL_MANAGER_CAP_FLAGS `mavenum:"uint32"`
	// Gimbal device ID that this gimbal manager is responsible for.
	GimbalDeviceId uint8
	// Minimum hardware roll angle (positive: rolling to the right, negative: rolling to the left)
	RollMin float32
	// Maximum hardware roll angle (positive: rolling to the right, negative: rolling to the left)
	RollMax float32
	// Minimum pitch angle (positive: up, negative: down)
	PitchMin float32
	// Maximum pitch angle (positive: up, negative: down)
	PitchMax float32
	// Minimum yaw angle (positive: to the right, negative: to the left)
	YawMin float32
	// Maximum yaw angle (positive: to the right, negative: to the left)
	YawMax float32
}

// GetID implements the msg.Message interface.
func (*MessageGimbalManagerInformation) GetID() uint32 {
	return 280
}
