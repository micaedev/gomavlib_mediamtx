//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ualberta

// Complete set of calibration parameters for the radio
type MessageRadioCalibration struct {
	// Aileron setpoints: left, center, right
	Aileron [3]uint16
	// Elevator setpoints: nose down, center, nose up
	Elevator [3]uint16
	// Rudder setpoints: nose left, center, nose right
	Rudder [3]uint16
	// Tail gyro mode/gain setpoints: heading hold, rate mode
	Gyro [2]uint16
	// Pitch curve setpoints (every 25%)
	Pitch [5]uint16
	// Throttle curve setpoints (every 25%)
	Throttle [5]uint16
}

// GetID implements the msg.Message interface.
func (*MessageRadioCalibration) GetID() uint32 {
	return 221
}
