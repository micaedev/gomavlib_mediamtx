//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// Temperature and humidity from hygrometer.
type MessageHygrometerSensor struct {
	// Hygrometer ID
	Id uint8
	// Temperature
	Temperature int16
	// Humidity
	Humidity uint16
}

// GetID implements the message.Message interface.
func (*MessageHygrometerSensor) GetID() uint32 {
	return 12920
}
