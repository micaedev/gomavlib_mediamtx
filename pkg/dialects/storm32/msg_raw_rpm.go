//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// RPM sensor data message.
type MessageRawRpm struct {
	// Index of this RPM sensor (0-indexed)
	Index uint8
	// Indicated rate
	Frequency float32
}

// GetID implements the message.Message interface.
func (*MessageRawRpm) GetID() uint32 {
	return 339
}
