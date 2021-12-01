//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package ualberta

import (
	"errors"
)

// Commands to be executed by the MAV. They can be executed on user request, or as part of a mission script. If the action is used in a mission, the parameter mapping to the waypoint/mission message is as follows: Param 1, Param 2, Param 3, Param 4, X: Param 5, Y:Param 6, Z:Param 7. This command list is similar what ARINC 424 is for commercial aircraft: A data format how to interpret waypoint/mission data. NaN and INT32_MAX may be used in float/integer params (respectively) to indicate optional/default values (e.g. to use the component's current yaw or latitude rather than a specific value). See https://mavlink.io/en/guide/xml_schema.html#MAV_CMD for information about the structure of the MAV_CMD entries
type MAV_CMD int

const (
	// Navigate to waypoint.
	MAV_CMD_NAV_WAYPOINT MAV_CMD = 16
	// Loiter around this waypoint an unlimited amount of time
	MAV_CMD_NAV_LOITER_UNLIM MAV_CMD = 17
	// Loiter around this waypoint for X turns
	MAV_CMD_NAV_LOITER_TURNS MAV_CMD = 18
	// Loiter at the specified latitude, longitude and altitude for a certain amount of time. Multicopter vehicles stop at the point (within a vehicle-specific acceptance radius). Forward-only moving vehicles (e.g. fixed-wing) circle the point with the specified radius/direction. If the Heading Required parameter (2) is non-zero forward moving aircraft will only leave the loiter circle once heading towards the next waypoint.
	MAV_CMD_NAV_LOITER_TIME MAV_CMD = 19
	// Return to launch location
	MAV_CMD_NAV_RETURN_TO_LAUNCH MAV_CMD = 20
	// Land at location.
	MAV_CMD_NAV_LAND MAV_CMD = 21
	// Takeoff from ground / hand. Vehicles that support multiple takeoff modes (e.g. VTOL quadplane) should take off using the currently configured mode.
	MAV_CMD_NAV_TAKEOFF MAV_CMD = 22
	// Land at local position (local frame only)
	MAV_CMD_NAV_LAND_LOCAL MAV_CMD = 23
	// Takeoff from local position (local frame only)
	MAV_CMD_NAV_TAKEOFF_LOCAL MAV_CMD = 24
	// Vehicle following, i.e. this waypoint represents the position of a moving vehicle
	MAV_CMD_NAV_FOLLOW MAV_CMD = 25
	// Continue on the current course and climb/descend to specified altitude.  When the altitude is reached continue to the next command (i.e., don't proceed to the next command until the desired altitude is reached.
	MAV_CMD_NAV_CONTINUE_AND_CHANGE_ALT MAV_CMD = 30
	// Begin loiter at the specified Latitude and Longitude.  If Lat=Lon=0, then loiter at the current position.  Don't consider the navigation command complete (don't leave loiter) until the altitude has been reached. Additionally, if the Heading Required parameter is non-zero the aircraft will not leave the loiter until heading toward the next waypoint.
	MAV_CMD_NAV_LOITER_TO_ALT MAV_CMD = 31
	// Begin following a target
	MAV_CMD_DO_FOLLOW MAV_CMD = 32
	// Reposition the MAV after a follow target command has been sent
	MAV_CMD_DO_FOLLOW_REPOSITION MAV_CMD = 33
	// Start orbiting on the circumference of a circle defined by the parameters. Setting values to NaN/INT32_MAX (as appropriate) results in using defaults.
	MAV_CMD_DO_ORBIT MAV_CMD = 34
	// Sets the region of interest (ROI) for a sensor set or the vehicle itself. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras.
	MAV_CMD_NAV_ROI MAV_CMD = 80
	// Control autonomous path planning on the MAV.
	MAV_CMD_NAV_PATHPLANNING MAV_CMD = 81
	// Navigate to waypoint using a spline path.
	MAV_CMD_NAV_SPLINE_WAYPOINT MAV_CMD = 82
	// Takeoff from ground using VTOL mode, and transition to forward flight with specified heading. The command should be ignored by vehicles that dont support both VTOL and fixed-wing flight (multicopters, boats,etc.).
	MAV_CMD_NAV_VTOL_TAKEOFF MAV_CMD = 84
	// Land using VTOL mode
	MAV_CMD_NAV_VTOL_LAND MAV_CMD = 85
	// hand control over to an external controller
	MAV_CMD_NAV_GUIDED_ENABLE MAV_CMD = 92
	// Delay the next navigation command a number of seconds or until a specified time
	MAV_CMD_NAV_DELAY MAV_CMD = 93
	// Descend and place payload. Vehicle moves to specified location, descends until it detects a hanging payload has reached the ground, and then releases the payload. If ground is not detected before the reaching the maximum descent value (param1), the command will complete without releasing the payload.
	MAV_CMD_NAV_PAYLOAD_PLACE MAV_CMD = 94
	// NOP - This command is only used to mark the upper limit of the NAV/ACTION commands in the enumeration
	MAV_CMD_NAV_LAST MAV_CMD = 95
	// Delay mission state machine.
	MAV_CMD_CONDITION_DELAY MAV_CMD = 112
	// Ascend/descend to target altitude at specified rate. Delay mission state machine until desired altitude reached.
	MAV_CMD_CONDITION_CHANGE_ALT MAV_CMD = 113
	// Delay mission state machine until within desired distance of next NAV point.
	MAV_CMD_CONDITION_DISTANCE MAV_CMD = 114
	// Reach a certain target angle.
	MAV_CMD_CONDITION_YAW MAV_CMD = 115
	// NOP - This command is only used to mark the upper limit of the CONDITION commands in the enumeration
	MAV_CMD_CONDITION_LAST MAV_CMD = 159
	// Set system mode.
	MAV_CMD_DO_SET_MODE MAV_CMD = 176
	// Jump to the desired command in the mission list.  Repeat this action only the specified number of times
	MAV_CMD_DO_JUMP MAV_CMD = 177
	// Change speed and/or throttle set points.
	MAV_CMD_DO_CHANGE_SPEED MAV_CMD = 178
	// Changes the home location either to the current location or a specified location.
	MAV_CMD_DO_SET_HOME MAV_CMD = 179
	// Set a system parameter.  Caution!  Use of this command requires knowledge of the numeric enumeration value of the parameter.
	MAV_CMD_DO_SET_PARAMETER MAV_CMD = 180
	// Set a relay to a condition.
	MAV_CMD_DO_SET_RELAY MAV_CMD = 181
	// Cycle a relay on and off for a desired number of cycles with a desired period.
	MAV_CMD_DO_REPEAT_RELAY MAV_CMD = 182
	// Set a servo to a desired PWM value.
	MAV_CMD_DO_SET_SERVO MAV_CMD = 183
	// Cycle a between its nominal setting and a desired PWM for a desired number of cycles with a desired period.
	MAV_CMD_DO_REPEAT_SERVO MAV_CMD = 184
	// Terminate flight immediately
	MAV_CMD_DO_FLIGHTTERMINATION MAV_CMD = 185
	// Change altitude set point.
	MAV_CMD_DO_CHANGE_ALTITUDE MAV_CMD = 186
	// Sets actuators (e.g. servos) to a desired value. The actuator numbers are mapped to specific outputs (e.g. on any MAIN or AUX PWM or UAVCAN) using a flight-stack specific mechanism (i.e. a parameter).
	MAV_CMD_DO_SET_ACTUATOR MAV_CMD = 187
	// Mission command to perform a landing. This is used as a marker in a mission to tell the autopilot where a sequence of mission items that represents a landing starts. It may also be sent via a COMMAND_LONG to trigger a landing, in which case the nearest (geographically) landing sequence in the mission will be used. The Latitude/Longitude is optional, and may be set to 0 if not needed. If specified then it will be used to help find the closest landing sequence.
	MAV_CMD_DO_LAND_START MAV_CMD = 189
	// Mission command to perform a landing from a rally point.
	MAV_CMD_DO_RALLY_LAND MAV_CMD = 190
	// Mission command to safely abort an autonomous landing.
	MAV_CMD_DO_GO_AROUND MAV_CMD = 191
	// Reposition the vehicle to a specific WGS84 global position.
	MAV_CMD_DO_REPOSITION MAV_CMD = 192
	// If in a GPS controlled position mode, hold the current position or continue.
	MAV_CMD_DO_PAUSE_CONTINUE MAV_CMD = 193
	// Set moving direction to forward or reverse.
	MAV_CMD_DO_SET_REVERSE MAV_CMD = 194
	// Sets the region of interest (ROI) to a location. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras. This command can be sent to a gimbal manager but not to a gimbal device. A gimbal is not to react to this message.
	MAV_CMD_DO_SET_ROI_LOCATION MAV_CMD = 195
	// Sets the region of interest (ROI) to be toward next waypoint, with optional pitch/roll/yaw offset. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras. This command can be sent to a gimbal manager but not to a gimbal device. A gimbal device is not to react to this message.
	MAV_CMD_DO_SET_ROI_WPNEXT_OFFSET MAV_CMD = 196
	// Cancels any previous ROI command returning the vehicle/sensors to default flight characteristics. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras. This command can be sent to a gimbal manager but not to a gimbal device. A gimbal device is not to react to this message. After this command the gimbal manager should go back to manual input if available, and otherwise assume a neutral position.
	MAV_CMD_DO_SET_ROI_NONE MAV_CMD = 197
	// Mount tracks system with specified system ID. Determination of target vehicle position may be done with GLOBAL_POSITION_INT or any other means. This command can be sent to a gimbal manager but not to a gimbal device. A gimbal device is not to react to this message.
	MAV_CMD_DO_SET_ROI_SYSID MAV_CMD = 198
	// Control onboard camera system.
	MAV_CMD_DO_CONTROL_VIDEO MAV_CMD = 200
	// Sets the region of interest (ROI) for a sensor set or the vehicle itself. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras.
	MAV_CMD_DO_SET_ROI MAV_CMD = 201
	// Configure digital camera. This is a fallback message for systems that have not yet implemented PARAM_EXT_XXX messages and camera definition files (see https://mavlink.io/en/services/camera_def.html ).
	MAV_CMD_DO_DIGICAM_CONFIGURE MAV_CMD = 202
	// Control digital camera. This is a fallback message for systems that have not yet implemented PARAM_EXT_XXX messages and camera definition files (see https://mavlink.io/en/services/camera_def.html ).
	MAV_CMD_DO_DIGICAM_CONTROL MAV_CMD = 203
	// Mission command to configure a camera or antenna mount
	MAV_CMD_DO_MOUNT_CONFIGURE MAV_CMD = 204
	// Mission command to control a camera or antenna mount
	MAV_CMD_DO_MOUNT_CONTROL MAV_CMD = 205
	// Mission command to set camera trigger distance for this flight. The camera is triggered each time this distance is exceeded. This command can also be used to set the shutter integration time for the camera.
	MAV_CMD_DO_SET_CAM_TRIGG_DIST MAV_CMD = 206
	// Mission command to enable the geofence
	MAV_CMD_DO_FENCE_ENABLE MAV_CMD = 207
	// Mission item/command to release a parachute or enable/disable auto release.
	MAV_CMD_DO_PARACHUTE MAV_CMD = 208
	// Command to perform motor test.
	MAV_CMD_DO_MOTOR_TEST MAV_CMD = 209
	// Change to/from inverted flight.
	MAV_CMD_DO_INVERTED_FLIGHT MAV_CMD = 210
	// Mission command to operate a gripper.
	MAV_CMD_DO_GRIPPER MAV_CMD = 211
	// Enable/disable autotune.
	MAV_CMD_DO_AUTOTUNE_ENABLE MAV_CMD = 212
	// Sets a desired vehicle turn angle and speed change.
	MAV_CMD_NAV_SET_YAW_SPEED MAV_CMD = 213
	// Mission command to set camera trigger interval for this flight. If triggering is enabled, the camera is triggered each time this interval expires. This command can also be used to set the shutter integration time for the camera.
	MAV_CMD_DO_SET_CAM_TRIGG_INTERVAL MAV_CMD = 214
	// Mission command to control a camera or antenna mount, using a quaternion as reference.
	MAV_CMD_DO_MOUNT_CONTROL_QUAT MAV_CMD = 220
	// set id of master controller
	MAV_CMD_DO_GUIDED_MASTER MAV_CMD = 221
	// Set limits for external control
	MAV_CMD_DO_GUIDED_LIMITS MAV_CMD = 222
	// Control vehicle engine. This is interpreted by the vehicles engine controller to change the target engine state. It is intended for vehicles with internal combustion engines
	MAV_CMD_DO_ENGINE_CONTROL MAV_CMD = 223
	// Set the mission item with sequence number seq as current item. This means that the MAV will continue to this mission item on the shortest path (not following the mission items in-between).
	MAV_CMD_DO_SET_MISSION_CURRENT MAV_CMD = 224
	// NOP - This command is only used to mark the upper limit of the DO commands in the enumeration
	MAV_CMD_DO_LAST MAV_CMD = 240
	// Trigger calibration. This command will be only accepted if in pre-flight mode. Except for Temperature Calibration, only one sensor should be set in a single message and all others should be zero.
	MAV_CMD_PREFLIGHT_CALIBRATION MAV_CMD = 241
	// Set sensor offsets. This command will be only accepted if in pre-flight mode.
	MAV_CMD_PREFLIGHT_SET_SENSOR_OFFSETS MAV_CMD = 242
	// Trigger UAVCAN configuration (actuator ID assignment and direction mapping). Note that this maps to the legacy UAVCAN v0 function UAVCAN_ENUMERATE, which is intended to be executed just once during initial vehicle configuration (it is not a normal pre-flight command and has been poorly named).
	MAV_CMD_PREFLIGHT_UAVCAN MAV_CMD = 243
	// Request storage of different parameter values and logs. This command will be only accepted if in pre-flight mode.
	MAV_CMD_PREFLIGHT_STORAGE MAV_CMD = 245
	// Request the reboot or shutdown of system components.
	MAV_CMD_PREFLIGHT_REBOOT_SHUTDOWN MAV_CMD = 246
	// Override current mission with command to pause mission, pause mission and move to position, continue/resume mission. When param 1 indicates that the mission is paused (MAV_GOTO_DO_HOLD), param 2 defines whether it holds in place or moves to another position.
	MAV_CMD_OVERRIDE_GOTO MAV_CMD = 252
	// Mission command to set a Camera Auto Mount Pivoting Oblique Survey (Replaces CAM_TRIGG_DIST for this purpose). The camera is triggered each time this distance is exceeded, then the mount moves to the next position. Params 4~6 set-up the angle limits and number of positions for oblique survey, where mount-enabled vehicles automatically roll the camera between shots to emulate an oblique camera setup (providing an increased HFOV). This command can also be used to set the shutter integration time for the camera.
	MAV_CMD_OBLIQUE_SURVEY MAV_CMD = 260
	// start running a mission
	MAV_CMD_MISSION_START MAV_CMD = 300
	// Actuator testing command. This is similar to MAV_CMD_DO_MOTOR_TEST but operates on the level of output functions, i.e. it is possible to test Motor1 independent from which output it is configured on. Autopilots typically refuse this command while armed.
	MAV_CMD_ACTUATOR_TEST MAV_CMD = 310
	// Actuator configuration command.
	MAV_CMD_CONFIGURE_ACTUATOR MAV_CMD = 311
	// Arms / Disarms a component
	MAV_CMD_COMPONENT_ARM_DISARM MAV_CMD = 400
	// Instructs system to run pre-arm checks. This command should return MAV_RESULT_TEMPORARILY_REJECTED in the case the system is armed, otherwise MAV_RESULT_ACCEPTED. Note that the return value from executing this command does not indicate whether the vehicle is armable or not, just whether the system has successfully run/is currently running the checks.  The result of the checks is reflected in the SYS_STATUS message.
	MAV_CMD_RUN_PREARM_CHECKS MAV_CMD = 401
	// Turns illuminators ON/OFF. An illuminator is a light source that is used for lighting up dark areas external to the sytstem: e.g. a torch or searchlight (as opposed to a light source for illuminating the system itself, e.g. an indicator light).
	MAV_CMD_ILLUMINATOR_ON_OFF MAV_CMD = 405
	// Request the home position from the vehicle.
	MAV_CMD_GET_HOME_POSITION MAV_CMD = 410
	// Inject artificial failure for testing purposes. Note that autopilots should implement an additional protection before accepting this command such as a specific param setting.
	MAV_CMD_INJECT_FAILURE MAV_CMD = 420
	// Starts receiver pairing.
	MAV_CMD_START_RX_PAIR MAV_CMD = 500
	// Request the interval between messages for a particular MAVLink message ID. The receiver should ACK the command and then emit its response in a MESSAGE_INTERVAL message.
	MAV_CMD_GET_MESSAGE_INTERVAL MAV_CMD = 510
	// Set the interval between messages for a particular MAVLink message ID. This interface replaces REQUEST_DATA_STREAM.
	MAV_CMD_SET_MESSAGE_INTERVAL MAV_CMD = 511
	// Request the target system(s) emit a single instance of a specified message (i.e. a "one-shot" version of MAV_CMD_SET_MESSAGE_INTERVAL).
	MAV_CMD_REQUEST_MESSAGE MAV_CMD = 512
	// Request MAVLink protocol version compatibility. All receivers should ACK the command and then emit their capabilities in an PROTOCOL_VERSION message
	MAV_CMD_REQUEST_PROTOCOL_VERSION MAV_CMD = 519
	// Request autopilot capabilities. The receiver should ACK the command and then emit its capabilities in an AUTOPILOT_VERSION message
	MAV_CMD_REQUEST_AUTOPILOT_CAPABILITIES MAV_CMD = 520
	// Request camera information (CAMERA_INFORMATION).
	MAV_CMD_REQUEST_CAMERA_INFORMATION MAV_CMD = 521
	// Request camera settings (CAMERA_SETTINGS).
	MAV_CMD_REQUEST_CAMERA_SETTINGS MAV_CMD = 522
	// Request storage information (STORAGE_INFORMATION). Use the command's target_component to target a specific component's storage.
	MAV_CMD_REQUEST_STORAGE_INFORMATION MAV_CMD = 525
	// Format a storage medium. Once format is complete, a STORAGE_INFORMATION message is sent. Use the command's target_component to target a specific component's storage.
	MAV_CMD_STORAGE_FORMAT MAV_CMD = 526
	// Request camera capture status (CAMERA_CAPTURE_STATUS)
	MAV_CMD_REQUEST_CAMERA_CAPTURE_STATUS MAV_CMD = 527
	// Request flight information (FLIGHT_INFORMATION)
	MAV_CMD_REQUEST_FLIGHT_INFORMATION MAV_CMD = 528
	// Reset all camera settings to Factory Default
	MAV_CMD_RESET_CAMERA_SETTINGS MAV_CMD = 529
	// Set camera running mode. Use NaN for reserved values. GCS will send a MAV_CMD_REQUEST_VIDEO_STREAM_STATUS command after a mode change if the camera supports video streaming.
	MAV_CMD_SET_CAMERA_MODE MAV_CMD = 530
	// Set camera zoom. Camera must respond with a CAMERA_SETTINGS message (on success).
	MAV_CMD_SET_CAMERA_ZOOM MAV_CMD = 531
	// Set camera focus. Camera must respond with a CAMERA_SETTINGS message (on success).
	MAV_CMD_SET_CAMERA_FOCUS MAV_CMD = 532
	// Set that a particular storage is the preferred location for saving photos, videos, and/or other media (e.g. to set that an SD card is used for storing videos).
	// There can only be one preferred save location for each particular media type: setting a media usage flag will clear/reset that same flag if set on any other storage.
	// If no flag is set the system should use its default storage.
	// A target system can choose to always use default storage, in which case it should ACK the command with MAV_RESULT_UNSUPPORTED.
	// A target system can choose to not allow a particular storage to be set as preferred storage, in which case it should ACK the command with MAV_RESULT_DENIED.
	MAV_CMD_SET_STORAGE_USAGE MAV_CMD = 533
	// Tagged jump target. Can be jumped to with MAV_CMD_DO_JUMP_TAG.
	MAV_CMD_JUMP_TAG MAV_CMD = 600
	// Jump to the matching tag in the mission list. Repeat this action for the specified number of times. A mission should contain a single matching tag for each jump. If this is not the case then a jump to a missing tag should complete the mission, and a jump where there are multiple matching tags should always select the one with the lowest mission sequence number.
	MAV_CMD_DO_JUMP_TAG MAV_CMD = 601
	// High level setpoint to be sent to a gimbal manager to set a gimbal attitude. It is possible to set combinations of the values below. E.g. an angle as well as a desired angular rate can be used to get to this angle at a certain angular rate, or an angular rate only will result in continuous turning. NaN is to be used to signal unset. Note: a gimbal is never to react to this command but only the gimbal manager.
	MAV_CMD_DO_GIMBAL_MANAGER_PITCHYAW MAV_CMD = 1000
	// Gimbal configuration to set which sysid/compid is in primary and secondary control.
	MAV_CMD_DO_GIMBAL_MANAGER_CONFIGURE MAV_CMD = 1001
	// Start image capture sequence. Sends CAMERA_IMAGE_CAPTURED after each capture. Use NaN for reserved values.
	MAV_CMD_IMAGE_START_CAPTURE MAV_CMD = 2000
	// Stop image capture sequence Use NaN for reserved values.
	MAV_CMD_IMAGE_STOP_CAPTURE MAV_CMD = 2001
	// Re-request a CAMERA_IMAGE_CAPTURED message.
	MAV_CMD_REQUEST_CAMERA_IMAGE_CAPTURE MAV_CMD = 2002
	// Enable or disable on-board camera triggering system.
	MAV_CMD_DO_TRIGGER_CONTROL MAV_CMD = 2003
	// If the camera supports point visual tracking (CAMERA_CAP_FLAGS_HAS_TRACKING_POINT is set), this command allows to initiate the tracking.
	MAV_CMD_CAMERA_TRACK_POINT MAV_CMD = 2004
	// If the camera supports rectangle visual tracking (CAMERA_CAP_FLAGS_HAS_TRACKING_RECTANGLE is set), this command allows to initiate the tracking.
	MAV_CMD_CAMERA_TRACK_RECTANGLE MAV_CMD = 2005
	// Stops ongoing tracking.
	MAV_CMD_CAMERA_STOP_TRACKING MAV_CMD = 2010
	// Starts video capture (recording).
	MAV_CMD_VIDEO_START_CAPTURE MAV_CMD = 2500
	// Stop the current video capture (recording).
	MAV_CMD_VIDEO_STOP_CAPTURE MAV_CMD = 2501
	// Start video streaming
	MAV_CMD_VIDEO_START_STREAMING MAV_CMD = 2502
	// Stop the given video stream
	MAV_CMD_VIDEO_STOP_STREAMING MAV_CMD = 2503
	// Request video stream information (VIDEO_STREAM_INFORMATION)
	MAV_CMD_REQUEST_VIDEO_STREAM_INFORMATION MAV_CMD = 2504
	// Request video stream status (VIDEO_STREAM_STATUS)
	MAV_CMD_REQUEST_VIDEO_STREAM_STATUS MAV_CMD = 2505
	// Request to start streaming logging data over MAVLink (see also LOGGING_DATA message)
	MAV_CMD_LOGGING_START MAV_CMD = 2510
	// Request to stop streaming log data over MAVLink
	MAV_CMD_LOGGING_STOP           MAV_CMD = 2511
	MAV_CMD_AIRFRAME_CONFIGURATION MAV_CMD = 2520
	// Request to start/stop transmitting over the high latency telemetry
	MAV_CMD_CONTROL_HIGH_LATENCY MAV_CMD = 2600
	// Create a panorama at the current position
	MAV_CMD_PANORAMA_CREATE MAV_CMD = 2800
	// Request VTOL transition
	MAV_CMD_DO_VTOL_TRANSITION MAV_CMD = 3000
	// Request authorization to arm the vehicle to a external entity, the arm authorizer is responsible to request all data that is needs from the vehicle before authorize or deny the request. If approved the progress of command_ack message should be set with period of time that this authorization is valid in seconds or in case it was denied it should be set with one of the reasons in ARM_AUTH_DENIED_REASON.
	MAV_CMD_ARM_AUTHORIZATION_REQUEST MAV_CMD = 3001
	// This command sets the submode to standard guided when vehicle is in guided mode. The vehicle holds position and altitude and the user can input the desired velocities along all three axes.
	MAV_CMD_SET_GUIDED_SUBMODE_STANDARD MAV_CMD = 4000
	// This command sets submode circle when vehicle is in guided mode. Vehicle flies along a circle facing the center of the circle. The user can input the velocity along the circle and change the radius. If no input is given the vehicle will hold position.
	MAV_CMD_SET_GUIDED_SUBMODE_CIRCLE MAV_CMD = 4001
	// Delay mission state machine until gate has been reached.
	MAV_CMD_CONDITION_GATE MAV_CMD = 4501
	// Fence return point (there can only be one such point in a geofence definition). If rally points are supported they should be used instead.
	MAV_CMD_NAV_FENCE_RETURN_POINT MAV_CMD = 5000
	// Fence vertex for an inclusion polygon (the polygon must not be self-intersecting). The vehicle must stay within this area. Minimum of 3 vertices required.
	MAV_CMD_NAV_FENCE_POLYGON_VERTEX_INCLUSION MAV_CMD = 5001
	// Fence vertex for an exclusion polygon (the polygon must not be self-intersecting). The vehicle must stay outside this area. Minimum of 3 vertices required.
	MAV_CMD_NAV_FENCE_POLYGON_VERTEX_EXCLUSION MAV_CMD = 5002
	// Circular fence area. The vehicle must stay inside this area.
	MAV_CMD_NAV_FENCE_CIRCLE_INCLUSION MAV_CMD = 5003
	// Circular fence area. The vehicle must stay outside this area.
	MAV_CMD_NAV_FENCE_CIRCLE_EXCLUSION MAV_CMD = 5004
	// Rally point. You can have multiple rally points defined.
	MAV_CMD_NAV_RALLY_POINT MAV_CMD = 5100
	// Commands the vehicle to respond with a sequence of messages UAVCAN_NODE_INFO, one message per every UAVCAN node that is online. Note that some of the response messages can be lost, which the receiver can detect easily by checking whether every received UAVCAN_NODE_STATUS has a matching message UAVCAN_NODE_INFO received earlier; if not, this command should be sent again in order to request re-transmission of the node information messages.
	MAV_CMD_UAVCAN_GET_NODE_INFO MAV_CMD = 5200
	// Trigger the start of an ADSB-out IDENT. This should only be used when requested to do so by an Air Traffic Controller in controlled airspace. This starts the IDENT which is then typically held for 18 seconds by the hardware per the Mode A, C, and S transponder spec.
	MAV_CMD_DO_ADSB_OUT_IDENT MAV_CMD = 10001
	// Deploy payload on a Lat / Lon / Alt position. This includes the navigation to reach the required release position and velocity.
	MAV_CMD_PAYLOAD_PREPARE_DEPLOY MAV_CMD = 30001
	// Control the payload deployment.
	MAV_CMD_PAYLOAD_CONTROL_DEPLOY MAV_CMD = 30002
	// Magnetometer calibration based on provided known yaw. This allows for fast calibration using WMM field tables in the vehicle, given only the known yaw of the vehicle. If Latitude and longitude are both zero then use the current vehicle location.
	MAV_CMD_FIXED_MAG_CAL_YAW MAV_CMD = 42006
	// Command to operate winch.
	MAV_CMD_DO_WINCH MAV_CMD = 42600
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	MAV_CMD_WAYPOINT_USER_1 MAV_CMD = 31000
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	MAV_CMD_WAYPOINT_USER_2 MAV_CMD = 31001
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	MAV_CMD_WAYPOINT_USER_3 MAV_CMD = 31002
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	MAV_CMD_WAYPOINT_USER_4 MAV_CMD = 31003
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	MAV_CMD_WAYPOINT_USER_5 MAV_CMD = 31004
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	MAV_CMD_SPATIAL_USER_1 MAV_CMD = 31005
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	MAV_CMD_SPATIAL_USER_2 MAV_CMD = 31006
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	MAV_CMD_SPATIAL_USER_3 MAV_CMD = 31007
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	MAV_CMD_SPATIAL_USER_4 MAV_CMD = 31008
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	MAV_CMD_SPATIAL_USER_5 MAV_CMD = 31009
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item.
	MAV_CMD_USER_1 MAV_CMD = 31010
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item.
	MAV_CMD_USER_2 MAV_CMD = 31011
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item.
	MAV_CMD_USER_3 MAV_CMD = 31012
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item.
	MAV_CMD_USER_4 MAV_CMD = 31013
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item.
	MAV_CMD_USER_5 MAV_CMD = 31014
)

var labels_MAV_CMD = map[MAV_CMD]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_CMD) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_CMD[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_CMD = map[string]MAV_CMD{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_CMD) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_CMD[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_CMD) String() string {
	if l, ok := labels_MAV_CMD[e]; ok {
		return l
	}
	return "invalid value"
}
