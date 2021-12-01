//autogenerated:yes
//nolint:golint,misspell,govet,lll
package pythonarraytest

// This message is sent to the MAV to write a partial list. If start index == end index, only one item will be transmitted / updated. If the start index is NOT 0 and above the current list size, this request should be REJECTED!
type MessageMissionWritePartialList struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Start index. Must be smaller / equal to the largest index of the current onboard list.
	StartIndex int16
	// End index, equal or greater than start index.
	EndIndex int16
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageMissionWritePartialList) GetID() uint32 {
	return 38
}
