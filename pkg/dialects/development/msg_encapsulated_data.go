//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

// Data packet for images sent using the Image Transmission Protocol: https://mavlink.io/en/services/image_transmission.html.
type MessageEncapsulatedData struct {
	// sequence number (starting with 0 on every transmission)
	Seqnr uint16
	// image data bytes
	Data [253]uint8
}

// GetID implements the message.Message interface.
func (*MessageEncapsulatedData) GetID() uint32 {
	return 131
}
