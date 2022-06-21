//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

// A forwarded CAN frame as requested by MAV_CMD_CAN_FORWARD.
type MessageCanFrame struct {
	// System ID.
	TargetSystem uint8
	// Component ID.
	TargetComponent uint8
	// Bus number
	Bus uint8
	// Frame length
	Len uint8
	// Frame ID
	Id uint32
	// Frame data
	Data [8]uint8
}

// GetID implements the message.Message interface.
func (*MessageCanFrame) GetID() uint32 {
	return 386
}
