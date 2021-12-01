//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ardupilotmega

// Camera Capture Feedback.
type MessageCameraFeedback struct {
	// Image timestamp (since UNIX epoch), as passed in by CAMERA_STATUS message (or autopilot if no CCB).
	TimeUsec uint64
	// System ID.
	TargetSystem uint8
	// Camera ID.
	CamIdx uint8
	// Image index.
	ImgIdx uint16
	// Latitude.
	Lat int32
	// Longitude.
	Lng int32
	// Altitude (MSL).
	AltMsl float32
	// Altitude (Relative to HOME location).
	AltRel float32
	// Camera Roll angle (earth frame, +-180).
	Roll float32
	// Camera Pitch angle (earth frame, +-180).
	Pitch float32
	// Camera Yaw (earth frame, 0-360, true).
	Yaw float32
	// Focal Length.
	FocLen float32
	// Feedback flags.
	Flags CAMERA_FEEDBACK_FLAGS `mavenum:"uint8"`
	// Completed image captures.
	CompletedCaptures uint16 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageCameraFeedback) GetID() uint32 {
	return 180
}
