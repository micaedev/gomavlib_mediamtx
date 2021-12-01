//autogenerated:yes
//nolint:golint,misspell,govet,lll
package common

// Low level message to control a gimbal device's attitude. This message is to be sent from the gimbal manager to the gimbal device component. Angles and rates can be set to NaN according to use case.
type MessageGimbalDeviceSetAttitude struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Low level gimbal flags.
	Flags GIMBAL_DEVICE_FLAGS `mavenum:"uint16"`
	// Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation, the frame is depends on whether the flag GIMBAL_DEVICE_FLAGS_YAW_LOCK is set, set all fields to NaN if only angular velocity should be used)
	Q [4]float32
	// X component of angular velocity, positive is rolling to the right, NaN to be ignored.
	AngularVelocityX float32
	// Y component of angular velocity, positive is pitching up, NaN to be ignored.
	AngularVelocityY float32
	// Z component of angular velocity, positive is yawing to the right, NaN to be ignored.
	AngularVelocityZ float32
}

// GetID implements the msg.Message interface.
func (*MessageGimbalDeviceSetAttitude) GetID() uint32 {
	return 284
}
