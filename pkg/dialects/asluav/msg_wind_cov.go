//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

// Wind estimate from vehicle. Note that despite the name, this message does not actually contain any covariances but instead variability and accuracy fields in terms of standard deviation (1-STD).
type MessageWindCov struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Wind in North (NED) direction (NAN if unknown)
	WindX float32
	// Wind in East (NED) direction (NAN if unknown)
	WindY float32
	// Wind in down (NED) direction (NAN if unknown)
	WindZ float32
	// Variability of wind in XY, 1-STD estimated from a 1 Hz lowpassed wind estimate (NAN if unknown)
	VarHoriz float32
	// Variability of wind in Z, 1-STD estimated from a 1 Hz lowpassed wind estimate (NAN if unknown)
	VarVert float32
	// Altitude (MSL) that this measurement was taken at (NAN if unknown)
	WindAlt float32
	// Horizontal speed 1-STD accuracy (0 if unknown)
	HorizAccuracy float32
	// Vertical speed 1-STD accuracy (0 if unknown)
	VertAccuracy float32
}

// GetID implements the message.Message interface.
func (*MessageWindCov) GetID() uint32 {
	return 231
}
