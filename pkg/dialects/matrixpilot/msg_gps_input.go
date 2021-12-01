//autogenerated:yes
//nolint:golint,misspell,govet,lll
package matrixpilot

// GPS sensor input message.  This is a raw sensor value sent by the GPS. This is NOT the global position estimate of the system.
type MessageGpsInput struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// ID of the GPS for multiple GPS inputs
	GpsId uint8
	// Bitmap indicating which GPS input flags fields to ignore.  All other fields must be provided.
	IgnoreFlags GPS_INPUT_IGNORE_FLAGS `mavenum:"uint16"`
	// GPS time (from start of GPS week)
	TimeWeekMs uint32
	// GPS week number
	TimeWeek uint16
	// 0-1: no fix, 2: 2D fix, 3: 3D fix. 4: 3D with DGPS. 5: 3D with RTK
	FixType uint8
	// Latitude (WGS84)
	Lat int32
	// Longitude (WGS84)
	Lon int32
	// Altitude (MSL). Positive for up.
	Alt float32
	// GPS HDOP horizontal dilution of position (unitless). If unknown, set to: UINT16_MAX
	Hdop float32
	// GPS VDOP vertical dilution of position (unitless). If unknown, set to: UINT16_MAX
	Vdop float32
	// GPS velocity in north direction in earth-fixed NED frame
	Vn float32
	// GPS velocity in east direction in earth-fixed NED frame
	Ve float32
	// GPS velocity in down direction in earth-fixed NED frame
	Vd float32
	// GPS speed accuracy
	SpeedAccuracy float32
	// GPS horizontal accuracy
	HorizAccuracy float32
	// GPS vertical accuracy
	VertAccuracy float32
	// Number of satellites visible.
	SatellitesVisible uint8
	// Yaw of vehicle relative to Earth's North, zero means not available, use 36000 for north
	Yaw uint16 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageGpsInput) GetID() uint32 {
	return 232
}
