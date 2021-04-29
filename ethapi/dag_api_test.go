package ethapi

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/require"
)

// PublicDAGChainAPI

func TestPublicDAGChainAPI_CurrentEpoch(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicDAGChainAPI(b)
	require.NotPanics(t, func() {
		res := api.CurrentEpoch(ctx)
		require.NotEmpty(t, res)
	})
}

func TestPublicDAGChainAPI_GetConsensusTime(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicDAGChainAPI(b)
	require.NotPanics(t, func() {
		res, err := api.GetConsensusTime(ctx, "1")
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPublicDAGChainAPI_GetEpochStats(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicDAGChainAPI(b)
	require.NotPanics(t, func() {
		res, err := api.GetEpochStats(ctx, rpc.BlockNumber(1))
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPublicDAGChainAPI_GetEvent(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicDAGChainAPI(b)
	require.NotPanics(t, func() {
		res, err := api.GetEvent(ctx, "1", true)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPublicDAGChainAPI_GetEventHeader(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicDAGChainAPI(b)
	require.NotPanics(t, func() {
		res, err := api.GetEventHeader(ctx, "1")
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPublicDAGChainAPI_GetHeads(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicDAGChainAPI(b)
	require.NotPanics(t, func() {
		res, err := api.GetHeads(ctx, rpc.BlockNumber(1))
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}
