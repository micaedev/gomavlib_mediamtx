//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

// The filtered local position (e.g. fused computer vision and accelerometers). Coordinate frame is right-handed, Z-axis down (aeronautical frame, NED / north-east-down convention)
type MessageLocalPositionNed struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// X Position
	X float32
	// Y Position
	Y float32
	// Z Position
	Z float32
	// X Speed
	Vx float32
	// Y Speed
	Vy float32
	// Z Speed
	Vz float32
}

// GetID implements the message.Message interface.
func (*MessageLocalPositionNed) GetID() uint32 {
	return 32
}
