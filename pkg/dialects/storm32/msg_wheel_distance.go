//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Cumulative distance traveled for each reported wheel.
type MessageWheelDistance struct {
	// Timestamp (synced to UNIX time or since system boot).
	TimeUsec uint64
	// Number of wheels reported.
	Count uint8
	// Distance reported by individual wheel encoders. Forward rotations increase values, reverse rotations decrease them. Not all wheels will necessarily have wheel encoders; the mapping of encoders to wheel positions must be agreed/understood by the endpoints.
	Distance [16]float64
}

// GetID implements the message.Message interface.
func (*MessageWheelDistance) GetID() uint32 {
	return 9000
}
