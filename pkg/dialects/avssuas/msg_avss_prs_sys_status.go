//autogenerated:yes
//nolint:revive,misspell,govet,lll
package avssuas

// AVSS PRS system status.
type MessageAvssPrsSysStatus struct {
	// Timestamp (time since PRS boot).
	TimeBootMs uint32
	// PRS error statuses
	ErrorStatus uint32
	// Estimated battery run-time without a remote connection and PRS battery voltage
	BatteryStatus uint32
	// PRS arm statuses
	ArmStatus uint8
	// PRS battery charge statuses
	ChargeStatus uint8
}

// GetID implements the message.Message interface.
func (*MessageAvssPrsSysStatus) GetID() uint32 {
	return 60050
}
