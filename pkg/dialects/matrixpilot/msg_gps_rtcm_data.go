//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// RTCM message for injecting into the onboard GPS (used for DGPS)
type MessageGpsRtcmData struct {
	// LSB: 1 means message is fragmented, next 2 bits are the fragment ID, the remaining 5 bits are used for the sequence ID. Messages are only to be flushed to the GPS when the entire message has been reconstructed on the autopilot. The fragment ID specifies which order the fragments should be assembled into a buffer, while the sequence ID is used to detect a mismatch between different buffers. The buffer is considered fully reconstructed when either all 4 fragments are present, or all the fragments before the first fragment with a non full payload is received. This management is used to ensure that normal GPS operation doesn't corrupt RTCM data, and to recover from a unreliable transport delivery order.
	Flags uint8
	// data length
	Len uint8
	// RTCM message (may be fragmented)
	Data [180]uint8
}

// GetID implements the message.Message interface.
func (*MessageGpsRtcmData) GetID() uint32 {
	return 233
}
