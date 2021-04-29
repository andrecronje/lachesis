package ethapi

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-lachesis/evmcore"
)

// PublicDebugAPI

func TestPublicDebugAPI_GetBlockRlp(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicDebugAPI(b)
	require.NotPanics(t, func() {
		b.EXPECT().BlockByNumber(gomock.Any(), gomock.Any()).
			Return(&evmcore.EvmBlock{
				EvmHeader: evmcore.EvmHeader{
					Number:     big.NewInt(1),
					Hash:       common.Hash{2},
					ParentHash: common.Hash{3},
					Root:       common.Hash{4},
					TxHash:     common.Hash{5},
					Time:       6,
					Coinbase:   common.Address{7},
					GasLimit:   8,
					GasUsed:    9,
				},
				Transactions: types.Transactions{
					types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
				},
			}, nil).
			Times(1)
		res, err := api.GetBlockRlp(ctx, 1)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPublicDebugAPI_PrintBlock(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicDebugAPI(b)
	require.NotPanics(t, func() {
		b.EXPECT().BlockByNumber(gomock.Any(), gomock.Any()).
			Return(&evmcore.EvmBlock{
				EvmHeader: evmcore.EvmHeader{
					Number:     big.NewInt(1),
					Hash:       common.Hash{2},
					ParentHash: common.Hash{3},
					Root:       common.Hash{4},
					TxHash:     common.Hash{5},
					Time:       6,
					Coinbase:   common.Address{7},
					GasLimit:   8,
					GasUsed:    9,
				},
				Transactions: types.Transactions{
					types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
				},
			}, nil).
			Times(1)
		res, err := api.PrintBlock(ctx, 1)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPublicDebugAPI_SeedHash(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicDebugAPI(b)
	require.NotPanics(t, func() {
		b.EXPECT().BlockByNumber(gomock.Any(), gomock.Any()).
			Return(&evmcore.EvmBlock{
				EvmHeader: evmcore.EvmHeader{
					Number:     big.NewInt(1),
					Hash:       common.Hash{2},
					ParentHash: common.Hash{3},
					Root:       common.Hash{4},
					TxHash:     common.Hash{5},
					Time:       6,
					Coinbase:   common.Address{7},
					GasLimit:   8,
					GasUsed:    9,
				},
				Transactions: types.Transactions{
					types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
				},
			}, nil).
			Times(1)
		res, err := api.SeedHash(ctx, 1)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPublicDebugAPI_TestSignCliqueBlock(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicDebugAPI(b)
	require.NotPanics(t, func() {
		res, err := api.TestSignCliqueBlock(ctx, common.Address{}, 1)
		require.Error(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPublicDebugAPI_ValidatorTimeDrifts(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicDebugAPI(b)
	require.NotPanics(t, func() {
		res, err := api.ValidatorTimeDrifts(ctx, 1, 1, 3)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPublicDebugAPI_ValidatorVersions(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicDebugAPI(b)
	require.NotPanics(t, func() {
		b.EXPECT().ForEachEpochEvent(ctx, gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)
		_, err := api.ValidatorVersions(ctx, 1, 1)
		require.NoError(t, err)
	})
}

// PrivateDebugAPI

func TestPrivateDebugAPI_ChaindbProperty(t *testing.T) {
	b := newTestBackend(t)

	api := NewPrivateDebugAPI(b)
	require.NotPanics(t, func() {
		_, _ = api.ChaindbProperty("p1")
	})
}

func TestPrivateDebugAPI_ChaindbCompact(t *testing.T) {
	b := newTestBackend(t)

	api := NewPrivateDebugAPI(b)
	require.NotPanics(t, func() {
		err := api.ChaindbCompact()
		require.NoError(t, err)
	})
}

func TestPrivateDebugAPI_SetHead(t *testing.T) {
	b := newTestBackend(t)

	api := NewPrivateDebugAPI(b)
	require.NotPanics(t, func() {
		err := api.SetHead(1)
		require.Error(t, err)
	})
}
