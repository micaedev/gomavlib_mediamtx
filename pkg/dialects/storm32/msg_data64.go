//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Data packet, size 64.
type MessageData64 struct {
	// Data type.
	Type uint8
	// Data length.
	Len uint8
	// Raw data.
	Data [64]uint8
}

// GetID implements the message.Message interface.
func (*MessageData64) GetID() uint32 {
	return 171
}
