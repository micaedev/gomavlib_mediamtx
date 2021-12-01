//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ardupilotmega

// RPM sensor output.
type MessageRpm struct {
	// RPM Sensor1.
	Rpm1 float32
	// RPM Sensor2.
	Rpm2 float32
}

// GetID implements the msg.Message interface.
func (*MessageRpm) GetID() uint32 {
	return 226
}
