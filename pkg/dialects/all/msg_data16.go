//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Data packet, size 16.
type MessageData16 struct {
	// Data type.
	Type uint8
	// Data length.
	Len uint8
	// Raw data.
	Data [16]uint8
}

// GetID implements the message.Message interface.
func (*MessageData16) GetID() uint32 {
	return 169
}
