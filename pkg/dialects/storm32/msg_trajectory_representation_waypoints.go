//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Describe a trajectory using an array of up-to 5 waypoints in the local frame (MAV_FRAME_LOCAL_NED).
type MessageTrajectoryRepresentationWaypoints struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Number of valid points (up-to 5 waypoints are possible)
	ValidPoints uint8
	// X-coordinate of waypoint, set to NaN if not being used
	PosX [5]float32
	// Y-coordinate of waypoint, set to NaN if not being used
	PosY [5]float32
	// Z-coordinate of waypoint, set to NaN if not being used
	PosZ [5]float32
	// X-velocity of waypoint, set to NaN if not being used
	VelX [5]float32
	// Y-velocity of waypoint, set to NaN if not being used
	VelY [5]float32
	// Z-velocity of waypoint, set to NaN if not being used
	VelZ [5]float32
	// X-acceleration of waypoint, set to NaN if not being used
	AccX [5]float32
	// Y-acceleration of waypoint, set to NaN if not being used
	AccY [5]float32
	// Z-acceleration of waypoint, set to NaN if not being used
	AccZ [5]float32
	// Yaw angle, set to NaN if not being used
	PosYaw [5]float32
	// Yaw rate, set to NaN if not being used
	VelYaw [5]float32
	// MAV_CMD command id of waypoint, set to UINT16_MAX if not being used.
	Command [5]MAV_CMD `mavenum:"uint16"`
}

// GetID implements the message.Message interface.
func (*MessageTrajectoryRepresentationWaypoints) GetID() uint32 {
	return 332
}
