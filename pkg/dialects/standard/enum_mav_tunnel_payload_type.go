//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package standard

import (
	"errors"
)

type MAV_TUNNEL_PAYLOAD_TYPE uint32

const (
	// Encoding of payload unknown.
	MAV_TUNNEL_PAYLOAD_TYPE_UNKNOWN MAV_TUNNEL_PAYLOAD_TYPE = 0
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED0 MAV_TUNNEL_PAYLOAD_TYPE = 200
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED1 MAV_TUNNEL_PAYLOAD_TYPE = 201
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED2 MAV_TUNNEL_PAYLOAD_TYPE = 202
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED3 MAV_TUNNEL_PAYLOAD_TYPE = 203
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED4 MAV_TUNNEL_PAYLOAD_TYPE = 204
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED5 MAV_TUNNEL_PAYLOAD_TYPE = 205
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED6 MAV_TUNNEL_PAYLOAD_TYPE = 206
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED7 MAV_TUNNEL_PAYLOAD_TYPE = 207
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED8 MAV_TUNNEL_PAYLOAD_TYPE = 208
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED9 MAV_TUNNEL_PAYLOAD_TYPE = 209
)

var labels_MAV_TUNNEL_PAYLOAD_TYPE = map[MAV_TUNNEL_PAYLOAD_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_TUNNEL_PAYLOAD_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_TUNNEL_PAYLOAD_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_TUNNEL_PAYLOAD_TYPE = map[string]MAV_TUNNEL_PAYLOAD_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_TUNNEL_PAYLOAD_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_TUNNEL_PAYLOAD_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_TUNNEL_PAYLOAD_TYPE) String() string {
	if l, ok := labels_MAV_TUNNEL_PAYLOAD_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
