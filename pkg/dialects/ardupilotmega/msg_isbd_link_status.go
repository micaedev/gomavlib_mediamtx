//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ardupilotmega

// Status of the Iridium SBD link.
type MessageIsbdLinkStatus struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Timestamp uint64
	// Timestamp of the last successful sbd session. The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	LastHeartbeat uint64
	// Number of failed SBD sessions.
	FailedSessions uint16
	// Number of successful SBD sessions.
	SuccessfulSessions uint16
	// Signal quality equal to the number of bars displayed on the ISU signal strength indicator. Range is 0 to 5, where 0 indicates no signal and 5 indicates maximum signal strength.
	SignalQuality uint8
	// 1: Ring call pending, 0: No call pending.
	RingPending uint8
	// 1: Transmission session pending, 0: No transmission session pending.
	TxSessionPending uint8
	// 1: Receiving session pending, 0: No receiving session pending.
	RxSessionPending uint8
}

// GetID implements the message.Message interface.
func (*MessageIsbdLinkStatus) GetID() uint32 {
	return 335
}
