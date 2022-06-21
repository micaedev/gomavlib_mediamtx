//autogenerated:yes
//nolint:revive,misspell,govet,lll
package avssuas

// A forwarded CANFD frame as requested by MAV_CMD_CAN_FORWARD. These are separated from CAN_FRAME as they need different handling (eg. TAO handling)
type MessageCanfdFrame struct {
	// System ID.
	TargetSystem uint8
	// Component ID.
	TargetComponent uint8
	// bus number
	Bus uint8
	// Frame length
	Len uint8
	// Frame ID
	Id uint32
	// Frame data
	Data [64]uint8
}

// GetID implements the message.Message interface.
func (*MessageCanfdFrame) GetID() uint32 {
	return 387
}
