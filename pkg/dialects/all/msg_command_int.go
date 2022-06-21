//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Message encoding a command with parameters as scaled integers. Scaling depends on the actual command value. NaN or INT32_MAX may be used in float/integer params (respectively) to indicate optional/default values (e.g. to use the component's current latitude, yaw rather than a specific value). The command microservice is documented at https://mavlink.io/en/services/command.html
type MessageCommandInt struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// The coordinate system of the COMMAND.
	Frame MAV_FRAME `mavenum:"uint8"`
	// The scheduled action for the mission item.
	Command MAV_CMD `mavenum:"uint16"`
	// Not used.
	Current uint8
	// Not used (set 0).
	Autocontinue uint8
	// PARAM1, see MAV_CMD enum
	Param1 float32
	// PARAM2, see MAV_CMD enum
	Param2 float32
	// PARAM3, see MAV_CMD enum
	Param3 float32
	// PARAM4, see MAV_CMD enum
	Param4 float32
	// PARAM5 / local: x position in meters * 1e4, global: latitude in degrees * 10^7
	X int32
	// PARAM6 / local: y position in meters * 1e4, global: longitude in degrees * 10^7
	Y int32
	// PARAM7 / z position: global: altitude in meters (relative or absolute, depending on frame).
	Z float32
}

// GetID implements the message.Message interface.
func (*MessageCommandInt) GetID() uint32 {
	return 75
}
