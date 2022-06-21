//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// High level message to control a gimbal manually. The angles or angular rates are unitless; the actual rates will depend on internal gimbal manager settings/configuration (e.g. set by parameters). This message is to be sent to the gimbal manager (e.g. from a ground station). Angles and rates can be set to NaN according to use case.
type MessageGimbalManagerSetManualControl struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// High level gimbal manager flags.
	Flags GIMBAL_MANAGER_FLAGS `mavenum:"uint32"`
	// Component ID of gimbal device to address (or 1-6 for non-MAVLink gimbal), 0 for all gimbal device components. Send command multiple times for more than one gimbal (but not all gimbals).
	GimbalDeviceId uint8
	// Pitch angle unitless (-1..1, positive: up, negative: down, NaN to be ignored).
	Pitch float32
	// Yaw angle unitless (-1..1, positive: to the right, negative: to the left, NaN to be ignored).
	Yaw float32
	// Pitch angular rate unitless (-1..1, positive: up, negative: down, NaN to be ignored).
	PitchRate float32
	// Yaw angular rate unitless (-1..1, positive: to the right, negative: to the left, NaN to be ignored).
	YawRate float32
}

// GetID implements the message.Message interface.
func (*MessageGimbalManagerSetManualControl) GetID() uint32 {
	return 288
}
