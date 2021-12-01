//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ualberta

// Status of geo-fencing. Sent in extended status stream when fencing enabled.
type MessageFenceStatus struct {
	// Breach status (0 if currently inside fence, 1 if outside).
	BreachStatus uint8
	// Number of fence breaches.
	BreachCount uint16
	// Last breach type.
	BreachType FENCE_BREACH `mavenum:"uint8"`
	// Time (since boot) of last breach.
	BreachTime uint32
	// Active action to prevent fence breach
	BreachMitigation FENCE_MITIGATE `mavenum:"uint8" mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageFenceStatus) GetID() uint32 {
	return 162
}
