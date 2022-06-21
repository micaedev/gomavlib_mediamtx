//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// Request to control this MAV
type MessageChangeOperatorControl struct {
	// System the GCS requests control for
	TargetSystem uint8
	// 0: request control of this MAV, 1: Release control of this MAV
	ControlRequest uint8
	// 0: key as plaintext, 1-255: future, different hashing/encryption variants. The GCS should in general use the safest mode possible initially and then gradually move down the encryption level if it gets a NACK message indicating an encryption mismatch.
	Version uint8
	// Password / Key, depending on version plaintext or encrypted. 25 or less characters, NULL terminated. The characters may involve A-Z, a-z, 0-9, and "!?,.-"
	Passkey string `mavlen:"25"`
}

// GetID implements the message.Message interface.
func (*MessageChangeOperatorControl) GetID() uint32 {
	return 5
}
