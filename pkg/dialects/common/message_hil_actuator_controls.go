//autogenerated:yes
//nolint:revive,misspell,govet,lll
package common

// Sent from autopilot to simulation. Hardware in the loop control outputs. Alternative to HIL_CONTROLS.
type MessageHilActuatorControls struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Control outputs -1 .. 1. Channel assignment depends on the simulated hardware.
	Controls [16]float32
	// System mode. Includes arming state.
	Mode MAV_MODE_FLAG `mavenum:"uint8"`
	// Flags bitmask.
	Flags HIL_ACTUATOR_CONTROLS_FLAGS `mavenum:"uint64"`
}

// GetID implements the message.Message interface.
func (*MessageHilActuatorControls) GetID() uint32 {
	return 93
}
