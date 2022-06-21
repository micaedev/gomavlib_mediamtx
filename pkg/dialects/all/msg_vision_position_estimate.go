//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Local position/attitude estimate from a vision source.
type MessageVisionPositionEstimate struct {
	// Timestamp (UNIX time or time since system boot)
	Usec uint64
	// Local X position
	X float32
	// Local Y position
	Y float32
	// Local Z position
	Z float32
	// Roll angle
	Roll float32
	// Pitch angle
	Pitch float32
	// Yaw angle
	Yaw float32
	// Row-major representation of pose 6x6 cross-covariance matrix upper right triangle (states: x, y, z, roll, pitch, yaw; first six entries are the first ROW, next five entries are the second ROW, etc.). If unknown, assign NaN value to first element in the array.
	Covariance [21]float32 `mavext:"true"`
	// Estimate reset counter. This should be incremented when the estimate resets in any of the dimensions (position, velocity, attitude, angular speed). This is designed to be used when e.g an external SLAM system detects a loop-closure and the estimate jumps.
	ResetCounter uint8 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageVisionPositionEstimate) GetID() uint32 {
	return 102
}
