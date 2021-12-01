//autogenerated:yes
//nolint:golint,misspell,govet,lll
package paparazzi

// Request to read the value of a parameter with either the param_id string id or param_index. PARAM_EXT_VALUE should be emitted in response.
type MessageParamExtRequestRead struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamId string `mavlen:"16"`
	// Parameter index. Set to -1 to use the Parameter ID field as identifier (else param_id will be ignored)
	ParamIndex int16
}

// GetID implements the msg.Message interface.
func (*MessageParamExtRequestRead) GetID() uint32 {
	return 320
}
