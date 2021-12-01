//autogenerated:yes
//nolint:golint,misspell,govet,lll
package asluav

// Sets a desired vehicle position in a local north-east-down coordinate frame. Used by an external controller to command the vehicle (manual controller or other system).
type MessageSetPositionTargetLocalNed struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Valid options are: MAV_FRAME_LOCAL_NED = 1, MAV_FRAME_LOCAL_OFFSET_NED = 7, MAV_FRAME_BODY_NED = 8, MAV_FRAME_BODY_OFFSET_NED = 9
	CoordinateFrame MAV_FRAME `mavenum:"uint8"`
	// Bitmap to indicate which dimensions should be ignored by the vehicle.
	TypeMask POSITION_TARGET_TYPEMASK `mavenum:"uint16"`
	// X Position in NED frame
	X float32
	// Y Position in NED frame
	Y float32
	// Z Position in NED frame (note, altitude is negative in NED)
	Z float32
	// X velocity in NED frame
	Vx float32
	// Y velocity in NED frame
	Vy float32
	// Z velocity in NED frame
	Vz float32
	// X acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afx float32
	// Y acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afy float32
	// Z acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afz float32
	// yaw setpoint
	Yaw float32
	// yaw rate setpoint
	YawRate float32
}

// GetID implements the msg.Message interface.
func (*MessageSetPositionTargetLocalNed) GetID() uint32 {
	return 84
}
