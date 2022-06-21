//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

// The RAW pressure readings for the typical setup of one absolute pressure and one differential pressure sensor. The sensor values should be the raw, UNSCALED ADC values.
type MessageRawPressure struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Absolute pressure (raw)
	PressAbs int16
	// Differential pressure 1 (raw, 0 if nonexistent)
	PressDiff1 int16
	// Differential pressure 2 (raw, 0 if nonexistent)
	PressDiff2 int16
	// Raw Temperature measurement (raw)
	Temperature int16
}

// GetID implements the message.Message interface.
func (*MessageRawPressure) GetID() uint32 {
	return 28
}
