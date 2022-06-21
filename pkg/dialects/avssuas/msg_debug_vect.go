//autogenerated:yes
//nolint:revive,misspell,govet,lll
package avssuas

// To debug something using a named 3D vector.
type MessageDebugVect struct {
	// Name
	Name string `mavlen:"10"`
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// x
	X float32
	// y
	Y float32
	// z
	Z float32
}

// GetID implements the message.Message interface.
func (*MessageDebugVect) GetID() uint32 {
	return 250
}
