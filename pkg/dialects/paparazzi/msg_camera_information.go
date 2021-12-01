//autogenerated:yes
//nolint:golint,misspell,govet,lll
package paparazzi

// Information about a camera. Can be requested with a MAV_CMD_REQUEST_MESSAGE command.
type MessageCameraInformation struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Name of the camera vendor
	VendorName [32]uint8
	// Name of the camera model
	ModelName [32]uint8
	// Version of the camera firmware, encoded as: (Dev &amp; 0xff) &lt;&lt; 24 | (Patch &amp; 0xff) &lt;&lt; 16 | (Minor &amp; 0xff) &lt;&lt; 8 | (Major &amp; 0xff)
	FirmwareVersion uint32
	// Focal length
	FocalLength float32
	// Image sensor size horizontal
	SensorSizeH float32
	// Image sensor size vertical
	SensorSizeV float32
	// Horizontal image resolution
	ResolutionH uint16
	// Vertical image resolution
	ResolutionV uint16
	// Reserved for a lens ID
	LensId uint8
	// Bitmap of camera capability flags.
	Flags CAMERA_CAP_FLAGS `mavenum:"uint32"`
	// Camera definition version (iteration)
	CamDefinitionVersion uint16
	// Camera definition URI (if any, otherwise only basic functions will be available). HTTP- (http://) and MAVLink FTP- (mavlinkftp://) formatted URIs are allowed (and both must be supported by any GCS that implements the Camera Protocol). The definition file may be xz compressed, which will be indicated by the file extension .xml.xz (a GCS that implements the protocol must support decompressing the file).
	CamDefinitionUri string `mavlen:"140"`
}

// GetID implements the msg.Message interface.
func (*MessageCameraInformation) GetID() uint32 {
	return 259
}
