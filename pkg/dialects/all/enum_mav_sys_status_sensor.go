//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package all

import (
	"errors"
)

// These encode the sensors whose status is sent as part of the SYS_STATUS message.
type MAV_SYS_STATUS_SENSOR uint32

const (
	// 0x01 3D gyro
	MAV_SYS_STATUS_SENSOR_3D_GYRO MAV_SYS_STATUS_SENSOR = 1
	// 0x02 3D accelerometer
	MAV_SYS_STATUS_SENSOR_3D_ACCEL MAV_SYS_STATUS_SENSOR = 2
	// 0x04 3D magnetometer
	MAV_SYS_STATUS_SENSOR_3D_MAG MAV_SYS_STATUS_SENSOR = 4
	// 0x08 absolute pressure
	MAV_SYS_STATUS_SENSOR_ABSOLUTE_PRESSURE MAV_SYS_STATUS_SENSOR = 8
	// 0x10 differential pressure
	MAV_SYS_STATUS_SENSOR_DIFFERENTIAL_PRESSURE MAV_SYS_STATUS_SENSOR = 16
	// 0x20 GPS
	MAV_SYS_STATUS_SENSOR_GPS MAV_SYS_STATUS_SENSOR = 32
	// 0x40 optical flow
	MAV_SYS_STATUS_SENSOR_OPTICAL_FLOW MAV_SYS_STATUS_SENSOR = 64
	// 0x80 computer vision position
	MAV_SYS_STATUS_SENSOR_VISION_POSITION MAV_SYS_STATUS_SENSOR = 128
	// 0x100 laser based position
	MAV_SYS_STATUS_SENSOR_LASER_POSITION MAV_SYS_STATUS_SENSOR = 256
	// 0x200 external ground truth (Vicon or Leica)
	MAV_SYS_STATUS_SENSOR_EXTERNAL_GROUND_TRUTH MAV_SYS_STATUS_SENSOR = 512
	// 0x400 3D angular rate control
	MAV_SYS_STATUS_SENSOR_ANGULAR_RATE_CONTROL MAV_SYS_STATUS_SENSOR = 1024
	// 0x800 attitude stabilization
	MAV_SYS_STATUS_SENSOR_ATTITUDE_STABILIZATION MAV_SYS_STATUS_SENSOR = 2048
	// 0x1000 yaw position
	MAV_SYS_STATUS_SENSOR_YAW_POSITION MAV_SYS_STATUS_SENSOR = 4096
	// 0x2000 z/altitude control
	MAV_SYS_STATUS_SENSOR_Z_ALTITUDE_CONTROL MAV_SYS_STATUS_SENSOR = 8192
	// 0x4000 x/y position control
	MAV_SYS_STATUS_SENSOR_XY_POSITION_CONTROL MAV_SYS_STATUS_SENSOR = 16384
	// 0x8000 motor outputs / control
	MAV_SYS_STATUS_SENSOR_MOTOR_OUTPUTS MAV_SYS_STATUS_SENSOR = 32768
	// 0x10000 rc receiver
	MAV_SYS_STATUS_SENSOR_RC_RECEIVER MAV_SYS_STATUS_SENSOR = 65536
	// 0x20000 2nd 3D gyro
	MAV_SYS_STATUS_SENSOR_3D_GYRO2 MAV_SYS_STATUS_SENSOR = 131072
	// 0x40000 2nd 3D accelerometer
	MAV_SYS_STATUS_SENSOR_3D_ACCEL2 MAV_SYS_STATUS_SENSOR = 262144
	// 0x80000 2nd 3D magnetometer
	MAV_SYS_STATUS_SENSOR_3D_MAG2 MAV_SYS_STATUS_SENSOR = 524288
	// 0x100000 geofence
	MAV_SYS_STATUS_GEOFENCE MAV_SYS_STATUS_SENSOR = 1048576
	// 0x200000 AHRS subsystem health
	MAV_SYS_STATUS_AHRS MAV_SYS_STATUS_SENSOR = 2097152
	// 0x400000 Terrain subsystem health
	MAV_SYS_STATUS_TERRAIN MAV_SYS_STATUS_SENSOR = 4194304
	// 0x800000 Motors are reversed
	MAV_SYS_STATUS_REVERSE_MOTOR MAV_SYS_STATUS_SENSOR = 8388608
	// 0x1000000 Logging
	MAV_SYS_STATUS_LOGGING MAV_SYS_STATUS_SENSOR = 16777216
	// 0x2000000 Battery
	MAV_SYS_STATUS_SENSOR_BATTERY MAV_SYS_STATUS_SENSOR = 33554432
	// 0x4000000 Proximity
	MAV_SYS_STATUS_SENSOR_PROXIMITY MAV_SYS_STATUS_SENSOR = 67108864
	// 0x8000000 Satellite Communication
	MAV_SYS_STATUS_SENSOR_SATCOM MAV_SYS_STATUS_SENSOR = 134217728
	// 0x10000000 pre-arm check status. Always healthy when armed
	MAV_SYS_STATUS_PREARM_CHECK MAV_SYS_STATUS_SENSOR = 268435456
	// 0x20000000 Avoidance/collision prevention
	MAV_SYS_STATUS_OBSTACLE_AVOIDANCE MAV_SYS_STATUS_SENSOR = 536870912
	// 0x40000000 propulsion (actuator, esc, motor or propellor)
	MAV_SYS_STATUS_SENSOR_PROPULSION MAV_SYS_STATUS_SENSOR = 1073741824
	// 0x80000000 Extended bit-field are used for further sensor status bits (needs to be set in onboard_control_sensors_present only)
	MAV_SYS_STATUS_EXTENSION_USED MAV_SYS_STATUS_SENSOR = 2147483648
)

var labels_MAV_SYS_STATUS_SENSOR = map[MAV_SYS_STATUS_SENSOR]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_SYS_STATUS_SENSOR) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_SYS_STATUS_SENSOR[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_SYS_STATUS_SENSOR = map[string]MAV_SYS_STATUS_SENSOR{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_SYS_STATUS_SENSOR) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_SYS_STATUS_SENSOR[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_SYS_STATUS_SENSOR) String() string {
	if l, ok := labels_MAV_SYS_STATUS_SENSOR[e]; ok {
		return l
	}
	return "invalid value"
}
