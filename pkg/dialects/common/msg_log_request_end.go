//autogenerated:yes
//nolint:golint,misspell,govet,lll
package common

// Stop log transfer and resume normal logging
type MessageLogRequestEnd struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
}

// GetID implements the msg.Message interface.
func (*MessageLogRequestEnd) GetID() uint32 {
	return 122
}
