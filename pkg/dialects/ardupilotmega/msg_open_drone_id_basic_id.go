//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ardupilotmega

// Data for filling the OpenDroneID Basic ID message. This and the below messages are primarily meant for feeding data to/from an OpenDroneID implementation. E.g. https://github.com/opendroneid/opendroneid-core-c. These messages are compatible with the ASTM Remote ID standard at https://www.astm.org/Standards/F3411.htm and the ASD-STAN Direct Remote ID standard. The usage of these messages is documented at https://mavlink.io/en/services/opendroneid.html.
type MessageOpenDroneIdBasicId struct {
	// System ID (0 for broadcast).
	TargetSystem uint8
	// Component ID (0 for broadcast).
	TargetComponent uint8
	// Only used for drone ID data received from other UAs. See detailed description at https://mavlink.io/en/services/opendroneid.html.
	IdOrMac [20]uint8
	// Indicates the format for the uas_id field of this message.
	IdType MAV_ODID_ID_TYPE `mavenum:"uint8"`
	// Indicates the type of UA (Unmanned Aircraft).
	UaType MAV_ODID_UA_TYPE `mavenum:"uint8"`
	// UAS (Unmanned Aircraft System) ID following the format specified by id_type. Shall be filled with nulls in the unused portion of the field.
	UasId [20]uint8
}

// GetID implements the msg.Message interface.
func (*MessageOpenDroneIdBasicId) GetID() uint32 {
	return 12900
}
