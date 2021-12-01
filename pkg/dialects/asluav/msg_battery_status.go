//autogenerated:yes
//nolint:golint,misspell,govet,lll
package asluav

// Battery information. Updates GCS with flight controller battery status. Smart batteries also use this message, but may additionally send SMART_BATTERY_INFO.
type MessageBatteryStatus struct {
	// Battery ID
	Id uint8
	// Function of the battery
	BatteryFunction MAV_BATTERY_FUNCTION `mavenum:"uint8"`
	// Type (chemistry) of the battery
	Type MAV_BATTERY_TYPE `mavenum:"uint8"`
	// Temperature of the battery. INT16_MAX for unknown temperature.
	Temperature int16
	// Battery voltage of cells 1 to 10 (see voltages_ext for cells 11-14). Cells in this field above the valid cell count for this battery should have the UINT16_MAX value. If individual cell voltages are unknown or not measured for this battery, then the overall battery voltage should be filled in cell 0, with all others set to UINT16_MAX. If the voltage of the battery is greater than (UINT16_MAX - 1), then cell 0 should be set to (UINT16_MAX - 1), and cell 1 to the remaining voltage. This can be extended to multiple cells if the total voltage is greater than 2 * (UINT16_MAX - 1).
	Voltages [10]uint16
	// Battery current, -1: autopilot does not measure the current
	CurrentBattery int16
	// Consumed charge, -1: autopilot does not provide consumption estimate
	CurrentConsumed int32
	// Consumed energy, -1: autopilot does not provide energy consumption estimate
	EnergyConsumed int32
	// Remaining battery energy. Values: [0-100], -1: autopilot does not estimate the remaining battery.
	BatteryRemaining int8
	// Remaining battery time, 0: autopilot does not provide remaining battery time estimate
	TimeRemaining int32 `mavext:"true"`
	// State for extent of discharge, provided by autopilot for warning or external reactions
	ChargeState MAV_BATTERY_CHARGE_STATE `mavenum:"uint8" mavext:"true"`
	// Battery voltages for cells 11 to 14. Cells above the valid cell count for this battery should have a value of 0, where zero indicates not supported (note, this is different than for the voltages field and allows empty byte truncation). If the measured value is 0 then 1 should be sent instead.
	VoltagesExt [4]uint16 `mavext:"true"`
	// Battery mode. Default (0) is that battery mode reporting is not supported or battery is in normal-use mode.
	Mode MAV_BATTERY_MODE `mavenum:"uint8" mavext:"true"`
	// Fault/health indications. These should be set when charge_state is MAV_BATTERY_CHARGE_STATE_FAILED or MAV_BATTERY_CHARGE_STATE_UNHEALTHY (if not, fault reporting is not supported).
	FaultBitmask MAV_BATTERY_FAULT `mavenum:"uint32" mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageBatteryStatus) GetID() uint32 {
	return 147
}
