//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

// Camera tracking status, sent while in active tracking. Use MAV_CMD_SET_MESSAGE_INTERVAL to define message interval.
type MessageCameraTrackingImageStatus struct {
	// Current tracking status
	TrackingStatus CAMERA_TRACKING_STATUS_FLAGS `mavenum:"uint8"`
	// Current tracking mode
	TrackingMode CAMERA_TRACKING_MODE `mavenum:"uint8"`
	// Defines location of target data
	TargetData CAMERA_TRACKING_TARGET_DATA `mavenum:"uint8"`
	// Current tracked point x value if CAMERA_TRACKING_MODE_POINT (normalized 0..1, 0 is left, 1 is right), NAN if unknown
	PointX float32
	// Current tracked point y value if CAMERA_TRACKING_MODE_POINT (normalized 0..1, 0 is top, 1 is bottom), NAN if unknown
	PointY float32
	// Current tracked radius if CAMERA_TRACKING_MODE_POINT (normalized 0..1, 0 is image left, 1 is image right), NAN if unknown
	Radius float32
	// Current tracked rectangle top x value if CAMERA_TRACKING_MODE_RECTANGLE (normalized 0..1, 0 is left, 1 is right), NAN if unknown
	RecTopX float32
	// Current tracked rectangle top y value if CAMERA_TRACKING_MODE_RECTANGLE (normalized 0..1, 0 is top, 1 is bottom), NAN if unknown
	RecTopY float32
	// Current tracked rectangle bottom x value if CAMERA_TRACKING_MODE_RECTANGLE (normalized 0..1, 0 is left, 1 is right), NAN if unknown
	RecBottomX float32
	// Current tracked rectangle bottom y value if CAMERA_TRACKING_MODE_RECTANGLE (normalized 0..1, 0 is top, 1 is bottom), NAN if unknown
	RecBottomY float32
}

// GetID implements the message.Message interface.
func (*MessageCameraTrackingImageStatus) GetID() uint32 {
	return 275
}
