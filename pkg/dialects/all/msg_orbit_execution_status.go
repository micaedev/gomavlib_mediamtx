//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Vehicle status report that is sent out while orbit execution is in progress (see MAV_CMD_DO_ORBIT).
type MessageOrbitExecutionStatus struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Radius of the orbit circle. Positive values orbit clockwise, negative values orbit counter-clockwise.
	Radius float32
	// The coordinate system of the fields: x, y, z.
	Frame MAV_FRAME `mavenum:"uint8"`
	// X coordinate of center point. Coordinate system depends on frame field: local = x position in meters * 1e4, global = latitude in degrees * 1e7.
	X int32
	// Y coordinate of center point.  Coordinate system depends on frame field: local = x position in meters * 1e4, global = latitude in degrees * 1e7.
	Y int32
	// Altitude of center point. Coordinate system depends on frame field.
	Z float32
}

// GetID implements the message.Message interface.
func (*MessageOrbitExecutionStatus) GetID() uint32 {
	return 360
}
