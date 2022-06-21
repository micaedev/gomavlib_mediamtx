//autogenerated:yes
//nolint:revive,misspell,govet,lll
package avssuas

// Optical flow from a flow sensor (e.g. optical mouse sensor)
type MessageOpticalFlow struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Sensor ID
	SensorId uint8
	// Flow in x-sensor direction
	FlowX int16
	// Flow in y-sensor direction
	FlowY int16
	// Flow in x-sensor direction, angular-speed compensated
	FlowCompMX float32
	// Flow in y-sensor direction, angular-speed compensated
	FlowCompMY float32
	// Optical flow quality / confidence. 0: bad, 255: maximum quality
	Quality uint8
	// Ground distance. Positive value: distance known. Negative value: Unknown distance
	GroundDistance float32
	// Flow rate about X axis
	FlowRateX float32 `mavext:"true"`
	// Flow rate about Y axis
	FlowRateY float32 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageOpticalFlow) GetID() uint32 {
	return 100
}
