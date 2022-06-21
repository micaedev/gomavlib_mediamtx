//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

// Status text message. These messages are printed in yellow in the COMM console of QGroundControl. WARNING: They consume quite some bandwidth, so use only for important status and error messages. If implemented wisely, these messages are buffered on the MCU and sent only at a limited rate (e.g. 10 Hz).
type MessageStatustext struct {
	// Severity of status. Relies on the definitions within RFC-5424.
	Severity MAV_SEVERITY `mavenum:"uint8"`
	// Status text message, without null termination character
	Text string `mavlen:"50"`
	// Unique (opaque) identifier for this statustext message.  May be used to reassemble a logical long-statustext message from a sequence of chunks.  A value of zero indicates this is the only chunk in the sequence and the message can be emitted immediately.
	Id uint16 `mavext:"true"`
	// This chunk's sequence number; indexing is from zero.  Any null character in the text field is taken to mean this was the last chunk.
	ChunkSeq uint8 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageStatustext) GetID() uint32 {
	return 253
}
