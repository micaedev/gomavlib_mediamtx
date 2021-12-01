//autogenerated:yes
//nolint:golint,misspell,govet,lll
package asluav

// Request for terrain data and terrain status. See terrain protocol docs: https://mavlink.io/en/services/terrain.html
type MessageTerrainRequest struct {
	// Latitude of SW corner of first grid
	Lat int32
	// Longitude of SW corner of first grid
	Lon int32
	// Grid spacing
	GridSpacing uint16
	// Bitmask of requested 4x4 grids (row major 8x7 array of grids, 56 bits)
	Mask uint64
}

// GetID implements the msg.Message interface.
func (*MessageTerrainRequest) GetID() uint32 {
	return 133
}
