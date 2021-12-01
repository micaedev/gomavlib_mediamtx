//autogenerated:yes
//nolint:golint,misspell,govet,lll
package pythonarraytest

// Array test #6.
type MessageArrayTest_6 struct {
	// Stub field
	V1 uint8
	// Stub field
	V2 uint16
	// Stub field
	V3 uint32
	// Value array
	ArU32 [2]uint32
	// Value array
	ArI32 [2]int32
	// Value array
	ArU16 [2]uint16
	// Value array
	ArI16 [2]int16
	// Value array
	ArU8 [2]uint8
	// Value array
	ArI8 [2]int8
	// Value array
	ArC string `mavlen:"32"`
	// Value array
	ArD [2]float64
	// Value array
	ArF [2]float32
}

// GetID implements the msg.Message interface.
func (*MessageArrayTest_6) GetID() uint32 {
	return 17156
}
