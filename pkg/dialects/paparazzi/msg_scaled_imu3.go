//autogenerated:yes
//nolint:golint,misspell,govet,lll
package paparazzi

// The RAW IMU readings for 3rd 9DOF sensor setup. This message should contain the scaled values to the described units
type MessageScaledImu3 struct {
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

// GetID implements the msg.Message interface.
func (*MessageScaledImu3) GetID() uint32 {
	return 129
}
