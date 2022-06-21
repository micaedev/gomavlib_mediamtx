//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Camera-IMU triggering and synchronisation message.
type MessageCameraTrigger struct {
	// Timestamp for image frame (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Image frame sequence
	Seq uint32
}

// GetID implements the message.Message interface.
func (*MessageCameraTrigger) GetID() uint32 {
	return 112
}
