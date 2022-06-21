//autogenerated:yes
//nolint:revive,misspell,govet,lll
package paparazzi

// High level message to control a gimbal's pitch and yaw angles. This message is to be sent to the gimbal manager (e.g. from a ground station). Angles and rates can be set to NaN according to use case.
type MessageGimbalManagerSetPitchyaw struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// High level gimbal manager flags to use.
	Flags GIMBAL_MANAGER_FLAGS `mavenum:"uint32"`
	// Component ID of gimbal device to address (or 1-6 for non-MAVLink gimbal), 0 for all gimbal device components. Send command multiple times for more than one gimbal (but not all gimbals).
	GimbalDeviceId uint8
	// Pitch angle (positive: up, negative: down, NaN to be ignored).
	Pitch float32
	// Yaw angle (positive: to the right, negative: to the left, NaN to be ignored).
	Yaw float32
	// Pitch angular rate (positive: up, negative: down, NaN to be ignored).
	PitchRate float32
	// Yaw angular rate (positive: to the right, negative: to the left, NaN to be ignored).
	YawRate float32
}

// GetID implements the message.Message interface.
func (*MessageGimbalManagerSetPitchyaw) GetID() uint32 {
	return 287
}
