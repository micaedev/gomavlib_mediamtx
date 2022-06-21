//autogenerated:yes
//nolint:revive,misspell,govet,lll
package uavionix

// Dynamic data used to generate ADS-B out transponder data (send at 5Hz)
type MessageUavionixAdsbOutDynamic struct {
	// UTC time in seconds since GPS epoch (Jan 6, 1980). If unknown set to UINT32_MAX
	Utctime uint32 `mavname:"utcTime"`
	// Latitude WGS84 (deg * 1E7). If unknown set to INT32_MAX
	Gpslat int32 `mavname:"gpsLat"`
	// Longitude WGS84 (deg * 1E7). If unknown set to INT32_MAX
	Gpslon int32 `mavname:"gpsLon"`
	// Altitude (WGS84). UP +ve. If unknown set to INT32_MAX
	Gpsalt int32 `mavname:"gpsAlt"`
	// 0-1: no fix, 2: 2D fix, 3: 3D fix, 4: DGPS, 5: RTK
	Gpsfix UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX `mavenum:"uint8" mavname:"gpsFix"`
	// Number of satellites visible. If unknown set to UINT8_MAX
	Numsats uint8 `mavname:"numSats"`
	// Barometric pressure altitude (MSL) relative to a standard atmosphere of 1013.2 mBar and NOT bar corrected altitude (m * 1E-3). (up +ve). If unknown set to INT32_MAX
	Baroaltmsl int32 `mavname:"baroAltMSL"`
	// Horizontal accuracy in mm (m * 1E-3). If unknown set to UINT32_MAX
	Accuracyhor uint32 `mavname:"accuracyHor"`
	// Vertical accuracy in cm. If unknown set to UINT16_MAX
	Accuracyvert uint16 `mavname:"accuracyVert"`
	// Velocity accuracy in mm/s (m * 1E-3). If unknown set to UINT16_MAX
	Accuracyvel uint16 `mavname:"accuracyVel"`
	// GPS vertical speed in cm/s. If unknown set to INT16_MAX
	Velvert int16 `mavname:"velVert"`
	// North-South velocity over ground in cm/s North +ve. If unknown set to INT16_MAX
	Velns int16 `mavname:"velNS"`
	// East-West velocity over ground in cm/s East +ve. If unknown set to INT16_MAX
	Velew int16 `mavname:"VelEW"`
	// Emergency status
	Emergencystatus UAVIONIX_ADSB_EMERGENCY_STATUS `mavenum:"uint8" mavname:"emergencyStatus"`
	// ADS-B transponder dynamic input state flags
	State UAVIONIX_ADSB_OUT_DYNAMIC_STATE `mavenum:"uint16"`
	// Mode A code (typically 1200 [0x04B0] for VFR)
	Squawk uint16
}

// GetID implements the message.Message interface.
func (*MessageUavionixAdsbOutDynamic) GetID() uint32 {
	return 10002
}
