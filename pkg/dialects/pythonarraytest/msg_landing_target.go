//autogenerated:yes
//nolint:golint,misspell,govet,lll
package pythonarraytest

// The location of a landing target. See: https://mavlink.io/en/services/landing_target.html
type MessageLandingTarget struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// The ID of the target if multiple targets are present
	TargetNum uint8
	// Coordinate frame used for following fields.
	Frame MAV_FRAME `mavenum:"uint8"`
	// X-axis angular offset of the target from the center of the image
	AngleX float32
	// Y-axis angular offset of the target from the center of the image
	AngleY float32
	// Distance to the target from the vehicle
	Distance float32
	// Size of target along x-axis
	SizeX float32
	// Size of target along y-axis
	SizeY float32
	// X Position of the landing target in MAV_FRAME
	X float32 `mavext:"true"`
	// Y Position of the landing target in MAV_FRAME
	Y float32 `mavext:"true"`
	// Z Position of the landing target in MAV_FRAME
	Z float32 `mavext:"true"`
	// Quaternion of landing target orientation (w, x, y, z order, zero-rotation is 1, 0, 0, 0)
	Q [4]float32 `mavext:"true"`
	// Type of landing target
	Type LANDING_TARGET_TYPE `mavenum:"uint8" mavext:"true"`
	// Boolean indicating whether the position fields (x, y, z, q, type) contain valid target position information (valid: 1, invalid: 0). Default is 0 (invalid).
	PositionValid uint8 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageLandingTarget) GetID() uint32 {
	return 149
}
