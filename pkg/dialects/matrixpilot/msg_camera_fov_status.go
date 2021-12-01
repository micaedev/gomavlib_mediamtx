//autogenerated:yes
//nolint:golint,misspell,govet,lll
package matrixpilot

// Information about the field of view of a camera. Can be requested with a MAV_CMD_REQUEST_MESSAGE command.
type MessageCameraFovStatus struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Latitude of camera (INT32_MAX if unknown).
	LatCamera int32
	// Longitude of camera (INT32_MAX if unknown).
	LonCamera int32
	// Altitude (MSL) of camera (INT32_MAX if unknown).
	AltCamera int32
	// Latitude of center of image (INT32_MAX if unknown, INT32_MIN if at infinity, not intersecting with horizon).
	LatImage int32
	// Longitude of center of image (INT32_MAX if unknown, INT32_MIN if at infinity, not intersecting with horizon).
	LonImage int32
	// Altitude (MSL) of center of image (INT32_MAX if unknown, INT32_MIN if at infinity, not intersecting with horizon).
	AltImage int32
	// Quaternion of camera orientation (w, x, y, z order, zero-rotation is 1, 0, 0, 0)
	Q [4]float32
	// Horizontal field of view (NaN if unknown).
	Hfov float32
	// Vertical field of view (NaN if unknown).
	Vfov float32
}

// GetID implements the msg.Message interface.
func (*MessageCameraFovStatus) GetID() uint32 {
	return 271
}
