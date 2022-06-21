//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Airspeed auto-calibration.
type MessageAirspeedAutocal struct {
	// GPS velocity north.
	Vx float32
	// GPS velocity east.
	Vy float32
	// GPS velocity down.
	Vz float32
	// Differential pressure.
	DiffPressure float32
	// Estimated to true airspeed ratio.
	Eas2tas float32 `mavname:"EAS2TAS"`
	// Airspeed ratio.
	Ratio float32
	// EKF state x.
	StateX float32
	// EKF state y.
	StateY float32
	// EKF state z.
	StateZ float32
	// EKF Pax.
	Pax float32 `mavname:"Pax"`
	// EKF Pby.
	Pby float32 `mavname:"Pby"`
	// EKF Pcz.
	Pcz float32 `mavname:"Pcz"`
}

// GetID implements the message.Message interface.
func (*MessageAirspeedAutocal) GetID() uint32 {
	return 174
}
