//autogenerated:yes
//nolint:golint,misspell,govet,lll
package asluav

// Status of GSM modem (connected to onboard computer)
type MessageGsmLinkStatus struct {
	// Timestamp (of OBC)
	Timestamp uint64
	// GSM modem used
	GsmModemType GSM_MODEM_TYPE `mavenum:"uint8"`
	// GSM link type
	GsmLinkType GSM_LINK_TYPE `mavenum:"uint8"`
	// RSSI as reported by modem (unconverted)
	Rssi uint8
	// RSRP (LTE) or RSCP (WCDMA) as reported by modem (unconverted)
	RsrpRscp uint8
	// SINR (LTE) or ECIO (WCDMA) as reported by modem (unconverted)
	SinrEcio uint8
	// RSRQ (LTE only) as reported by modem (unconverted)
	Rsrq uint8
}

// GetID implements the msg.Message interface.
func (*MessageGsmLinkStatus) GetID() uint32 {
	return 8014
}
