//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// The raw values of the actuator outputs (e.g. on Pixhawk, from MAIN, AUX ports). This message supersedes SERVO_OUTPUT_RAW.
type MessageActuatorOutputStatus struct {
	// Timestamp (since system boot).
	TimeUsec uint64
	// Active outputs
	Active uint32
	// Servo / motor output array values. Zero values indicate unused channels.
	Actuator [32]float32
}

// GetID implements the message.Message interface.
func (*MessageActuatorOutputStatus) GetID() uint32 {
	return 375
}
