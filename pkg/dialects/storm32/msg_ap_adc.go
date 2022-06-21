//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Raw ADC output.
type MessageApAdc struct {
	// ADC output 1.
	Adc1 uint16
	// ADC output 2.
	Adc2 uint16
	// ADC output 3.
	Adc3 uint16
	// ADC output 4.
	Adc4 uint16
	// ADC output 5.
	Adc5 uint16
	// ADC output 6.
	Adc6 uint16
}

// GetID implements the message.Message interface.
func (*MessageApAdc) GetID() uint32 {
	return 153
}
