//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

// Set the vehicle attitude and body angular rates.
type MessageActuatorControlTarget struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Actuator group. The "_mlx" indicates this is a multi-instance message and a MAVLink parser should use this field to difference between instances.
	GroupMlx uint8
	// Actuator controls. Normed to -1..+1 where 0 is neutral position. Throttle for single rotation direction motors is 0..1, negative range for reverse direction. Standard mapping for attitude controls (group 0): (index 0-7): roll, pitch, yaw, throttle, flaps, spoilers, airbrakes, landing gear. Load a pass-through mixer to repurpose them as generic outputs.
	Controls [8]float32
}

// GetID implements the message.Message interface.
func (*MessageActuatorControlTarget) GetID() uint32 {
	return 140
}
