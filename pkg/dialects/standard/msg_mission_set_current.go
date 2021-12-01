//autogenerated:yes
//nolint:golint,misspell,govet,lll
package standard

// Set the mission item with sequence number seq as current item. This means that the MAV will continue to this mission item on the shortest path (not following the mission items in-between).
type MessageMissionSetCurrent struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Sequence
	Seq uint16
}

// GetID implements the msg.Message interface.
func (*MessageMissionSetCurrent) GetID() uint32 {
	return 41
}
