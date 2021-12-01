//autogenerated:yes
//nolint:golint,misspell,govet,lll
package uavionix

// Send a key-value pair as float. The use of this message is discouraged for normal packets, but a quite efficient way for testing new messages and getting experimental debug output.
type MessageNamedValueFloat struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Name of the debug variable
	Name string `mavlen:"10"`
	// Floating point value
	Value float32
}

// GetID implements the msg.Message interface.
func (*MessageNamedValueFloat) GetID() uint32 {
	return 251
}
