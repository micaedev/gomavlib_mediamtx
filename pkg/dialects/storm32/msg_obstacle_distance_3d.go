//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Obstacle located as a 3D vector.
type MessageObstacleDistance_3d struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Class id of the distance sensor type.
	SensorType MAV_DISTANCE_SENSOR `mavenum:"uint8"`
	// Coordinate frame of reference.
	Frame MAV_FRAME `mavenum:"uint8"`
	// Unique ID given to each obstacle so that its movement can be tracked. Use UINT16_MAX if object ID is unknown or cannot be determined.
	ObstacleId uint16
	// X position of the obstacle.
	X float32
	// Y position of the obstacle.
	Y float32
	// Z position of the obstacle.
	Z float32
	// Minimum distance the sensor can measure.
	MinDistance float32
	// Maximum distance the sensor can measure.
	MaxDistance float32
}

// GetID implements the message.Message interface.
func (*MessageObstacleDistance_3d) GetID() uint32 {
	return 11037
}
