//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

// Information about flight since last arming.
// This can be requested using MAV_CMD_REQUEST_MESSAGE.
type MessageFlightInformation struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Timestamp at arming (time since UNIX epoch) in UTC, 0 for unknown
	ArmingTimeUtc uint64
	// Timestamp at takeoff (time since UNIX epoch) in UTC, 0 for unknown
	TakeoffTimeUtc uint64
	// Universally unique identifier (UUID) of flight, should correspond to name of log files
	FlightUuid uint64
}

// GetID implements the message.Message interface.
func (*MessageFlightInformation) GetID() uint32 {
	return 264
}
