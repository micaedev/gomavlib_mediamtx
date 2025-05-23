//autogenerated:yes
//nolint:revive,govet,errcheck
package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnum_SET_FOCUS_TYPE(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var e SET_FOCUS_TYPE
		e.UnmarshalText([]byte{})
		e.MarshalText()
		e.String()
	})

	t.Run("first entry", func(t *testing.T) {
		enc, err := FOCUS_TYPE_STEP.MarshalText()
		require.NoError(t, err)

		var dec SET_FOCUS_TYPE
		err = dec.UnmarshalText(enc)
		require.NoError(t, err)

		require.Equal(t, FOCUS_TYPE_STEP, dec)
	})
}
