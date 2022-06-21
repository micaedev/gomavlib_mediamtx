//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ardupilotmega

// Response from a GOPRO_COMMAND set request.
type MessageGoproSetResponse struct {
	// Command ID.
	CmdId GOPRO_COMMAND `mavenum:"uint8"`
	// Status.
	Status GOPRO_REQUEST_STATUS `mavenum:"uint8"`
}

// GetID implements the message.Message interface.
func (*MessageGoproSetResponse) GetID() uint32 {
	return 219
}
