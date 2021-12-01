//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ardupilotmega

// Request a partial list of mission items from the system/component. https://mavlink.io/en/services/mission.html. If start and end index are the same, just send one waypoint.
type MessageMissionRequestPartialList struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Start index
	StartIndex int16
	// End index, -1 by default (-1: send list to end). Else a valid index of the list
	EndIndex int16
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageMissionRequestPartialList) GetID() uint32 {
	return 37
}
