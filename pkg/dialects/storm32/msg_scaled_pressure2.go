//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Barometer readings for 2nd barometer
type MessageScaledPressure2 struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Absolute pressure
	PressAbs float32
	// Differential pressure
	PressDiff float32
	// Absolute pressure temperature
	Temperature int16
	// Differential pressure temperature (0, if not available). Report values of 0 (or 1) as 1 cdegC.
	TemperaturePressDiff int16 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageScaledPressure2) GetID() uint32 {
	return 137
}
