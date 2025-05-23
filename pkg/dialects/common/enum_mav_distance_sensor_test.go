//autogenerated:yes
//nolint:revive,govet,errcheck
package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnum_MAV_DISTANCE_SENSOR(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var e MAV_DISTANCE_SENSOR
		e.UnmarshalText([]byte{})
		e.MarshalText()
		e.String()
	})

	t.Run("first entry", func(t *testing.T) {
		enc, err := MAV_DISTANCE_SENSOR_LASER.MarshalText()
		require.NoError(t, err)

		var dec MAV_DISTANCE_SENSOR
		err = dec.UnmarshalText(enc)
		require.NoError(t, err)

		require.Equal(t, MAV_DISTANCE_SENSOR_LASER, dec)
	})
}
