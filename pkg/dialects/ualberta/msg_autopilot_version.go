//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ualberta

// Version and capability of autopilot software. This should be emitted in response to a request with MAV_CMD_REQUEST_MESSAGE.
type MessageAutopilotVersion struct {
	// Bitmap of capabilities
	Capabilities MAV_PROTOCOL_CAPABILITY `mavenum:"uint64"`
	// Firmware version number
	FlightSwVersion uint32
	// Middleware version number
	MiddlewareSwVersion uint32
	// Operating system version number
	OsSwVersion uint32
	// HW / board version (last 8 bits should be silicon ID, if any). The first 16 bits of this field specify https://github.com/PX4/PX4-Bootloader/blob/master/board_types.txt
	BoardVersion uint32
	// Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases.
	FlightCustomVersion [8]uint8
	// Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases.
	MiddlewareCustomVersion [8]uint8
	// Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases.
	OsCustomVersion [8]uint8
	// ID of the board vendor
	VendorId uint16
	// ID of the product
	ProductId uint16
	// UID if provided by hardware (see uid2)
	Uid uint64
	// UID if provided by hardware (supersedes the uid field. If this is non-zero, use this field, otherwise use uid)
	Uid2 [18]uint8 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageAutopilotVersion) GetID() uint32 {
	return 148
}
