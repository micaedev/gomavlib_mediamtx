//autogenerated:yes
//nolint:revive,govet,errcheck
package loweheiser

import (
	"testing"
)

func TestEnum_MAV_CMD(t *testing.T) {
	var e MAV_CMD
	e.UnmarshalText([]byte{})
	e.MarshalText()
	e.String()
}
