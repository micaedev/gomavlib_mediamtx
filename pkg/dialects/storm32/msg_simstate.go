//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Status of simulation environment, if used.
type MessageSimstate struct {
	// Roll angle.
	Roll float32
	// Pitch angle.
	Pitch float32
	// Yaw angle.
	Yaw float32
	// X acceleration.
	Xacc float32
	// Y acceleration.
	Yacc float32
	// Z acceleration.
	Zacc float32
	// Angular speed around X axis.
	Xgyro float32
	// Angular speed around Y axis.
	Ygyro float32
	// Angular speed around Z axis.
	Zgyro float32
	// Latitude.
	Lat int32
	// Longitude.
	Lng int32
}

// GetID implements the message.Message interface.
func (*MessageSimstate) GetID() uint32 {
	return 164
}
