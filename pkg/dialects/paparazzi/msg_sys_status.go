//autogenerated:yes
//nolint:revive,misspell,govet,lll
package paparazzi

// The general system state. If the system is following the MAVLink standard, the system state is mainly defined by three orthogonal states/modes: The system mode, which is either LOCKED (motors shut down and locked), MANUAL (system under RC control), GUIDED (system with autonomous position control, position setpoint controlled manually) or AUTO (system guided by path/waypoint planner). The NAV_MODE defined the current flight state: LIFTOFF (often an open-loop maneuver), LANDING, WAYPOINTS or VECTOR. This represents the internal navigation state machine. The system status shows whether the system is currently active or not and if an emergency occurred. During the CRITICAL and EMERGENCY states the MAV is still considered to be active, but should start emergency procedures autonomously. After a failure occurred it should first move from active to critical to allow manual intervention and then move to emergency after a certain timeout.
type MessageSysStatus struct {
	// Bitmap showing which onboard controllers and sensors are present. Value of 0: not present. Value of 1: present.
	OnboardControlSensorsPresent MAV_SYS_STATUS_SENSOR `mavenum:"uint32"`
	// Bitmap showing which onboard controllers and sensors are enabled:  Value of 0: not enabled. Value of 1: enabled.
	OnboardControlSensorsEnabled MAV_SYS_STATUS_SENSOR `mavenum:"uint32"`
	// Bitmap showing which onboard controllers and sensors have an error (or are operational). Value of 0: error. Value of 1: healthy.
	OnboardControlSensorsHealth MAV_SYS_STATUS_SENSOR `mavenum:"uint32"`
	// Maximum usage in percent of the mainloop time. Values: [0-1000] - should always be below 1000
	Load uint16
	// Battery voltage, UINT16_MAX: Voltage not sent by autopilot
	VoltageBattery uint16
	// Battery current, -1: Current not sent by autopilot
	CurrentBattery int16
	// Battery energy remaining, -1: Battery remaining energy not sent by autopilot
	BatteryRemaining int8
	// Communication drop rate, (UART, I2C, SPI, CAN), dropped packets on all links (packets that were corrupted on reception on the MAV)
	DropRateComm uint16
	// Communication errors (UART, I2C, SPI, CAN), dropped packets on all links (packets that were corrupted on reception on the MAV)
	ErrorsComm uint16
	// Autopilot-specific errors
	ErrorsCount1 uint16
	// Autopilot-specific errors
	ErrorsCount2 uint16
	// Autopilot-specific errors
	ErrorsCount3 uint16
	// Autopilot-specific errors
	ErrorsCount4 uint16
	// Bitmap showing which onboard controllers and sensors are present. Value of 0: not present. Value of 1: present.
	OnboardControlSensorsPresentExtended MAV_SYS_STATUS_SENSOR_EXTENDED `mavenum:"uint32" mavext:"true"`
	// Bitmap showing which onboard controllers and sensors are enabled:  Value of 0: not enabled. Value of 1: enabled.
	OnboardControlSensorsEnabledExtended MAV_SYS_STATUS_SENSOR_EXTENDED `mavenum:"uint32" mavext:"true"`
	// Bitmap showing which onboard controllers and sensors have an error (or are operational). Value of 0: error. Value of 1: healthy.
	OnboardControlSensorsHealthExtended MAV_SYS_STATUS_SENSOR_EXTENDED `mavenum:"uint32" mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageSysStatus) GetID() uint32 {
	return 1
}
