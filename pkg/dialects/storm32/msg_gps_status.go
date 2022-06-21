//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// The positioning status, as reported by GPS. This message is intended to display status information about each satellite visible to the receiver. See message GLOBAL_POSITION_INT for the global position estimate. This message can contain information for up to 20 satellites.
type MessageGpsStatus struct {
	// Number of satellites visible
	SatellitesVisible uint8
	// Global satellite ID
	SatellitePrn [20]uint8
	// 0: Satellite not used, 1: used for localization
	SatelliteUsed [20]uint8
	// Elevation (0: right on top of receiver, 90: on the horizon) of satellite
	SatelliteElevation [20]uint8
	// Direction of satellite, 0: 0 deg, 255: 360 deg.
	SatelliteAzimuth [20]uint8
	// Signal to noise ratio of satellite
	SatelliteSnr [20]uint8
}

// GetID implements the message.Message interface.
func (*MessageGpsStatus) GetID() uint32 {
	return 25
}
