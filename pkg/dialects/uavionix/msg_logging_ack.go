//autogenerated:yes
//nolint:golint,misspell,govet,lll
package uavionix

// An ack for a LOGGING_DATA_ACKED message
type MessageLoggingAck struct {
	// system ID of the target
	TargetSystem uint8
	// component ID of the target
	TargetComponent uint8
	// sequence number (must match the one in LOGGING_DATA_ACKED)
	Sequence uint16
}

// GetID implements the msg.Message interface.
func (*MessageLoggingAck) GetID() uint32 {
	return 268
}
