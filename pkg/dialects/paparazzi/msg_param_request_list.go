//autogenerated:yes
//nolint:golint,misspell,govet,lll
package paparazzi

// Request all parameters of this component. After this request, all parameters are emitted. The parameter microservice is documented at https://mavlink.io/en/services/parameter.html
type MessageParamRequestList struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
}

// GetID implements the msg.Message interface.
func (*MessageParamRequestList) GetID() uint32 {
	return 21
}
