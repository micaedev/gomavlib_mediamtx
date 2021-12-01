//autogenerated:yes
//nolint:golint,misspell,govet,lll
package common

// Accept / deny control of this MAV
type MessageChangeOperatorControlAck struct {
	// ID of the GCS this message
	GcsSystemId uint8
	// 0: request control of this MAV, 1: Release control of this MAV
	ControlRequest uint8
	// 0: ACK, 1: NACK: Wrong passkey, 2: NACK: Unsupported passkey encryption method, 3: NACK: Already under control
	Ack uint8
}

// GetID implements the msg.Message interface.
func (*MessageChangeOperatorControlAck) GetID() uint32 {
	return 6
}
