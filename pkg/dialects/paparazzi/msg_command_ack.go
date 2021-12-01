//autogenerated:yes
//nolint:golint,misspell,govet,lll
package paparazzi

// Report status of a command. Includes feedback whether the command was executed. The command microservice is documented at https://mavlink.io/en/services/command.html
type MessageCommandAck struct {
	// Command ID (of acknowledged command).
	Command MAV_CMD `mavenum:"uint16"`
	// Result of command.
	Result MAV_RESULT `mavenum:"uint8"`
	// Also used as result_param1, it can be set with an enum containing the errors reasons of why the command was denied, or the progress percentage when result is MAV_RESULT_IN_PROGRESS (UINT8_MAX if the progress is unknown).
	Progress uint8 `mavext:"true"`
	// Additional parameter of the result, example: which parameter of MAV_CMD_NAV_WAYPOINT caused it to be denied.
	ResultParam2 int32 `mavext:"true"`
	// System ID of the target recipient. This is the ID of the system that sent the command for which this COMMAND_ACK is an acknowledgement.
	TargetSystem uint8 `mavext:"true"`
	// Component ID of the target recipient. This is the ID of the system that sent the command for which this COMMAND_ACK is an acknowledgement.
	TargetComponent uint8 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageCommandAck) GetID() uint32 {
	return 77
}
