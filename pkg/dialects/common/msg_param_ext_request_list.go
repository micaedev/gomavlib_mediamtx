//autogenerated:yes
//nolint:golint,misspell,govet,lll
package common

// Request all parameters of this component. All parameters should be emitted in response as PARAM_EXT_VALUE.
type MessageParamExtRequestList struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
}

// GetID implements the msg.Message interface.
func (*MessageParamExtRequestList) GetID() uint32 {
	return 321
}
