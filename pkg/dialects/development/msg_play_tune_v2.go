//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

// Play vehicle tone/tune (buzzer). Supersedes message PLAY_TUNE.
type MessagePlayTuneV2 struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Tune format
	Format TUNE_FORMAT `mavenum:"uint32"`
	// Tune definition as a NULL-terminated string.
	Tune string `mavlen:"248"`
}

// GetID implements the message.Message interface.
func (*MessagePlayTuneV2) GetID() uint32 {
	return 400
}
