//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ualberta

// Streamed from drone to report progress of terrain map download (initiated by TERRAIN_REQUEST), or sent as a response to a TERRAIN_CHECK request. See terrain protocol docs: https://mavlink.io/en/services/terrain.html
type MessageTerrainReport struct {
	// Latitude
	Lat int32
	// Longitude
	Lon int32
	// grid spacing (zero if terrain at this location unavailable)
	Spacing uint16
	// Terrain height MSL
	TerrainHeight float32
	// Current vehicle height above lat/lon terrain height
	CurrentHeight float32
	// Number of 4x4 terrain blocks waiting to be received or read from disk
	Pending uint16
	// Number of 4x4 terrain blocks in memory
	Loaded uint16
}

// GetID implements the msg.Message interface.
func (*MessageTerrainReport) GetID() uint32 {
	return 136
}
