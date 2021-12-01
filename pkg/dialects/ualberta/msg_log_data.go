//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ualberta

// Reply to LOG_REQUEST_DATA
type MessageLogData struct {
	// Log id (from LOG_ENTRY reply)
	Id uint16
	// Offset into the log
	Ofs uint32
	// Number of bytes (zero for end of log)
	Count uint8
	// log data
	Data [90]uint8
}

// GetID implements the msg.Message interface.
func (*MessageLogData) GetID() uint32 {
	return 120
}
