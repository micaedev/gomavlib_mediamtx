//autogenerated:yes
//nolint:golint,misspell,govet,lll
package minimal

// The heartbeat message shows that a system or component is present and responding. The type and autopilot fields (along with the message component id), allow the receiving system to treat further messages from this system appropriately (e.g. by laying out the user interface based on the autopilot). This microservice is documented at https://mavlink.io/en/services/heartbeat.html
type MessageHeartbeat struct {
	// Vehicle or component type. For a flight controller component the vehicle type (quadrotor, helicopter, etc.). For other components the component type (e.g. camera, gimbal, etc.). This should be used in preference to component id for identifying the component type.
	Type MAV_TYPE `mavenum:"uint8"`
	// Autopilot type / class. Use MAV_AUTOPILOT_INVALID for components that are not flight controllers.
	Autopilot MAV_AUTOPILOT `mavenum:"uint8"`
	// System mode bitmap.
	BaseMode MAV_MODE_FLAG `mavenum:"uint8"`
	// A bitfield for use for autopilot-specific flags
	CustomMode uint32
	// System status flag.
	SystemStatus MAV_STATE `mavenum:"uint8"`
	// MAVLink version, not writable by user, gets added by protocol because of magic data type: uint8_t_mavlink_version
	MavlinkVersion uint8
}

// GetID implements the msg.Message interface.
func (*MessageHeartbeat) GetID() uint32 {
	return 0
}
