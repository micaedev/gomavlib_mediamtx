//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

// Acknowledgment message during waypoint handling. The type field states if this message is a positive ack (type=0) or if an error happened (type=non-zero).
type MessageMissionAck struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Mission result.
	Type MAV_MISSION_RESULT `mavenum:"uint8"`
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageMissionAck) GetID() uint32 {
	return 47
}
