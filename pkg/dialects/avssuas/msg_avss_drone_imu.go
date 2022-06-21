//autogenerated:yes
//nolint:revive,misspell,govet,lll
package avssuas

// Drone IMU data. Quaternion order is w, x, y, z and a zero rotation would be expressed as (1 0 0 0).
type MessageAvssDroneImu struct {
	// Timestamp (time since FC boot).
	TimeBootMs uint32
	// Quaternion component 1, w (1 in null-rotation)
	Q1 float32
	// Quaternion component 2, x (0 in null-rotation)
	Q2 float32
	// Quaternion component 3, y (0 in null-rotation)
	Q3 float32
	// Quaternion component 4, z (0 in null-rotation)
	Q4 float32
	// X acceleration
	Xacc float32
	// Y acceleration
	Yacc float32
	// Z acceleration
	Zacc float32
	// Angular speed around X axis
	Xgyro float32
	// Angular speed around Y axis
	Ygyro float32
	// Angular speed around Z axis
	Zgyro float32
}

// GetID implements the message.Message interface.
func (*MessageAvssDroneImu) GetID() uint32 {
	return 60052
}
