//autogenerated:yes
//nolint:golint,misspell,govet,lll
package uavionix

// Response to a REQUEST_EVENT in case of an error (e.g. the event is not available anymore).
type MessageResponseEventError struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Sequence number.
	Sequence uint16
	// Oldest Sequence number that is still available after the sequence set in REQUEST_EVENT.
	SequenceOldestAvailable uint16
	// Error reason.
	Reason MAV_EVENT_ERROR_REASON `mavenum:"uint8"`
}

// GetID implements the msg.Message interface.
func (*MessageResponseEventError) GetID() uint32 {
	return 413
}
