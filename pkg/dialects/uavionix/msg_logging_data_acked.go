//autogenerated:yes
//nolint:revive,misspell,govet,lll
package uavionix

// A message containing logged data which requires a LOGGING_ACK to be sent back
type MessageLoggingDataAcked struct {
	// system ID of the target
	TargetSystem uint8
	// component ID of the target
	TargetComponent uint8
	// sequence number (can wrap)
	Sequence uint16
	// data length
	Length uint8
	// offset into data where first message starts. This can be used for recovery, when a previous message got lost (set to UINT8_MAX if no start exists).
	FirstMessageOffset uint8
	// logged data
	Data [249]uint8
}

// GetID implements the message.Message interface.
func (*MessageLoggingDataAcked) GetID() uint32 {
	return 267
}
