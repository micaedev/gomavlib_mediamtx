//autogenerated:yes
//nolint:golint,misspell,govet,lll
package matrixpilot

// Depreciated but used as a compiler flag.  Do not remove
type MessageFlexifunctionSet struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
}

// GetID implements the msg.Message interface.
func (*MessageFlexifunctionSet) GetID() uint32 {
	return 150
}
