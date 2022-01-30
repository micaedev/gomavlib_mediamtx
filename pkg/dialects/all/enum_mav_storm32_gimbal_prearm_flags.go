//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package all

import (
	"errors"
)

// STorM32 gimbal prearm check flags.
type MAV_STORM32_GIMBAL_PREARM_FLAGS uint32

const (
	// STorM32 gimbal is in normal state.
	MAV_STORM32_GIMBAL_PREARM_FLAGS_IS_NORMAL MAV_STORM32_GIMBAL_PREARM_FLAGS = 1
	// The IMUs are healthy and working normally.
	MAV_STORM32_GIMBAL_PREARM_FLAGS_IMUS_WORKING MAV_STORM32_GIMBAL_PREARM_FLAGS = 2
	// The motors are active and working normally.
	MAV_STORM32_GIMBAL_PREARM_FLAGS_MOTORS_WORKING MAV_STORM32_GIMBAL_PREARM_FLAGS = 4
	// The encoders are healthy and working normally.
	MAV_STORM32_GIMBAL_PREARM_FLAGS_ENCODERS_WORKING MAV_STORM32_GIMBAL_PREARM_FLAGS = 8
	// A battery voltage is applied and is in range.
	MAV_STORM32_GIMBAL_PREARM_FLAGS_VOLTAGE_OK MAV_STORM32_GIMBAL_PREARM_FLAGS = 16
	// ???.
	MAV_STORM32_GIMBAL_PREARM_FLAGS_VIRTUALCHANNELS_RECEIVING MAV_STORM32_GIMBAL_PREARM_FLAGS = 32
	// Mavlink messages are being received.
	MAV_STORM32_GIMBAL_PREARM_FLAGS_MAVLINK_RECEIVING MAV_STORM32_GIMBAL_PREARM_FLAGS = 64
	// The STorM32Link data indicates QFix.
	MAV_STORM32_GIMBAL_PREARM_FLAGS_STORM32LINK_QFIX MAV_STORM32_GIMBAL_PREARM_FLAGS = 128
	// The STorM32Link is working.
	MAV_STORM32_GIMBAL_PREARM_FLAGS_STORM32LINK_WORKING MAV_STORM32_GIMBAL_PREARM_FLAGS = 256
	// The camera has been found and is connected.
	MAV_STORM32_GIMBAL_PREARM_FLAGS_CAMERA_CONNECTED MAV_STORM32_GIMBAL_PREARM_FLAGS = 512
	// The signal on the AUX0 input pin is low.
	MAV_STORM32_GIMBAL_PREARM_FLAGS_AUX0_LOW MAV_STORM32_GIMBAL_PREARM_FLAGS = 1024
	// The signal on the AUX1 input pin is low.
	MAV_STORM32_GIMBAL_PREARM_FLAGS_AUX1_LOW MAV_STORM32_GIMBAL_PREARM_FLAGS = 2048
	// The NTLogger is working normally.
	MAV_STORM32_GIMBAL_PREARM_FLAGS_NTLOGGER_WORKING MAV_STORM32_GIMBAL_PREARM_FLAGS = 4096
)

var labels_MAV_STORM32_GIMBAL_PREARM_FLAGS = map[MAV_STORM32_GIMBAL_PREARM_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_STORM32_GIMBAL_PREARM_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_STORM32_GIMBAL_PREARM_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_STORM32_GIMBAL_PREARM_FLAGS = map[string]MAV_STORM32_GIMBAL_PREARM_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_STORM32_GIMBAL_PREARM_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_STORM32_GIMBAL_PREARM_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_STORM32_GIMBAL_PREARM_FLAGS) String() string {
	if l, ok := labels_MAV_STORM32_GIMBAL_PREARM_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
