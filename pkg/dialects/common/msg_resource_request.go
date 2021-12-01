//autogenerated:yes
//nolint:golint,misspell,govet,lll
package common

// The autopilot is requesting a resource (file, binary, other type of data)
type MessageResourceRequest struct {
	// Request ID. This ID should be re-used when sending back URI contents
	RequestId uint8
	// The type of requested URI. 0 = a file via URL. 1 = a UAVCAN binary
	UriType uint8
	// The requested unique resource identifier (URI). It is not necessarily a straight domain name (depends on the URI type enum)
	Uri [120]uint8
	// The way the autopilot wants to receive the URI. 0 = MAVLink FTP. 1 = binary stream.
	TransferType uint8
	// The storage path the autopilot wants the URI to be stored in. Will only be valid if the transfer_type has a storage associated (e.g. MAVLink FTP).
	Storage [120]uint8
}

// GetID implements the msg.Message interface.
func (*MessageResourceRequest) GetID() uint32 {
	return 142
}
