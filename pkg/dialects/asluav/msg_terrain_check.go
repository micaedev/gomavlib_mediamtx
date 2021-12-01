//autogenerated:yes
//nolint:golint,misspell,govet,lll
package asluav

// Request that the vehicle report terrain height at the given location (expected response is a TERRAIN_REPORT). Used by GCS to check if vehicle has all terrain data needed for a mission.
type MessageTerrainCheck struct {
	// Latitude
	Lat int32
	// Longitude
	Lon int32
}

// GetID implements the msg.Message interface.
func (*MessageTerrainCheck) GetID() uint32 {
	return 135
}
