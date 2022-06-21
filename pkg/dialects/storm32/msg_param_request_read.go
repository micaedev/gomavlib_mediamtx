//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Request to read the onboard parameter with the param_id string id. Onboard parameters are stored as key[const char*] -> value[float]. This allows to send a parameter to any other component (such as the GCS) without the need of previous knowledge of possible parameter names. Thus the same GCS can store different parameters for different autopilots. See also https://mavlink.io/en/services/parameter.html for a full documentation of QGroundControl and IMU code.
type MessageParamRequestRead struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamId string `mavlen:"16"`
	// Parameter index. Send -1 to use the param ID field as identifier (else the param id will be ignored)
	ParamIndex int16
}

// GetID implements the message.Message interface.
func (*MessageParamRequestRead) GetID() uint32 {
	return 20
}
