//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// EKF Status message including flags and variances.
type MessageEkfStatusReport struct {
	// Flags.
	Flags EKF_STATUS_FLAGS `mavenum:"uint16"`
	// Velocity variance.
	VelocityVariance float32
	// Horizontal Position variance.
	PosHorizVariance float32
	// Vertical Position variance.
	PosVertVariance float32
	// Compass variance.
	CompassVariance float32
	// Terrain Altitude variance.
	TerrainAltVariance float32
	// Airspeed variance.
	AirspeedVariance float32 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageEkfStatusReport) GetID() uint32 {
	return 193
}
