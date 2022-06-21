//autogenerated:yes
//nolint:revive,misspell,govet,lll
package pythonarraytest

// Request all parameters of this component. After this request, all parameters are emitted. The parameter microservice is documented at https://mavlink.io/en/services/parameter.html
type MessageParamRequestList struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
}

// GetID implements the message.Message interface.
func (*MessageParamRequestList) GetID() uint32 {
	return 21
}
