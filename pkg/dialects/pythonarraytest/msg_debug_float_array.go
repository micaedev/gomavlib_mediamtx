//autogenerated:yes
//nolint:revive,misspell,govet,lll
package pythonarraytest

// Large debug/prototyping array. The message uses the maximum available payload for data. The array_id and name fields are used to discriminate between messages in code and in user interfaces (respectively). Do not use in production code.
type MessageDebugFloatArray struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Name, for human-friendly display in a Ground Control Station
	Name string `mavlen:"10"`
	// Unique ID used to discriminate between arrays
	ArrayId uint16
	// data
	Data [58]float32 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageDebugFloatArray) GetID() uint32 {
	return 350
}
