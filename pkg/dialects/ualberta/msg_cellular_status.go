//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

// Report current used cellular network status
type MessageCellularStatus struct {
	// Cellular modem status
	Status CELLULAR_STATUS_FLAG `mavenum:"uint8"`
	// Failure reason when status in in CELLUAR_STATUS_FAILED
	FailureReason CELLULAR_NETWORK_FAILED_REASON `mavenum:"uint8"`
	// Cellular network radio type: gsm, cdma, lte...
	Type CELLULAR_NETWORK_RADIO_TYPE `mavenum:"uint8"`
	// Signal quality in percent. If unknown, set to UINT8_MAX
	Quality uint8
	// Mobile country code. If unknown, set to UINT16_MAX
	Mcc uint16
	// Mobile network code. If unknown, set to UINT16_MAX
	Mnc uint16
	// Location area code. If unknown, set to 0
	Lac uint16
}

// GetID implements the message.Message interface.
func (*MessageCellularStatus) GetID() uint32 {
	return 334
}
