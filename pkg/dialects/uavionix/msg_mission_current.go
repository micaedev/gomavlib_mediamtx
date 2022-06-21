//autogenerated:yes
//nolint:revive,misspell,govet,lll
package uavionix

// Message that announces the sequence number of the current active mission item. The MAV will fly towards this mission item.
type MessageMissionCurrent struct {
	// Sequence
	Seq uint16
}

// GetID implements the message.Message interface.
func (*MessageMissionCurrent) GetID() uint32 {
	return 42
}
