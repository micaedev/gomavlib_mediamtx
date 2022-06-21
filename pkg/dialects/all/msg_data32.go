//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Data packet, size 32.
type MessageData32 struct {
	// Data type.
	Type uint8
	// Data length.
	Len uint8
	// Raw data.
	Data [32]uint8
}

// GetID implements the message.Message interface.
func (*MessageData32) GetID() uint32 {
	return 170
}
