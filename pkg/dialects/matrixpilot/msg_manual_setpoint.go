//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// Setpoint in roll, pitch, yaw and thrust from the operator
type MessageManualSetpoint struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Desired roll rate
	Roll float32
	// Desired pitch rate
	Pitch float32
	// Desired yaw rate
	Yaw float32
	// Collective thrust, normalized to 0 .. 1
	Thrust float32
	// Flight mode switch position, 0.. 255
	ModeSwitch uint8
	// Override mode switch position, 0.. 255
	ManualOverrideSwitch uint8
}

// GetID implements the message.Message interface.
func (*MessageManualSetpoint) GetID() uint32 {
	return 81
}
