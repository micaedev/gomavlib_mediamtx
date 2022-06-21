//autogenerated:yes
//nolint:revive,misspell,govet,lll
package paparazzi

// ESC information for higher rate streaming. Recommended streaming rate is ~10 Hz. Information that changes more slowly is sent in ESC_INFO. It should typically only be streamed on high-bandwidth links (i.e. to a companion computer).
type MessageEscStatus struct {
	// Index of the first ESC in this message. minValue = 0, maxValue = 60, increment = 4.
	Index uint8
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number.
	TimeUsec uint64
	// Reported motor RPM from each ESC (negative for reverse rotation).
	Rpm [4]int32
	// Voltage measured from each ESC.
	Voltage [4]float32
	// Current measured from each ESC.
	Current [4]float32
}

// GetID implements the message.Message interface.
func (*MessageEscStatus) GetID() uint32 {
	return 291
}
