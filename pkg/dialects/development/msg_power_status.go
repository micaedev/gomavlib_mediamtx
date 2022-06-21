//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

// Power supply status
type MessagePowerStatus struct {
	// 5V rail voltage.
	Vcc uint16 `mavname:"Vcc"`
	// Servo rail voltage.
	Vservo uint16 `mavname:"Vservo"`
	// Bitmap of power supply status flags.
	Flags MAV_POWER_STATUS `mavenum:"uint16"`
}

// GetID implements the message.Message interface.
func (*MessagePowerStatus) GetID() uint32 {
	return 125
}
