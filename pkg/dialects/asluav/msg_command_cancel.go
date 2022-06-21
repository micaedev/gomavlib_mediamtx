//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

// Cancel a long running command. The target system should respond with a COMMAND_ACK to the original command with result=MAV_RESULT_CANCELLED if the long running process was cancelled. If it has already completed, the cancel action can be ignored. The cancel action can be retried until some sort of acknowledgement to the original command has been received. The command microservice is documented at https://mavlink.io/en/services/command.html
type MessageCommandCancel struct {
	// System executing long running command. Should not be broadcast (0).
	TargetSystem uint8
	// Component executing long running command.
	TargetComponent uint8
	// Command ID (of command to cancel).
	Command MAV_CMD `mavenum:"uint16"`
}

// GetID implements the message.Message interface.
func (*MessageCommandCancel) GetID() uint32 {
	return 80
}
