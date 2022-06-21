//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

// Request a chunk of a log
type MessageLogRequestData struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Log id (from LOG_ENTRY reply)
	Id uint16
	// Offset into the log
	Ofs uint32
	// Number of bytes
	Count uint32
}

// GetID implements the message.Message interface.
func (*MessageLogRequestData) GetID() uint32 {
	return 119
}
