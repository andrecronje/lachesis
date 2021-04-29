package ethapi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// PublicTxPoolAPI

func TestPublicTxPoolAPI_Content(t *testing.T) {
	b := newTestBackend(t)

	api := NewPublicTxPoolAPI(b)
	require.NotPanics(t, func() {
		res := api.Content()
		require.NotEmpty(t, res)
	})
}

func TestPublicTxPoolAPI_Status(t *testing.T) {
	b := newTestBackend(t)

	api := NewPublicTxPoolAPI(b)
	require.NotPanics(t, func() {
		res := api.Status()
		require.NotEmpty(t, res)
	})
}

func TestPublicTxPoolAPI_Inspect(t *testing.T) {
	b := newTestBackend(t)

	api := NewPublicTxPoolAPI(b)
	require.NotPanics(t, func() {
		res := api.Inspect()
		require.NotEmpty(t, res)
	})
}
