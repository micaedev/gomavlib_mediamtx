//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ualberta

// Distance sensor information for an onboard rangefinder.
type MessageDistanceSensor struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Minimum distance the sensor can measure
	MinDistance uint16
	// Maximum distance the sensor can measure
	MaxDistance uint16
	// Current distance reading
	CurrentDistance uint16
	// Type of distance sensor.
	Type MAV_DISTANCE_SENSOR `mavenum:"uint8"`
	// Onboard ID of the sensor
	Id uint8
	// Direction the sensor faces. downward-facing: ROTATION_PITCH_270, upward-facing: ROTATION_PITCH_90, backward-facing: ROTATION_PITCH_180, forward-facing: ROTATION_NONE, left-facing: ROTATION_YAW_90, right-facing: ROTATION_YAW_270
	Orientation MAV_SENSOR_ORIENTATION `mavenum:"uint8"`
	// Measurement variance. Max standard deviation is 6cm. UINT8_MAX if unknown.
	Covariance uint8
	// Horizontal Field of View (angle) where the distance measurement is valid and the field of view is known. Otherwise this is set to 0.
	HorizontalFov float32 `mavext:"true"`
	// Vertical Field of View (angle) where the distance measurement is valid and the field of view is known. Otherwise this is set to 0.
	VerticalFov float32 `mavext:"true"`
	// Quaternion of the sensor orientation in vehicle body frame (w, x, y, z order, zero-rotation is 1, 0, 0, 0). Zero-rotation is along the vehicle body x-axis. This field is required if the orientation is set to MAV_SENSOR_ROTATION_CUSTOM. Set it to 0 if invalid."
	Quaternion [4]float32 `mavext:"true"`
	// Signal quality of the sensor. Specific to each sensor type, representing the relation of the signal strength with the target reflectivity, distance, size or aspect, but normalised as a percentage. 0 = unknown/unset signal quality, 1 = invalid signal, 100 = perfect signal.
	SignalQuality uint8 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageDistanceSensor) GetID() uint32 {
	return 132
}
