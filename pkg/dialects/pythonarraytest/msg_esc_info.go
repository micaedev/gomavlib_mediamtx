//autogenerated:yes
//nolint:golint,misspell,govet,lll
package pythonarraytest

// ESC information for lower rate streaming. Recommended streaming rate 1Hz. See ESC_STATUS for higher-rate ESC data.
type MessageEscInfo struct {
	// Index of the first ESC in this message. minValue = 0, maxValue = 60, increment = 4.
	Index uint8
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number.
	TimeUsec uint64
	// Counter of data packets received.
	Counter uint16
	// Total number of ESCs in all messages of this type. Message fields with an index higher than this should be ignored because they contain invalid data.
	Count uint8
	// Connection type protocol for all ESC.
	ConnectionType ESC_CONNECTION_TYPE `mavenum:"uint8"`
	// Information regarding online/offline status of each ESC.
	Info uint8
	// Bitmap of ESC failure flags.
	FailureFlags [4]ESC_FAILURE_FLAGS `mavenum:"uint16"`
	// Number of reported errors by each ESC since boot.
	ErrorCount [4]uint32
	// Temperature of each ESC. INT16_MAX: if data not supplied by ESC.
	Temperature [4]int16
}

// GetID implements the msg.Message interface.
func (*MessageEscInfo) GetID() uint32 {
	return 290
}
