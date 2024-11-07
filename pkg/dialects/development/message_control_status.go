//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

// Information about GCS in control of this MAV. This should be broadcast at low rate (nominally 1 Hz) and emitted when ownership or takeover status change. Control over MAV is requested using MAV_CMD_REQUEST_OPERATOR_CONTROL.
type MessageControlStatus struct {
	// System ID of GCS MAVLink component in control (0: no GCS in control).
	SysidInControl uint8
	// Control status. For example, whether takeover is allowed, and whether this message instance defines the default controlling GCS for the whole system.
	Flags GCS_CONTROL_STATUS_FLAGS `mavenum:"uint8"`
}

// GetID implements the message.Message interface.
func (*MessageControlStatus) GetID() uint32 {
	return 512
}
