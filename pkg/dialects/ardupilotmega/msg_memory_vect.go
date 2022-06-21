//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ardupilotmega

// Send raw controller memory. The use of this message is discouraged for normal packets, but a quite efficient way for testing new messages and getting experimental debug output.
type MessageMemoryVect struct {
	// Starting address of the debug variables
	Address uint16
	// Version code of the type variable. 0=unknown, type ignored and assumed int16_t. 1=as below
	Ver uint8
	// Type code of the memory variables. for ver = 1: 0=16 x int16_t, 1=16 x uint16_t, 2=16 x Q15, 3=16 x 1Q14
	Type uint8
	// Memory contents at specified address
	Value [32]int8
}

// GetID implements the message.Message interface.
func (*MessageMemoryVect) GetID() uint32 {
	return 249
}
