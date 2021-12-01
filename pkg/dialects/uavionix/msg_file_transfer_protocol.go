//autogenerated:yes
//nolint:golint,misspell,govet,lll
package uavionix

// File transfer message
type MessageFileTransferProtocol struct {
	// Network ID (0 for broadcast)
	TargetNetwork uint8
	// System ID (0 for broadcast)
	TargetSystem uint8
	// Component ID (0 for broadcast)
	TargetComponent uint8
	// Variable length payload. The length is defined by the remaining message length when subtracting the header and other fields.  The entire content of this block is opaque unless you understand any the encoding message_type.  The particular encoding used can be extension specific and might not always be documented as part of the mavlink specification.
	Payload [251]uint8
}

// GetID implements the msg.Message interface.
func (*MessageFileTransferProtocol) GetID() uint32 {
	return 110
}
