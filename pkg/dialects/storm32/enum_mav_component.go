//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package storm32

import (
	"errors"
)

// Component ids (values) for the different types and instances of onboard hardware/software that might make up a MAVLink system (autopilot, cameras, servos, GPS systems, avoidance systems etc.).
// Components must use the appropriate ID in their source address when sending messages. Components can also use IDs to determine if they are the intended recipient of an incoming message. The MAV_COMP_ID_ALL value is used to indicate messages that must be processed by all components.
// When creating new entries, components that can have multiple instances (e.g. cameras, servos etc.) should be allocated sequential values. An appropriate number of values should be left free after these components to allow the number of instances to be expanded.
type MAV_COMPONENT uint32

const (
	// Target id (target_component) used to broadcast messages to all components of the receiving system. Components should attempt to process messages with this component ID and forward to components on any other interfaces. Note: This is not a valid *source* component id for a message.
	MAV_COMP_ID_ALL MAV_COMPONENT = 0
	// System flight controller component ("autopilot"). Only one autopilot is expected in a particular system.
	MAV_COMP_ID_AUTOPILOT1 MAV_COMPONENT = 1
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER1 MAV_COMPONENT = 25
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER2 MAV_COMPONENT = 26
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER3 MAV_COMPONENT = 27
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER4 MAV_COMPONENT = 28
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER5 MAV_COMPONENT = 29
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER6 MAV_COMPONENT = 30
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER7 MAV_COMPONENT = 31
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER8 MAV_COMPONENT = 32
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER9 MAV_COMPONENT = 33
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER10 MAV_COMPONENT = 34
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER11 MAV_COMPONENT = 35
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER12 MAV_COMPONENT = 36
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER13 MAV_COMPONENT = 37
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER14 MAV_COMPONENT = 38
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER15 MAV_COMPONENT = 39
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER16 MAV_COMPONENT = 40
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER17 MAV_COMPONENT = 41
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER18 MAV_COMPONENT = 42
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER19 MAV_COMPONENT = 43
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER20 MAV_COMPONENT = 44
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER21 MAV_COMPONENT = 45
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER22 MAV_COMPONENT = 46
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER23 MAV_COMPONENT = 47
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER24 MAV_COMPONENT = 48
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER25 MAV_COMPONENT = 49
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER26 MAV_COMPONENT = 50
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER27 MAV_COMPONENT = 51
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER28 MAV_COMPONENT = 52
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER29 MAV_COMPONENT = 53
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER30 MAV_COMPONENT = 54
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER31 MAV_COMPONENT = 55
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER32 MAV_COMPONENT = 56
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER33 MAV_COMPONENT = 57
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER34 MAV_COMPONENT = 58
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER35 MAV_COMPONENT = 59
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER36 MAV_COMPONENT = 60
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER37 MAV_COMPONENT = 61
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER38 MAV_COMPONENT = 62
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER39 MAV_COMPONENT = 63
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER40 MAV_COMPONENT = 64
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER41 MAV_COMPONENT = 65
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER42 MAV_COMPONENT = 66
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER43 MAV_COMPONENT = 67
	// Telemetry radio (e.g. SiK radio, or other component that emits RADIO_STATUS messages).
	MAV_COMP_ID_TELEMETRY_RADIO MAV_COMPONENT = 68
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER45 MAV_COMPONENT = 69
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER46 MAV_COMPONENT = 70
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER47 MAV_COMPONENT = 71
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER48 MAV_COMPONENT = 72
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER49 MAV_COMPONENT = 73
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER50 MAV_COMPONENT = 74
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER51 MAV_COMPONENT = 75
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER52 MAV_COMPONENT = 76
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER53 MAV_COMPONENT = 77
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER54 MAV_COMPONENT = 78
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER55 MAV_COMPONENT = 79
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER56 MAV_COMPONENT = 80
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER57 MAV_COMPONENT = 81
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER58 MAV_COMPONENT = 82
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER59 MAV_COMPONENT = 83
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER60 MAV_COMPONENT = 84
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER61 MAV_COMPONENT = 85
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER62 MAV_COMPONENT = 86
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER63 MAV_COMPONENT = 87
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER64 MAV_COMPONENT = 88
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER65 MAV_COMPONENT = 89
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER66 MAV_COMPONENT = 90
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER67 MAV_COMPONENT = 91
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER68 MAV_COMPONENT = 92
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER69 MAV_COMPONENT = 93
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER70 MAV_COMPONENT = 94
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER71 MAV_COMPONENT = 95
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER72 MAV_COMPONENT = 96
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER73 MAV_COMPONENT = 97
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER74 MAV_COMPONENT = 98
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER75 MAV_COMPONENT = 99
	// Camera #1.
	MAV_COMP_ID_CAMERA MAV_COMPONENT = 100
	// Camera #2.
	MAV_COMP_ID_CAMERA2 MAV_COMPONENT = 101
	// Camera #3.
	MAV_COMP_ID_CAMERA3 MAV_COMPONENT = 102
	// Camera #4.
	MAV_COMP_ID_CAMERA4 MAV_COMPONENT = 103
	// Camera #5.
	MAV_COMP_ID_CAMERA5 MAV_COMPONENT = 104
	// Camera #6.
	MAV_COMP_ID_CAMERA6 MAV_COMPONENT = 105
	// Servo #1.
	MAV_COMP_ID_SERVO1 MAV_COMPONENT = 140
	// Servo #2.
	MAV_COMP_ID_SERVO2 MAV_COMPONENT = 141
	// Servo #3.
	MAV_COMP_ID_SERVO3 MAV_COMPONENT = 142
	// Servo #4.
	MAV_COMP_ID_SERVO4 MAV_COMPONENT = 143
	// Servo #5.
	MAV_COMP_ID_SERVO5 MAV_COMPONENT = 144
	// Servo #6.
	MAV_COMP_ID_SERVO6 MAV_COMPONENT = 145
	// Servo #7.
	MAV_COMP_ID_SERVO7 MAV_COMPONENT = 146
	// Servo #8.
	MAV_COMP_ID_SERVO8 MAV_COMPONENT = 147
	// Servo #9.
	MAV_COMP_ID_SERVO9 MAV_COMPONENT = 148
	// Servo #10.
	MAV_COMP_ID_SERVO10 MAV_COMPONENT = 149
	// Servo #11.
	MAV_COMP_ID_SERVO11 MAV_COMPONENT = 150
	// Servo #12.
	MAV_COMP_ID_SERVO12 MAV_COMPONENT = 151
	// Servo #13.
	MAV_COMP_ID_SERVO13 MAV_COMPONENT = 152
	// Servo #14.
	MAV_COMP_ID_SERVO14 MAV_COMPONENT = 153
	// Gimbal #1.
	MAV_COMP_ID_GIMBAL MAV_COMPONENT = 154
	// Logging component.
	MAV_COMP_ID_LOG MAV_COMPONENT = 155
	// Automatic Dependent Surveillance-Broadcast (ADS-B) component.
	MAV_COMP_ID_ADSB MAV_COMPONENT = 156
	// On Screen Display (OSD) devices for video links.
	MAV_COMP_ID_OSD MAV_COMPONENT = 157
	// Generic autopilot peripheral component ID. Meant for devices that do not implement the parameter microservice.
	MAV_COMP_ID_PERIPHERAL MAV_COMPONENT = 158
	// Gimbal ID for QX1.
	MAV_COMP_ID_QX1_GIMBAL MAV_COMPONENT = 159
	// FLARM collision alert component.
	MAV_COMP_ID_FLARM MAV_COMPONENT = 160
	// Parachute component.
	MAV_COMP_ID_PARACHUTE MAV_COMPONENT = 161
	// Gimbal #2.
	MAV_COMP_ID_GIMBAL2 MAV_COMPONENT = 171
	// Gimbal #3.
	MAV_COMP_ID_GIMBAL3 MAV_COMPONENT = 172
	// Gimbal #4
	MAV_COMP_ID_GIMBAL4 MAV_COMPONENT = 173
	// Gimbal #5.
	MAV_COMP_ID_GIMBAL5 MAV_COMPONENT = 174
	// Gimbal #6.
	MAV_COMP_ID_GIMBAL6 MAV_COMPONENT = 175
	// Battery #1.
	MAV_COMP_ID_BATTERY MAV_COMPONENT = 180
	// Battery #2.
	MAV_COMP_ID_BATTERY2 MAV_COMPONENT = 181
	// CAN over MAVLink client.
	MAV_COMP_ID_MAVCAN MAV_COMPONENT = 189
	// Component that can generate/supply a mission flight plan (e.g. GCS or developer API).
	MAV_COMP_ID_MISSIONPLANNER MAV_COMPONENT = 190
	// Component that lives on the onboard computer (companion computer) and has some generic functionalities, such as settings system parameters and monitoring the status of some processes that don't directly speak mavlink and so on.
	MAV_COMP_ID_ONBOARD_COMPUTER MAV_COMPONENT = 191
	// Component that lives on the onboard computer (companion computer) and has some generic functionalities, such as settings system parameters and monitoring the status of some processes that don't directly speak mavlink and so on.
	MAV_COMP_ID_ONBOARD_COMPUTER2 MAV_COMPONENT = 192
	// Component that lives on the onboard computer (companion computer) and has some generic functionalities, such as settings system parameters and monitoring the status of some processes that don't directly speak mavlink and so on.
	MAV_COMP_ID_ONBOARD_COMPUTER3 MAV_COMPONENT = 193
	// Component that lives on the onboard computer (companion computer) and has some generic functionalities, such as settings system parameters and monitoring the status of some processes that don't directly speak mavlink and so on.
	MAV_COMP_ID_ONBOARD_COMPUTER4 MAV_COMPONENT = 194
	// Component that finds an optimal path between points based on a certain constraint (e.g. minimum snap, shortest path, cost, etc.).
	MAV_COMP_ID_PATHPLANNER MAV_COMPONENT = 195
	// Component that plans a collision free path between two points.
	MAV_COMP_ID_OBSTACLE_AVOIDANCE MAV_COMPONENT = 196
	// Component that provides position estimates using VIO techniques.
	MAV_COMP_ID_VISUAL_INERTIAL_ODOMETRY MAV_COMPONENT = 197
	// Component that manages pairing of vehicle and GCS.
	MAV_COMP_ID_PAIRING_MANAGER MAV_COMPONENT = 198
	// Inertial Measurement Unit (IMU) #1.
	MAV_COMP_ID_IMU MAV_COMPONENT = 200
	// Inertial Measurement Unit (IMU) #2.
	MAV_COMP_ID_IMU_2 MAV_COMPONENT = 201
	// Inertial Measurement Unit (IMU) #3.
	MAV_COMP_ID_IMU_3 MAV_COMPONENT = 202
	// GPS #1.
	MAV_COMP_ID_GPS MAV_COMPONENT = 220
	// GPS #2.
	MAV_COMP_ID_GPS2 MAV_COMPONENT = 221
	// Open Drone ID transmitter/receiver (Bluetooth/WiFi/Internet).
	MAV_COMP_ID_ODID_TXRX_1 MAV_COMPONENT = 236
	// Open Drone ID transmitter/receiver (Bluetooth/WiFi/Internet).
	MAV_COMP_ID_ODID_TXRX_2 MAV_COMPONENT = 237
	// Open Drone ID transmitter/receiver (Bluetooth/WiFi/Internet).
	MAV_COMP_ID_ODID_TXRX_3 MAV_COMPONENT = 238
	// Component to bridge MAVLink to UDP (i.e. from a UART).
	MAV_COMP_ID_UDP_BRIDGE MAV_COMPONENT = 240
	// Component to bridge to UART (i.e. from UDP).
	MAV_COMP_ID_UART_BRIDGE MAV_COMPONENT = 241
	// Component handling TUNNEL messages (e.g. vendor specific GUI of a component).
	MAV_COMP_ID_TUNNEL_NODE MAV_COMPONENT = 242
	// Deprecated, don't use. Component for handling system messages (e.g. to ARM, takeoff, etc.).
	MAV_COMP_ID_SYSTEM_CONTROL MAV_COMPONENT = 250
)

var labels_MAV_COMPONENT = map[MAV_COMPONENT]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_COMPONENT) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_COMPONENT[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_COMPONENT = map[string]MAV_COMPONENT{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_COMPONENT) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_COMPONENT[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_COMPONENT) String() string {
	if l, ok := labels_MAV_COMPONENT[e]; ok {
		return l
	}
	return "invalid value"
}
