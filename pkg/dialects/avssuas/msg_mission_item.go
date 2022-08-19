//autogenerated:yes
//nolint:revive,misspell,govet,lll
package avssuas

// Message encoding a mission item. This message is emitted to announce
// the presence of a mission item and to set a mission item on the system. The mission item can be either in x, y, z meters (type: LOCAL) or x:lat, y:lon, z:altitude. Local frame is Z-down, right handed (NED), global frame is Z-up, right handed (ENU). NaN may be used to indicate an optional/default value (e.g. to use the system's current latitude or yaw rather than a specific value). See also https://mavlink.io/en/services/mission.html.
type MessageMissionItem struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Sequence
	Seq uint16
	// The coordinate system of the waypoint.
	Frame MAV_FRAME `mavenum:"uint8"`
	// The scheduled action for the waypoint.
	Command MAV_CMD `mavenum:"uint16"`
	// false:0, true:1
	Current uint8
	// Autocontinue to next waypoint. 0: false, 1: true. Set false to pause mission after the item completes.
	Autocontinue uint8
	// PARAM1, see MAV_CMD enum
	Param1 float32
	// PARAM2, see MAV_CMD enum
	Param2 float32
	// PARAM3, see MAV_CMD enum
	Param3 float32
	// PARAM4, see MAV_CMD enum
	Param4 float32
	// PARAM5 / local: X coordinate, global: latitude
	X float32
	// PARAM6 / local: Y coordinate, global: longitude
	Y float32
	// PARAM7 / local: Z coordinate, global: altitude (relative or absolute, depending on frame).
	Z float32
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageMissionItem) GetID() uint32 {
	return 39
}
