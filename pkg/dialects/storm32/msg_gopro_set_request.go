//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Request to set a GOPRO_COMMAND with a desired.
type MessageGoproSetRequest struct {
	// System ID.
	TargetSystem uint8
	// Component ID.
	TargetComponent uint8
	// Command ID.
	CmdId GOPRO_COMMAND `mavenum:"uint8"`
	// Value.
	Value [4]uint8
}

// GetID implements the message.Message interface.
func (*MessageGoproSetRequest) GetID() uint32 {
	return 218
}
