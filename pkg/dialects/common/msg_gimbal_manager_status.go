//autogenerated:yes
//nolint:revive,misspell,govet,lll
package common

// Current status about a high level gimbal manager. This message should be broadcast at a low regular rate (e.g. 5Hz).
type MessageGimbalManagerStatus struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// High level gimbal manager flags currently applied.
	Flags GIMBAL_MANAGER_FLAGS `mavenum:"uint32"`
	// Gimbal device ID that this gimbal manager is responsible for.
	GimbalDeviceId uint8
	// System ID of MAVLink component with primary control, 0 for none.
	PrimaryControlSysid uint8
	// Component ID of MAVLink component with primary control, 0 for none.
	PrimaryControlCompid uint8
	// System ID of MAVLink component with secondary control, 0 for none.
	SecondaryControlSysid uint8
	// Component ID of MAVLink component with secondary control, 0 for none.
	SecondaryControlCompid uint8
}

// GetID implements the message.Message interface.
func (*MessageGimbalManagerStatus) GetID() uint32 {
	return 281
}
