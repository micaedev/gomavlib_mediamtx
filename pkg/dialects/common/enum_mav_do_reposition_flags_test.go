//autogenerated:yes
//nolint:revive,govet,errcheck
package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnum_MAV_DO_REPOSITION_FLAGS(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var e MAV_DO_REPOSITION_FLAGS
		e.UnmarshalText([]byte{})
		e.MarshalText()
		e.String()
	})

	t.Run("first entry", func(t *testing.T) {
		enc, err := MAV_DO_REPOSITION_FLAGS_CHANGE_MODE.MarshalText()
		require.NoError(t, err)

		var dec MAV_DO_REPOSITION_FLAGS
		err = dec.UnmarshalText(enc)
		require.NoError(t, err)

		require.Equal(t, MAV_DO_REPOSITION_FLAGS_CHANGE_MODE, dec)
	})
}
