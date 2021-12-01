//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ualberta

// Sets the GPS co-ordinates of the vehicle local origin (0,0,0) position. Vehicle should emit GPS_GLOBAL_ORIGIN irrespective of whether the origin is changed. This enables transform between the local coordinate frame and the global (GPS) coordinate frame, which may be necessary when (for example) indoor and outdoor settings are connected and the MAV should move from in- to outdoor.
type MessageSetGpsGlobalOrigin struct {
	// System ID
	TargetSystem uint8
	// Latitude (WGS84)
	Latitude int32
	// Longitude (WGS84)
	Longitude int32
	// Altitude (MSL). Positive for up.
	Altitude int32
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageSetGpsGlobalOrigin) GetID() uint32 {
	return 48
}
