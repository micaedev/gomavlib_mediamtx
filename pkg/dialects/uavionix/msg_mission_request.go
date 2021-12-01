//autogenerated:yes
//nolint:golint,misspell,govet,lll
package uavionix

// Request the information of the mission item with the sequence number seq. The response of the system to this message should be a MISSION_ITEM message. https://mavlink.io/en/services/mission.html
type MessageMissionRequest struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Sequence
	Seq uint16
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageMissionRequest) GetID() uint32 {
	return 40
}
