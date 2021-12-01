//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ardupilotmega

// Status of key hardware.
type MessageHwstatus struct {
	// Board voltage.
	Vcc uint16 `mavname:"Vcc"`
	// I2C error count.
	I2cerr uint8 `mavname:"I2Cerr"`
}

// GetID implements the msg.Message interface.
func (*MessageHwstatus) GetID() uint32 {
	return 165
}
