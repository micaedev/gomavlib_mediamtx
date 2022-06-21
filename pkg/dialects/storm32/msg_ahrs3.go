//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Status of third AHRS filter if available. This is for ANU research group (Ali and Sean).
type MessageAhrs3 struct {
	// Roll angle.
	Roll float32
	// Pitch angle.
	Pitch float32
	// Yaw angle.
	Yaw float32
	// Altitude (MSL).
	Altitude float32
	// Latitude.
	Lat int32
	// Longitude.
	Lng int32
	// Test variable1.
	V1 float32
	// Test variable2.
	V2 float32
	// Test variable3.
	V3 float32
	// Test variable4.
	V4 float32
}

// GetID implements the message.Message interface.
func (*MessageAhrs3) GetID() uint32 {
	return 182
}
