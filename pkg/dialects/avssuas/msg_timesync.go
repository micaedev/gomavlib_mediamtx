//autogenerated:yes
//nolint:revive,misspell,govet,lll
package avssuas

// Time synchronization message.
type MessageTimesync struct {
	// Time sync timestamp 1
	Tc1 int64
	// Time sync timestamp 2
	Ts1 int64
}

// GetID implements the message.Message interface.
func (*MessageTimesync) GetID() uint32 {
	return 111
}
