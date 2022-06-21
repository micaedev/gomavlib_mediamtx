//autogenerated:yes
//nolint:revive,misspell,govet,lll
package uavionix

// An OpenDroneID message pack is a container for multiple encoded OpenDroneID messages (i.e. not in the format given for the above message descriptions but after encoding into the compressed OpenDroneID byte format). Used e.g. when transmitting on Bluetooth 5.0 Long Range/Extended Advertising or on WiFi Neighbor Aware Networking or on WiFi Beacon.
type MessageOpenDroneIdMessagePack struct {
	// System ID (0 for broadcast).
	TargetSystem uint8
	// Component ID (0 for broadcast).
	TargetComponent uint8
	// Only used for drone ID data received from other UAs. See detailed description at https://mavlink.io/en/services/opendroneid.html.
	IdOrMac [20]uint8
	// This field must currently always be equal to 25 (bytes), since all encoded OpenDroneID messages are specificed to have this length.
	SingleMessageSize uint8
	// Number of encoded messages in the pack (not the number of bytes). Allowed range is 1 - 9.
	MsgPackSize uint8
	// Concatenation of encoded OpenDroneID messages. Shall be filled with nulls in the unused portion of the field.
	Messages [225]uint8
}

// GetID implements the message.Message interface.
func (*MessageOpenDroneIdMessagePack) GetID() uint32 {
	return 12915
}
