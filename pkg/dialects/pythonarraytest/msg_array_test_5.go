//autogenerated:yes
//nolint:revive,misspell,govet,lll
package pythonarraytest

// Array test #5.
type MessageArrayTest_5 struct {
	// Value array
	C1 string `mavlen:"5"`
	// Value array
	C2 string `mavlen:"5"`
}

// GetID implements the message.Message interface.
func (*MessageArrayTest_5) GetID() uint32 {
	return 17155
}
