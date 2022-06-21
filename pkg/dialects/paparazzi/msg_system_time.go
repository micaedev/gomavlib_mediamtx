//autogenerated:yes
//nolint:revive,misspell,govet,lll
package paparazzi

// The system time is the time of the master clock, typically the computer clock of the main onboard computer.
type MessageSystemTime struct {
	// Timestamp (UNIX epoch time).
	TimeUnixUsec uint64
	// Timestamp (time since system boot).
	TimeBootMs uint32
}

// GetID implements the message.Message interface.
func (*MessageSystemTime) GetID() uint32 {
	return 2
}
