//autogenerated:yes
//nolint:golint,misspell,govet,lll
package matrixpilot

// Backwards compatible MAVLink version of SERIAL_UDB_EXTRA - F2: Format Part A
type MessageSerialUdbExtraF2A struct {
	// Serial UDB Extra Time
	SueTime uint32
	// Serial UDB Extra Status
	SueStatus uint8
	// Serial UDB Extra Latitude
	SueLatitude int32
	// Serial UDB Extra Longitude
	SueLongitude int32
	// Serial UDB Extra Altitude
	SueAltitude int32
	// Serial UDB Extra Waypoint Index
	SueWaypointIndex uint16
	// Serial UDB Extra Rmat 0
	SueRmat0 int16
	// Serial UDB Extra Rmat 1
	SueRmat1 int16
	// Serial UDB Extra Rmat 2
	SueRmat2 int16
	// Serial UDB Extra Rmat 3
	SueRmat3 int16
	// Serial UDB Extra Rmat 4
	SueRmat4 int16
	// Serial UDB Extra Rmat 5
	SueRmat5 int16
	// Serial UDB Extra Rmat 6
	SueRmat6 int16
	// Serial UDB Extra Rmat 7
	SueRmat7 int16
	// Serial UDB Extra Rmat 8
	SueRmat8 int16
	// Serial UDB Extra GPS Course Over Ground
	SueCog uint16
	// Serial UDB Extra Speed Over Ground
	SueSog int16
	// Serial UDB Extra CPU Load
	SueCpuLoad uint16
	// Serial UDB Extra 3D IMU Air Speed
	SueAirSpeed_3dimu uint16 `mavname:"sue_air_speed_3DIMU"`
	// Serial UDB Extra Estimated Wind 0
	SueEstimatedWind_0 int16
	// Serial UDB Extra Estimated Wind 1
	SueEstimatedWind_1 int16
	// Serial UDB Extra Estimated Wind 2
	SueEstimatedWind_2 int16
	// Serial UDB Extra Magnetic Field Earth 0
	SueMagfieldearth0 int16 `mavname:"sue_magFieldEarth0"`
	// Serial UDB Extra Magnetic Field Earth 1
	SueMagfieldearth1 int16 `mavname:"sue_magFieldEarth1"`
	// Serial UDB Extra Magnetic Field Earth 2
	SueMagfieldearth2 int16 `mavname:"sue_magFieldEarth2"`
	// Serial UDB Extra Number of Sattelites in View
	SueSvs int16
	// Serial UDB Extra GPS Horizontal Dilution of Precision
	SueHdop int16
}

// GetID implements the msg.Message interface.
func (*MessageSerialUdbExtraF2A) GetID() uint32 {
	return 170
}
