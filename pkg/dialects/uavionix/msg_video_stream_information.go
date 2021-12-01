//autogenerated:yes
//nolint:golint,misspell,govet,lll
package uavionix

// Information about video stream. It may be requested using MAV_CMD_REQUEST_MESSAGE, where param2 indicates the video stream id: 0 for all streams, 1 for first, 2 for second, etc.
type MessageVideoStreamInformation struct {
	// Video Stream ID (1 for first, 2 for second, etc.)
	StreamId uint8
	// Number of streams available.
	Count uint8
	// Type of stream.
	Type VIDEO_STREAM_TYPE `mavenum:"uint8"`
	// Bitmap of stream status flags.
	Flags VIDEO_STREAM_STATUS_FLAGS `mavenum:"uint16"`
	// Frame rate.
	Framerate float32
	// Horizontal resolution.
	ResolutionH uint16
	// Vertical resolution.
	ResolutionV uint16
	// Bit rate.
	Bitrate uint32
	// Video image rotation clockwise.
	Rotation uint16
	// Horizontal Field of view.
	Hfov uint16
	// Stream name.
	Name string `mavlen:"32"`
	// Video stream URI (TCP or RTSP URI ground station should connect to) or port number (UDP port ground station should listen to).
	Uri string `mavlen:"160"`
}

// GetID implements the msg.Message interface.
func (*MessageVideoStreamInformation) GetID() uint32 {
	return 269
}
