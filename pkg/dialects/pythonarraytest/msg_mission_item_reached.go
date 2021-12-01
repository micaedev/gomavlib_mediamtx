//autogenerated:yes
//nolint:golint,misspell,govet,lll
package pythonarraytest

// A certain mission item has been reached. The system will either hold this position (or circle on the orbit) or (if the autocontinue on the WP was set) continue to the next waypoint.
type MessageMissionItemReached struct {
	// Sequence
	Seq uint16
}

// GetID implements the msg.Message interface.
func (*MessageMissionItemReached) GetID() uint32 {
	return 46
}
