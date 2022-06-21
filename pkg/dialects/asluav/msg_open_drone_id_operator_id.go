//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

// Data for filling the OpenDroneID Operator ID message, which contains the CAA (Civil Aviation Authority) issued operator ID.
type MessageOpenDroneIdOperatorId struct {
	// System ID (0 for broadcast).
	TargetSystem uint8
	// Component ID (0 for broadcast).
	TargetComponent uint8
	// Only used for drone ID data received from other UAs. See detailed description at https://mavlink.io/en/services/opendroneid.html.
	IdOrMac [20]uint8
	// Indicates the type of the operator_id field.
	OperatorIdType MAV_ODID_OPERATOR_ID_TYPE `mavenum:"uint8"`
	// Text description or numeric value expressed as ASCII characters. Shall be filled with nulls in the unused portion of the field.
	OperatorId string `mavlen:"20"`
}

// GetID implements the message.Message interface.
func (*MessageOpenDroneIdOperatorId) GetID() uint32 {
	return 12905
}
