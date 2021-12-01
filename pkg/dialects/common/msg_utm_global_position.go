//autogenerated:yes
//nolint:golint,misspell,govet,lll
package common

// The global position resulting from GPS and sensor fusion.
type MessageUtmGlobalPosition struct {
	// Time of applicability of position (microseconds since UNIX epoch).
	Time uint64
	// Unique UAS ID.
	UasId [18]uint8
	// Latitude (WGS84)
	Lat int32
	// Longitude (WGS84)
	Lon int32
	// Altitude (WGS84)
	Alt int32
	// Altitude above ground
	RelativeAlt int32
	// Ground X speed (latitude, positive north)
	Vx int16
	// Ground Y speed (longitude, positive east)
	Vy int16
	// Ground Z speed (altitude, positive down)
	Vz int16
	// Horizontal position uncertainty (standard deviation)
	HAcc uint16
	// Altitude uncertainty (standard deviation)
	VAcc uint16
	// Speed uncertainty (standard deviation)
	VelAcc uint16
	// Next waypoint, latitude (WGS84)
	NextLat int32
	// Next waypoint, longitude (WGS84)
	NextLon int32
	// Next waypoint, altitude (WGS84)
	NextAlt int32
	// Time until next update. Set to 0 if unknown or in data driven mode.
	UpdateRate uint16
	// Flight state
	FlightState UTM_FLIGHT_STATE `mavenum:"uint8"`
	// Bitwise OR combination of the data available flags.
	Flags UTM_DATA_AVAIL_FLAGS `mavenum:"uint8"`
}

// GetID implements the msg.Message interface.
func (*MessageUtmGlobalPosition) GetID() uint32 {
	return 340
}
