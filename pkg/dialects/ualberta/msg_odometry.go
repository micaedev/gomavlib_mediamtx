//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ualberta

// Odometry message to communicate odometry information with an external interface. Fits ROS REP 147 standard for aerial vehicles (http://www.ros.org/reps/rep-0147.html).
type MessageOdometry struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Coordinate frame of reference for the pose data.
	FrameId MAV_FRAME `mavenum:"uint8"`
	// Coordinate frame of reference for the velocity in free space (twist) data.
	ChildFrameId MAV_FRAME `mavenum:"uint8"`
	// X Position
	X float32
	// Y Position
	Y float32
	// Z Position
	Z float32
	// Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation)
	Q [4]float32
	// X linear speed
	Vx float32
	// Y linear speed
	Vy float32
	// Z linear speed
	Vz float32
	// Roll angular speed
	Rollspeed float32
	// Pitch angular speed
	Pitchspeed float32
	// Yaw angular speed
	Yawspeed float32
	// Row-major representation of a 6x6 pose cross-covariance matrix upper right triangle (states: x, y, z, roll, pitch, yaw; first six entries are the first ROW, next five entries are the second ROW, etc.). If unknown, assign NaN value to first element in the array.
	PoseCovariance [21]float32
	// Row-major representation of a 6x6 velocity cross-covariance matrix upper right triangle (states: vx, vy, vz, rollspeed, pitchspeed, yawspeed; first six entries are the first ROW, next five entries are the second ROW, etc.). If unknown, assign NaN value to first element in the array.
	VelocityCovariance [21]float32
	// Estimate reset counter. This should be incremented when the estimate resets in any of the dimensions (position, velocity, attitude, angular speed). This is designed to be used when e.g an external SLAM system detects a loop-closure and the estimate jumps.
	ResetCounter uint8 `mavext:"true"`
	// Type of estimator that is providing the odometry.
	EstimatorType MAV_ESTIMATOR_TYPE `mavenum:"uint8" mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageOdometry) GetID() uint32 {
	return 331
}
