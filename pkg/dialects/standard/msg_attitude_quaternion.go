//autogenerated:yes
//nolint:revive,misspell,govet,lll
package standard

// The attitude in the aeronautical frame (right-handed, Z-down, X-front, Y-right), expressed as quaternion. Quaternion order is w, x, y, z and a zero rotation would be expressed as (1 0 0 0).
type MessageAttitudeQuaternion struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Quaternion component 1, w (1 in null-rotation)
	Q1 float32
	// Quaternion component 2, x (0 in null-rotation)
	Q2 float32
	// Quaternion component 3, y (0 in null-rotation)
	Q3 float32
	// Quaternion component 4, z (0 in null-rotation)
	Q4 float32
	// Roll angular speed
	Rollspeed float32
	// Pitch angular speed
	Pitchspeed float32
	// Yaw angular speed
	Yawspeed float32
	// Rotation offset by which the attitude quaternion and angular speed vector should be rotated for user display (quaternion with [w, x, y, z] order, zero-rotation is [1, 0, 0, 0], send [0, 0, 0, 0] if field not supported). This field is intended for systems in which the reference attitude may change during flight. For example, tailsitters VTOLs rotate their reference attitude by 90 degrees between hover mode and fixed wing mode, thus repr_offset_q is equal to [1, 0, 0, 0] in hover mode and equal to [0.7071, 0, 0.7071, 0] in fixed wing mode.
	ReprOffsetQ [4]float32 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageAttitudeQuaternion) GetID() uint32 {
	return 31
}
