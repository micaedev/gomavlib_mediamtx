//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

// Status generated in each node in the communication chain and injected into MAVLink stream.
type MessageLinkNodeStatus struct {
	// Timestamp (time since system boot).
	Timestamp uint64
	// Remaining free transmit buffer space
	TxBuf uint8
	// Remaining free receive buffer space
	RxBuf uint8
	// Transmit rate
	TxRate uint32
	// Receive rate
	RxRate uint32
	// Number of bytes that could not be parsed correctly.
	RxParseErr uint16
	// Transmit buffer overflows. This number wraps around as it reaches UINT16_MAX
	TxOverflows uint16
	// Receive buffer overflows. This number wraps around as it reaches UINT16_MAX
	RxOverflows uint16
	// Messages sent
	MessagesSent uint32
	// Messages received (estimated from counting seq)
	MessagesReceived uint32
	// Messages lost (estimated from counting seq)
	MessagesLost uint32
}

// GetID implements the message.Message interface.
func (*MessageLinkNodeStatus) GetID() uint32 {
	return 8
}
