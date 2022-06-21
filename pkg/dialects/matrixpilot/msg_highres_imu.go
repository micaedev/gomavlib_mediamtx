//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// The IMU readings in SI units in NED body frame
type MessageHighresImu struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// X acceleration
	Xacc float32
	// Y acceleration
	Yacc float32
	// Z acceleration
	Zacc float32
	// Angular speed around X axis
	Xgyro float32
	// Angular speed around Y axis
	Ygyro float32
	// Angular speed around Z axis
	Zgyro float32
	// X Magnetic field
	Xmag float32
	// Y Magnetic field
	Ymag float32
	// Z Magnetic field
	Zmag float32
	// Absolute pressure
	AbsPressure float32
	// Differential pressure
	DiffPressure float32
	// Altitude calculated from pressure
	PressureAlt float32
	// Temperature
	Temperature float32
	// Bitmap for fields that have updated since last message
	FieldsUpdated HIGHRES_IMU_UPDATED_FLAGS `mavenum:"uint16"`
	// Id. Ids are numbered from 0 and map to IMUs numbered from 1 (e.g. IMU1 will have a message with id=0)
	Id uint8 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageHighresImu) GetID() uint32 {
	return 105
}
