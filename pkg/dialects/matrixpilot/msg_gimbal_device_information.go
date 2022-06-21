//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// Information about a low level gimbal. This message should be requested by the gimbal manager or a ground station using MAV_CMD_REQUEST_MESSAGE. The maximum angles and rates are the limits by hardware. However, the limits by software used are likely different/smaller and dependent on mode/settings/etc..
type MessageGimbalDeviceInformation struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Name of the gimbal vendor.
	VendorName string `mavlen:"32"`
	// Name of the gimbal model.
	ModelName string `mavlen:"32"`
	// Custom name of the gimbal given to it by the user.
	CustomName string `mavlen:"32"`
	// Version of the gimbal firmware, encoded as: (Dev &amp; 0xff) &lt;&lt; 24 | (Patch &amp; 0xff) &lt;&lt; 16 | (Minor &amp; 0xff) &lt;&lt; 8 | (Major &amp; 0xff).
	FirmwareVersion uint32
	// Version of the gimbal hardware, encoded as: (Dev &amp; 0xff) &lt;&lt; 24 | (Patch &amp; 0xff) &lt;&lt; 16 | (Minor &amp; 0xff) &lt;&lt; 8 | (Major &amp; 0xff).
	HardwareVersion uint32
	// UID of gimbal hardware (0 if unknown).
	Uid uint64
	// Bitmap of gimbal capability flags.
	CapFlags GIMBAL_DEVICE_CAP_FLAGS `mavenum:"uint16"`
	// Bitmap for use for gimbal-specific capability flags.
	CustomCapFlags uint16
	// Minimum hardware roll angle (positive: rolling to the right, negative: rolling to the left)
	RollMin float32
	// Maximum hardware roll angle (positive: rolling to the right, negative: rolling to the left)
	RollMax float32
	// Minimum hardware pitch angle (positive: up, negative: down)
	PitchMin float32
	// Maximum hardware pitch angle (positive: up, negative: down)
	PitchMax float32
	// Minimum hardware yaw angle (positive: to the right, negative: to the left)
	YawMin float32
	// Maximum hardware yaw angle (positive: to the right, negative: to the left)
	YawMax float32
}

// GetID implements the message.Message interface.
func (*MessageGimbalDeviceInformation) GetID() uint32 {
	return 283
}
