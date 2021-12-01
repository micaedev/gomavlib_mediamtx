//autogenerated:yes
//nolint:golint,misspell,govet,lll
package minimal

// Version and capability of protocol version. This message can be requested with MAV_CMD_REQUEST_MESSAGE and is used as part of the handshaking to establish which MAVLink version should be used on the network. Every node should respond to a request for PROTOCOL_VERSION to enable the handshaking. Library implementers should consider adding this into the default decoding state machine to allow the protocol core to respond directly.
type MessageProtocolVersion struct {
	// Currently active MAVLink version number * 100: v1.0 is 100, v2.0 is 200, etc.
	Version uint16
	// Minimum MAVLink version supported
	MinVersion uint16
	// Maximum MAVLink version supported (set to the same value as version by default)
	MaxVersion uint16
	// The first 8 bytes (not characters printed in hex!) of the git hash.
	SpecVersionHash [8]uint8
	// The first 8 bytes (not characters printed in hex!) of the git hash.
	LibraryVersionHash [8]uint8
}

// GetID implements the msg.Message interface.
func (*MessageProtocolVersion) GetID() uint32 {
	return 300
}
