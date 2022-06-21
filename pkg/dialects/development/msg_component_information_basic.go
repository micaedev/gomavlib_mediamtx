//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

// Basic component information data.
type MessageComponentInformationBasic struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Name of the component vendor
	VendorName [32]uint8
	// Name of the component model
	ModelName [32]uint8
	// Sofware version. The version format can be custom, recommended is SEMVER 'major.minor.patch'.
	SoftwareVersion string `mavlen:"24"`
	// Hardware version. The version format can be custom, recommended is SEMVER 'major.minor.patch'.
	HardwareVersion string `mavlen:"24"`
	// Component capability flags
	Capabilities MAV_PROTOCOL_CAPABILITY `mavenum:"uint64"`
}

// GetID implements the message.Message interface.
func (*MessageComponentInformationBasic) GetID() uint32 {
	return 396
}
