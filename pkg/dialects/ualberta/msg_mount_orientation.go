//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ualberta

// Orientation of a mount
type MessageMountOrientation struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Roll in global frame (set to NaN for invalid).
	Roll float32
	// Pitch in global frame (set to NaN for invalid).
	Pitch float32
	// Yaw relative to vehicle (set to NaN for invalid).
	Yaw float32
	// Yaw in absolute frame relative to Earth's North, north is 0 (set to NaN for invalid).
	YawAbsolute float32 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageMountOrientation) GetID() uint32 {
	return 265
}
