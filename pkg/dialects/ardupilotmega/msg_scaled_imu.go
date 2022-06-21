//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ardupilotmega

// The RAW IMU readings for the usual 9DOF sensor setup. This message should contain the scaled values to the described units
type MessageScaledImu struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// X acceleration
	Xacc int16
	// Y acceleration
	Yacc int16
	// Z acceleration
	Zacc int16
	// Angular speed around X axis
	Xgyro int16
	// Angular speed around Y axis
	Ygyro int16
	// Angular speed around Z axis
	Zgyro int16
	// X Magnetic field
	Xmag int16
	// Y Magnetic field
	Ymag int16
	// Z Magnetic field
	Zmag int16
	// Temperature, 0: IMU does not provide temperature values. If the IMU is at 0C it must send 1 (0.01C).
	Temperature int16 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageScaledImu) GetID() uint32 {
	return 26
}
