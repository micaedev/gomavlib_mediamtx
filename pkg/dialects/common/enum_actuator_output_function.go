//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strconv"
)

// Actuator output function. Values greater or equal to 1000 are autopilot-specific.
type ACTUATOR_OUTPUT_FUNCTION uint64

const (
	// No function (disabled).
	ACTUATOR_OUTPUT_FUNCTION_NONE ACTUATOR_OUTPUT_FUNCTION = 0
	// Motor 1
	ACTUATOR_OUTPUT_FUNCTION_MOTOR1 ACTUATOR_OUTPUT_FUNCTION = 1
	// Motor 2
	ACTUATOR_OUTPUT_FUNCTION_MOTOR2 ACTUATOR_OUTPUT_FUNCTION = 2
	// Motor 3
	ACTUATOR_OUTPUT_FUNCTION_MOTOR3 ACTUATOR_OUTPUT_FUNCTION = 3
	// Motor 4
	ACTUATOR_OUTPUT_FUNCTION_MOTOR4 ACTUATOR_OUTPUT_FUNCTION = 4
	// Motor 5
	ACTUATOR_OUTPUT_FUNCTION_MOTOR5 ACTUATOR_OUTPUT_FUNCTION = 5
	// Motor 6
	ACTUATOR_OUTPUT_FUNCTION_MOTOR6 ACTUATOR_OUTPUT_FUNCTION = 6
	// Motor 7
	ACTUATOR_OUTPUT_FUNCTION_MOTOR7 ACTUATOR_OUTPUT_FUNCTION = 7
	// Motor 8
	ACTUATOR_OUTPUT_FUNCTION_MOTOR8 ACTUATOR_OUTPUT_FUNCTION = 8
	// Motor 9
	ACTUATOR_OUTPUT_FUNCTION_MOTOR9 ACTUATOR_OUTPUT_FUNCTION = 9
	// Motor 10
	ACTUATOR_OUTPUT_FUNCTION_MOTOR10 ACTUATOR_OUTPUT_FUNCTION = 10
	// Motor 11
	ACTUATOR_OUTPUT_FUNCTION_MOTOR11 ACTUATOR_OUTPUT_FUNCTION = 11
	// Motor 12
	ACTUATOR_OUTPUT_FUNCTION_MOTOR12 ACTUATOR_OUTPUT_FUNCTION = 12
	// Motor 13
	ACTUATOR_OUTPUT_FUNCTION_MOTOR13 ACTUATOR_OUTPUT_FUNCTION = 13
	// Motor 14
	ACTUATOR_OUTPUT_FUNCTION_MOTOR14 ACTUATOR_OUTPUT_FUNCTION = 14
	// Motor 15
	ACTUATOR_OUTPUT_FUNCTION_MOTOR15 ACTUATOR_OUTPUT_FUNCTION = 15
	// Motor 16
	ACTUATOR_OUTPUT_FUNCTION_MOTOR16 ACTUATOR_OUTPUT_FUNCTION = 16
	// Servo 1
	ACTUATOR_OUTPUT_FUNCTION_SERVO1 ACTUATOR_OUTPUT_FUNCTION = 33
	// Servo 2
	ACTUATOR_OUTPUT_FUNCTION_SERVO2 ACTUATOR_OUTPUT_FUNCTION = 34
	// Servo 3
	ACTUATOR_OUTPUT_FUNCTION_SERVO3 ACTUATOR_OUTPUT_FUNCTION = 35
	// Servo 4
	ACTUATOR_OUTPUT_FUNCTION_SERVO4 ACTUATOR_OUTPUT_FUNCTION = 36
	// Servo 5
	ACTUATOR_OUTPUT_FUNCTION_SERVO5 ACTUATOR_OUTPUT_FUNCTION = 37
	// Servo 6
	ACTUATOR_OUTPUT_FUNCTION_SERVO6 ACTUATOR_OUTPUT_FUNCTION = 38
	// Servo 7
	ACTUATOR_OUTPUT_FUNCTION_SERVO7 ACTUATOR_OUTPUT_FUNCTION = 39
	// Servo 8
	ACTUATOR_OUTPUT_FUNCTION_SERVO8 ACTUATOR_OUTPUT_FUNCTION = 40
	// Servo 9
	ACTUATOR_OUTPUT_FUNCTION_SERVO9 ACTUATOR_OUTPUT_FUNCTION = 41
	// Servo 10
	ACTUATOR_OUTPUT_FUNCTION_SERVO10 ACTUATOR_OUTPUT_FUNCTION = 42
	// Servo 11
	ACTUATOR_OUTPUT_FUNCTION_SERVO11 ACTUATOR_OUTPUT_FUNCTION = 43
	// Servo 12
	ACTUATOR_OUTPUT_FUNCTION_SERVO12 ACTUATOR_OUTPUT_FUNCTION = 44
	// Servo 13
	ACTUATOR_OUTPUT_FUNCTION_SERVO13 ACTUATOR_OUTPUT_FUNCTION = 45
	// Servo 14
	ACTUATOR_OUTPUT_FUNCTION_SERVO14 ACTUATOR_OUTPUT_FUNCTION = 46
	// Servo 15
	ACTUATOR_OUTPUT_FUNCTION_SERVO15 ACTUATOR_OUTPUT_FUNCTION = 47
	// Servo 16
	ACTUATOR_OUTPUT_FUNCTION_SERVO16 ACTUATOR_OUTPUT_FUNCTION = 48
)

var value_to_label_ACTUATOR_OUTPUT_FUNCTION = map[ACTUATOR_OUTPUT_FUNCTION]string{
	ACTUATOR_OUTPUT_FUNCTION_NONE:    "ACTUATOR_OUTPUT_FUNCTION_NONE",
	ACTUATOR_OUTPUT_FUNCTION_MOTOR1:  "ACTUATOR_OUTPUT_FUNCTION_MOTOR1",
	ACTUATOR_OUTPUT_FUNCTION_MOTOR2:  "ACTUATOR_OUTPUT_FUNCTION_MOTOR2",
	ACTUATOR_OUTPUT_FUNCTION_MOTOR3:  "ACTUATOR_OUTPUT_FUNCTION_MOTOR3",
	ACTUATOR_OUTPUT_FUNCTION_MOTOR4:  "ACTUATOR_OUTPUT_FUNCTION_MOTOR4",
	ACTUATOR_OUTPUT_FUNCTION_MOTOR5:  "ACTUATOR_OUTPUT_FUNCTION_MOTOR5",
	ACTUATOR_OUTPUT_FUNCTION_MOTOR6:  "ACTUATOR_OUTPUT_FUNCTION_MOTOR6",
	ACTUATOR_OUTPUT_FUNCTION_MOTOR7:  "ACTUATOR_OUTPUT_FUNCTION_MOTOR7",
	ACTUATOR_OUTPUT_FUNCTION_MOTOR8:  "ACTUATOR_OUTPUT_FUNCTION_MOTOR8",
	ACTUATOR_OUTPUT_FUNCTION_MOTOR9:  "ACTUATOR_OUTPUT_FUNCTION_MOTOR9",
	ACTUATOR_OUTPUT_FUNCTION_MOTOR10: "ACTUATOR_OUTPUT_FUNCTION_MOTOR10",
	ACTUATOR_OUTPUT_FUNCTION_MOTOR11: "ACTUATOR_OUTPUT_FUNCTION_MOTOR11",
	ACTUATOR_OUTPUT_FUNCTION_MOTOR12: "ACTUATOR_OUTPUT_FUNCTION_MOTOR12",
	ACTUATOR_OUTPUT_FUNCTION_MOTOR13: "ACTUATOR_OUTPUT_FUNCTION_MOTOR13",
	ACTUATOR_OUTPUT_FUNCTION_MOTOR14: "ACTUATOR_OUTPUT_FUNCTION_MOTOR14",
	ACTUATOR_OUTPUT_FUNCTION_MOTOR15: "ACTUATOR_OUTPUT_FUNCTION_MOTOR15",
	ACTUATOR_OUTPUT_FUNCTION_MOTOR16: "ACTUATOR_OUTPUT_FUNCTION_MOTOR16",
	ACTUATOR_OUTPUT_FUNCTION_SERVO1:  "ACTUATOR_OUTPUT_FUNCTION_SERVO1",
	ACTUATOR_OUTPUT_FUNCTION_SERVO2:  "ACTUATOR_OUTPUT_FUNCTION_SERVO2",
	ACTUATOR_OUTPUT_FUNCTION_SERVO3:  "ACTUATOR_OUTPUT_FUNCTION_SERVO3",
	ACTUATOR_OUTPUT_FUNCTION_SERVO4:  "ACTUATOR_OUTPUT_FUNCTION_SERVO4",
	ACTUATOR_OUTPUT_FUNCTION_SERVO5:  "ACTUATOR_OUTPUT_FUNCTION_SERVO5",
	ACTUATOR_OUTPUT_FUNCTION_SERVO6:  "ACTUATOR_OUTPUT_FUNCTION_SERVO6",
	ACTUATOR_OUTPUT_FUNCTION_SERVO7:  "ACTUATOR_OUTPUT_FUNCTION_SERVO7",
	ACTUATOR_OUTPUT_FUNCTION_SERVO8:  "ACTUATOR_OUTPUT_FUNCTION_SERVO8",
	ACTUATOR_OUTPUT_FUNCTION_SERVO9:  "ACTUATOR_OUTPUT_FUNCTION_SERVO9",
	ACTUATOR_OUTPUT_FUNCTION_SERVO10: "ACTUATOR_OUTPUT_FUNCTION_SERVO10",
	ACTUATOR_OUTPUT_FUNCTION_SERVO11: "ACTUATOR_OUTPUT_FUNCTION_SERVO11",
	ACTUATOR_OUTPUT_FUNCTION_SERVO12: "ACTUATOR_OUTPUT_FUNCTION_SERVO12",
	ACTUATOR_OUTPUT_FUNCTION_SERVO13: "ACTUATOR_OUTPUT_FUNCTION_SERVO13",
	ACTUATOR_OUTPUT_FUNCTION_SERVO14: "ACTUATOR_OUTPUT_FUNCTION_SERVO14",
	ACTUATOR_OUTPUT_FUNCTION_SERVO15: "ACTUATOR_OUTPUT_FUNCTION_SERVO15",
	ACTUATOR_OUTPUT_FUNCTION_SERVO16: "ACTUATOR_OUTPUT_FUNCTION_SERVO16",
}

var label_to_value_ACTUATOR_OUTPUT_FUNCTION = map[string]ACTUATOR_OUTPUT_FUNCTION{
	"ACTUATOR_OUTPUT_FUNCTION_NONE":    ACTUATOR_OUTPUT_FUNCTION_NONE,
	"ACTUATOR_OUTPUT_FUNCTION_MOTOR1":  ACTUATOR_OUTPUT_FUNCTION_MOTOR1,
	"ACTUATOR_OUTPUT_FUNCTION_MOTOR2":  ACTUATOR_OUTPUT_FUNCTION_MOTOR2,
	"ACTUATOR_OUTPUT_FUNCTION_MOTOR3":  ACTUATOR_OUTPUT_FUNCTION_MOTOR3,
	"ACTUATOR_OUTPUT_FUNCTION_MOTOR4":  ACTUATOR_OUTPUT_FUNCTION_MOTOR4,
	"ACTUATOR_OUTPUT_FUNCTION_MOTOR5":  ACTUATOR_OUTPUT_FUNCTION_MOTOR5,
	"ACTUATOR_OUTPUT_FUNCTION_MOTOR6":  ACTUATOR_OUTPUT_FUNCTION_MOTOR6,
	"ACTUATOR_OUTPUT_FUNCTION_MOTOR7":  ACTUATOR_OUTPUT_FUNCTION_MOTOR7,
	"ACTUATOR_OUTPUT_FUNCTION_MOTOR8":  ACTUATOR_OUTPUT_FUNCTION_MOTOR8,
	"ACTUATOR_OUTPUT_FUNCTION_MOTOR9":  ACTUATOR_OUTPUT_FUNCTION_MOTOR9,
	"ACTUATOR_OUTPUT_FUNCTION_MOTOR10": ACTUATOR_OUTPUT_FUNCTION_MOTOR10,
	"ACTUATOR_OUTPUT_FUNCTION_MOTOR11": ACTUATOR_OUTPUT_FUNCTION_MOTOR11,
	"ACTUATOR_OUTPUT_FUNCTION_MOTOR12": ACTUATOR_OUTPUT_FUNCTION_MOTOR12,
	"ACTUATOR_OUTPUT_FUNCTION_MOTOR13": ACTUATOR_OUTPUT_FUNCTION_MOTOR13,
	"ACTUATOR_OUTPUT_FUNCTION_MOTOR14": ACTUATOR_OUTPUT_FUNCTION_MOTOR14,
	"ACTUATOR_OUTPUT_FUNCTION_MOTOR15": ACTUATOR_OUTPUT_FUNCTION_MOTOR15,
	"ACTUATOR_OUTPUT_FUNCTION_MOTOR16": ACTUATOR_OUTPUT_FUNCTION_MOTOR16,
	"ACTUATOR_OUTPUT_FUNCTION_SERVO1":  ACTUATOR_OUTPUT_FUNCTION_SERVO1,
	"ACTUATOR_OUTPUT_FUNCTION_SERVO2":  ACTUATOR_OUTPUT_FUNCTION_SERVO2,
	"ACTUATOR_OUTPUT_FUNCTION_SERVO3":  ACTUATOR_OUTPUT_FUNCTION_SERVO3,
	"ACTUATOR_OUTPUT_FUNCTION_SERVO4":  ACTUATOR_OUTPUT_FUNCTION_SERVO4,
	"ACTUATOR_OUTPUT_FUNCTION_SERVO5":  ACTUATOR_OUTPUT_FUNCTION_SERVO5,
	"ACTUATOR_OUTPUT_FUNCTION_SERVO6":  ACTUATOR_OUTPUT_FUNCTION_SERVO6,
	"ACTUATOR_OUTPUT_FUNCTION_SERVO7":  ACTUATOR_OUTPUT_FUNCTION_SERVO7,
	"ACTUATOR_OUTPUT_FUNCTION_SERVO8":  ACTUATOR_OUTPUT_FUNCTION_SERVO8,
	"ACTUATOR_OUTPUT_FUNCTION_SERVO9":  ACTUATOR_OUTPUT_FUNCTION_SERVO9,
	"ACTUATOR_OUTPUT_FUNCTION_SERVO10": ACTUATOR_OUTPUT_FUNCTION_SERVO10,
	"ACTUATOR_OUTPUT_FUNCTION_SERVO11": ACTUATOR_OUTPUT_FUNCTION_SERVO11,
	"ACTUATOR_OUTPUT_FUNCTION_SERVO12": ACTUATOR_OUTPUT_FUNCTION_SERVO12,
	"ACTUATOR_OUTPUT_FUNCTION_SERVO13": ACTUATOR_OUTPUT_FUNCTION_SERVO13,
	"ACTUATOR_OUTPUT_FUNCTION_SERVO14": ACTUATOR_OUTPUT_FUNCTION_SERVO14,
	"ACTUATOR_OUTPUT_FUNCTION_SERVO15": ACTUATOR_OUTPUT_FUNCTION_SERVO15,
	"ACTUATOR_OUTPUT_FUNCTION_SERVO16": ACTUATOR_OUTPUT_FUNCTION_SERVO16,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ACTUATOR_OUTPUT_FUNCTION) MarshalText() ([]byte, error) {
	if name, ok := value_to_label_ACTUATOR_OUTPUT_FUNCTION[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ACTUATOR_OUTPUT_FUNCTION) UnmarshalText(text []byte) error {
	if value, ok := label_to_value_ACTUATOR_OUTPUT_FUNCTION[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = ACTUATOR_OUTPUT_FUNCTION(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e ACTUATOR_OUTPUT_FUNCTION) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
