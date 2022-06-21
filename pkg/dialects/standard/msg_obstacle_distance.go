//autogenerated:yes
//nolint:revive,misspell,govet,lll
package standard

// Obstacle distances in front of the sensor, starting from the left in increment degrees to the right
type MessageObstacleDistance struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Class id of the distance sensor type.
	SensorType MAV_DISTANCE_SENSOR `mavenum:"uint8"`
	// Distance of obstacles around the vehicle with index 0 corresponding to north + angle_offset, unless otherwise specified in the frame. A value of 0 is valid and means that the obstacle is practically touching the sensor. A value of max_distance +1 means no obstacle is present. A value of UINT16_MAX for unknown/not used. In a array element, one unit corresponds to 1cm.
	Distances [72]uint16
	// Angular width in degrees of each array element. Increment direction is clockwise. This field is ignored if increment_f is non-zero.
	Increment uint8
	// Minimum distance the sensor can measure.
	MinDistance uint16
	// Maximum distance the sensor can measure.
	MaxDistance uint16
	// Angular width in degrees of each array element as a float. If non-zero then this value is used instead of the uint8_t increment field. Positive is clockwise direction, negative is counter-clockwise.
	IncrementF float32 `mavext:"true"`
	// Relative angle offset of the 0-index element in the distances array. Value of 0 corresponds to forward. Positive is clockwise direction, negative is counter-clockwise.
	AngleOffset float32 `mavext:"true"`
	// Coordinate frame of reference for the yaw rotation and offset of the sensor data. Defaults to MAV_FRAME_GLOBAL, which is north aligned. For body-mounted sensors use MAV_FRAME_BODY_FRD, which is vehicle front aligned.
	Frame MAV_FRAME `mavenum:"uint8" mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageObstacleDistance) GetID() uint32 {
	return 330
}
