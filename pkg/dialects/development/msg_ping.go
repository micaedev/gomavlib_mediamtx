//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

// A ping message either requesting or responding to a ping. This allows to measure the system latencies, including serial port, radio modem and UDP connections. The ping microservice is documented at https://mavlink.io/en/services/ping.html
type MessagePing struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// PING sequence
	Seq uint32
	// 0: request ping from all receiving systems. If greater than 0: message is a ping response and number is the system id of the requesting system
	TargetSystem uint8
	// 0: request ping from all receiving components. If greater than 0: message is a ping response and number is the component id of the requesting component.
	TargetComponent uint8
}

// GetID implements the message.Message interface.
func (*MessagePing) GetID() uint32 {
	return 4
}
