//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ualberta

// The smoothed, monotonic system state used to feed the control loops of the system.
type MessageControlSystemState struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// X acceleration in body frame
	XAcc float32
	// Y acceleration in body frame
	YAcc float32
	// Z acceleration in body frame
	ZAcc float32
	// X velocity in body frame
	XVel float32
	// Y velocity in body frame
	YVel float32
	// Z velocity in body frame
	ZVel float32
	// X position in local frame
	XPos float32
	// Y position in local frame
	YPos float32
	// Z position in local frame
	ZPos float32
	// Airspeed, set to -1 if unknown
	Airspeed float32
	// Variance of body velocity estimate
	VelVariance [3]float32
	// Variance in local position
	PosVariance [3]float32
	// The attitude, represented as Quaternion
	Q [4]float32
	// Angular rate in roll axis
	RollRate float32
	// Angular rate in pitch axis
	PitchRate float32
	// Angular rate in yaw axis
	YawRate float32
}

// GetID implements the msg.Message interface.
func (*MessageControlSystemState) GetID() uint32 {
	return 146
}
