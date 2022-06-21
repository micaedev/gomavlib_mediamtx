//autogenerated:yes
//nolint:revive,misspell,govet,lll
package paparazzi

// Camera tracking status, sent while in active tracking. Use MAV_CMD_SET_MESSAGE_INTERVAL to define message interval.
type MessageCameraTrackingGeoStatus struct {
	// Current tracking status
	TrackingStatus CAMERA_TRACKING_STATUS_FLAGS `mavenum:"uint8"`
	// Latitude of tracked object
	Lat int32
	// Longitude of tracked object
	Lon int32
	// Altitude of tracked object(AMSL, WGS84)
	Alt float32
	// Horizontal accuracy. NAN if unknown
	HAcc float32
	// Vertical accuracy. NAN if unknown
	VAcc float32
	// North velocity of tracked object. NAN if unknown
	VelN float32
	// East velocity of tracked object. NAN if unknown
	VelE float32
	// Down velocity of tracked object. NAN if unknown
	VelD float32
	// Velocity accuracy. NAN if unknown
	VelAcc float32
	// Distance between camera and tracked object. NAN if unknown
	Dist float32
	// Heading in radians, in NED. NAN if unknown
	Hdg float32
	// Accuracy of heading, in NED. NAN if unknown
	HdgAcc float32
}

// GetID implements the message.Message interface.
func (*MessageCameraTrackingGeoStatus) GetID() uint32 {
	return 276
}
