//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// Backwards compatible version of SERIAL_UDB_EXTRA F6: format
type MessageSerialUdbExtraF6 struct {
	// Serial UDB Extra PITCHGAIN Proportional Control
	SuePitchgain float32 `mavname:"sue_PITCHGAIN"`
	// Serial UDB Extra Pitch Rate Control
	SuePitchkd float32 `mavname:"sue_PITCHKD"`
	// Serial UDB Extra Rudder to Elevator Mix
	SueRudderElevMix float32 `mavname:"sue_RUDDER_ELEV_MIX"`
	// Serial UDB Extra Roll to Elevator Mix
	SueRollElevMix float32 `mavname:"sue_ROLL_ELEV_MIX"`
	// Gain For Boosting Manual Elevator control When Plane Stabilized
	SueElevatorBoost float32 `mavname:"sue_ELEVATOR_BOOST"`
}

// GetID implements the message.Message interface.
func (*MessageSerialUdbExtraF6) GetID() uint32 {
	return 174
}
