//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Read a configured an OSD parameter slot.
type MessageOsdParamShowConfig struct {
	// System ID.
	TargetSystem uint8
	// Component ID.
	TargetComponent uint8
	// Request ID - copied to reply.
	RequestId uint32
	// OSD parameter screen index.
	OsdScreen uint8
	// OSD parameter display index.
	OsdIndex uint8
}

// GetID implements the message.Message interface.
func (*MessageOsdParamShowConfig) GetID() uint32 {
	return 11035
}
