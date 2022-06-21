//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

//
// Sets the home position.
// 	The home position is the default position that the system will return to and land on.
// The position is set automatically by the system during the takeoff (and may also be set using this message).
// The global and local positions encode the position in the respective coordinate frames, while the q parameter encodes the orientation of the surface.
// Under normal conditions it describes the heading and terrain slope, which can be used by the aircraft to adjust the approach.
// The approach 3D vector describes the point to which the system should fly in normal flight mode and then perform a landing sequence along the vector.
// Note: the current home position may be emitted in a HOME_POSITION message on request (using MAV_CMD_REQUEST_MESSAGE with param1=242).
type MessageSetHomePosition struct {
	// System ID.
	TargetSystem uint8
	// Latitude (WGS84)
	Latitude int32
	// Longitude (WGS84)
	Longitude int32
	// Altitude (MSL). Positive for up.
	Altitude int32
	// Local X position of this position in the local coordinate frame
	X float32
	// Local Y position of this position in the local coordinate frame
	Y float32
	// Local Z position of this position in the local coordinate frame
	Z float32
	// World to surface normal and heading transformation of the takeoff position. Used to indicate the heading and slope of the ground
	Q [4]float32
	// Local X position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	ApproachX float32
	// Local Y position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	ApproachY float32
	// Local Z position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	ApproachZ float32
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageSetHomePosition) GetID() uint32 {
	return 243
}
