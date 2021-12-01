//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ardupilotmega

// Control a serial port. This can be used for raw access to an onboard serial peripheral such as a GPS or telemetry radio. It is designed to make it possible to update the devices firmware via MAVLink messages or change the devices settings. A message with zero bytes can be used to change just the baudrate.
type MessageSerialControl struct {
	// Serial control device type.
	Device SERIAL_CONTROL_DEV `mavenum:"uint8"`
	// Bitmap of serial control flags.
	Flags SERIAL_CONTROL_FLAG `mavenum:"uint8"`
	// Timeout for reply data
	Timeout uint16
	// Baudrate of transfer. Zero means no change.
	Baudrate uint32
	// how many bytes in this transfer
	Count uint8
	// serial data
	Data [70]uint8
	// System ID
	TargetSystem uint8 `mavext:"true"`
	// Component ID
	TargetComponent uint8 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageSerialControl) GetID() uint32 {
	return 126
}
