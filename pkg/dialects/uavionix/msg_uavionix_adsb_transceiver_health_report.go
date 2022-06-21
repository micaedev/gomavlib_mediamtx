//autogenerated:yes
//nolint:revive,misspell,govet,lll
package uavionix

// Transceiver heartbeat with health report (updated every 10s)
type MessageUavionixAdsbTransceiverHealthReport struct {
	// ADS-B transponder messages
	Rfhealth UAVIONIX_ADSB_RF_HEALTH `mavenum:"uint8" mavname:"rfHealth"`
}

// GetID implements the message.Message interface.
func (*MessageUavionixAdsbTransceiverHealthReport) GetID() uint32 {
	return 10003
}
