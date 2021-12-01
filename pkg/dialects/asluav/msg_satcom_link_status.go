//autogenerated:yes
//nolint:golint,misspell,govet,lll
package asluav

// Status of the SatCom link
type MessageSatcomLinkStatus struct {
	// Timestamp
	Timestamp uint64
	// Timestamp of the last successful sbd session
	LastHeartbeat uint64
	// Number of failed sessions
	FailedSessions uint16
	// Number of successful sessions
	SuccessfulSessions uint16
	// Signal quality
	SignalQuality uint8
	// Ring call pending
	RingPending uint8
	// Transmission session pending
	TxSessionPending uint8
	// Receiving session pending
	RxSessionPending uint8
}

// GetID implements the msg.Message interface.
func (*MessageSatcomLinkStatus) GetID() uint32 {
	return 8015
}
