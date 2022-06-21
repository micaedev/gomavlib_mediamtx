//autogenerated:yes
//nolint:revive,misspell,govet,lll
package pythonarraytest

// Send a debug value. The index is used to discriminate between values. These values show up in the plot of QGroundControl as DEBUG N.
type MessageDebug struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// index of debug variable
	Ind uint8
	// DEBUG value
	Value float32
}

// GetID implements the message.Message interface.
func (*MessageDebug) GetID() uint32 {
	return 254
}
