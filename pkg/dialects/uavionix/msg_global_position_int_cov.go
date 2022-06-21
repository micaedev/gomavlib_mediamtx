//autogenerated:yes
//nolint:revive,misspell,govet,lll
package uavionix

// The filtered global position (e.g. fused GPS and accelerometers). The position is in GPS-frame (right-handed, Z-up). It  is designed as scaled integer message since the resolution of float is not sufficient. NOTE: This message is intended for onboard networks / companion computers and higher-bandwidth links and optimized for accuracy and completeness. Please use the GLOBAL_POSITION_INT message for a minimal subset.
type MessageGlobalPositionIntCov struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Class id of the estimator this estimate originated from.
	EstimatorType MAV_ESTIMATOR_TYPE `mavenum:"uint8"`
	// Latitude
	Lat int32
	// Longitude
	Lon int32
	// Altitude in meters above MSL
	Alt int32
	// Altitude above ground
	RelativeAlt int32
	// Ground X Speed (Latitude)
	Vx float32
	// Ground Y Speed (Longitude)
	Vy float32
	// Ground Z Speed (Altitude)
	Vz float32
	// Row-major representation of a 6x6 position and velocity 6x6 cross-covariance matrix (states: lat, lon, alt, vx, vy, vz; first six entries are the first ROW, next six entries are the second row, etc.). If unknown, assign NaN value to first element in the array.
	Covariance [36]float32
}

// GetID implements the message.Message interface.
func (*MessageGlobalPositionIntCov) GetID() uint32 {
	return 63
}
