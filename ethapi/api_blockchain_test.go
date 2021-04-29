package ethapi

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-lachesis/evmcore"
)

var (
	ErrBackendTest = errors.New("backend test error")
)

// PublicBlockChainAPI

func TestPublicBlockChainAPI_Call(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	gas := hexutil.Uint64(0)
	gasPrice := hexutil.Big(*big.NewInt(0))
	value := hexutil.Big(*big.NewInt(0))
	data := hexutil.Bytes([]byte{1, 2, 3})
	code := hexutil.Bytes([]byte{1, 2, 3})
	balance := &hexutil.Big{}
	nonce := hexutil.Uint64(1)

	api := NewPublicBlockChainAPI(b)
	require.NotPanics(t, func() {
		api.Call(ctx, CallArgs{
			Gas:      &gas,
			GasPrice: &gasPrice,
			Value:    &value,
			Data:     &data,
		}, rpc.BlockNumberOrHashWithNumber(1), &map[common.Address]account{
			common.HexToAddress("0x0"): account{
				Nonce:   &nonce,
				Code:    &code,
				Balance: &balance,
				StateDiff: &map[common.Hash]common.Hash{
					common.Hash{1}: {1},
				},
			},
		})
		// require.NoError(t, err)
	})
}

func TestPublicBlockChainAPI_BlockNumber(t *testing.T) {
	b := newTestBackend(t)

	b.EXPECT().HeaderByNumber(gomock.Any(), gomock.Any()).
		Return(&evmcore.EvmHeader{
			Number: big.NewInt(1),
		}, nil).
		Times(1)

	api := NewPublicBlockChainAPI(b)
	require.NotPanics(t, func() {
		res := api.BlockNumber()
		require.NotEmpty(t, res)
	})
}

func TestPublicBlockChainAPI_ChainID(t *testing.T) {
	b := newTestBackend(t)

	api := NewPublicBlockChainAPI(b)
	require.NotPanics(t, func() {
		res := api.ChainID()
		require.NotEmpty(t, res)
	})
}

func TestPublicBlockChainAPI_EstimateGas(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicBlockChainAPI(b)
	require.NotPanics(t, func() {
		blockNr := rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber)
		_, _ = api.EstimateGas(ctx, CallArgs{}, &blockNr)
	})
}

func TestPublicBlockChainAPI_GetBalance(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicBlockChainAPI(b)
	require.NotPanics(t, func() {
		balance, err := api.GetBalance(ctx, common.Address{1}, rpc.BlockNumberOrHashWithNumber(1))
		require.NoError(t, err)
		require.Equal(t, big.NewInt(10), balance.ToInt())
	})

	require.NotPanics(t, func() {
		b.EXPECT().StateAndHeaderByNumberOrHash(ctx, gomock.Any()).
			Return(nil, nil, ErrBackendTest).
			Times(1)
		_, err := api.GetBalance(ctx, common.Address{1}, rpc.BlockNumberOrHashWithNumber(1))
		require.Error(t, err)
	})
}

func TestPublicBlockChainAPI_GetBlockByHash(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicBlockChainAPI(b)

	require.NotPanics(t, func() {
		b.EXPECT().CalcLogsBloom().
			Return(false).
			Times(1)
		res, err := api.GetBlockByHash(ctx, common.Hash{1}, true)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	require.NotPanics(t, func() {
		b.EXPECT().BlockByHash(gomock.Any(), gomock.Any()).
			Return(nil, ErrBackendTest).
			Times(1)
		b.EXPECT().CalcLogsBloom().
			Return(false).
			Times(1)
		res, err := api.GetBlockByHash(ctx, common.Hash{1}, true)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPublicBlockChainAPI_GetBlockByNumber(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicBlockChainAPI(b)

	require.NotPanics(t, func() {
		b.EXPECT().CalcLogsBloom().
			Return(false).
			Times(1)
		res, err := api.GetBlockByNumber(ctx, rpc.BlockNumber(rpc.PendingBlockNumber), true)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	require.NotPanics(t, func() {
		b.EXPECT().CalcLogsBloom().
			Return(false).
			Times(1)
		b.EXPECT().BlockByNumber(gomock.Any(), gomock.Any()).
			Return(nil, ErrBackendTest).
			Times(1)
		_, err := api.GetBlockByNumber(ctx, rpc.BlockNumber(rpc.PendingBlockNumber), true)
		require.NoError(t, err)
	})
}

func TestPublicBlockChainAPI_GetCode(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicBlockChainAPI(b)

	require.NotPanics(t, func() {
		res, err := api.GetCode(ctx, common.Address{1}, rpc.BlockNumberOrHashWithNumber(1))
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	require.NotPanics(t, func() {
		b.EXPECT().StateAndHeaderByNumberOrHash(ctx, gomock.Any()).
			Return(nil, nil, ErrBackendTest).
			Times(1)
		res, err := api.GetCode(ctx, common.Address{1}, rpc.BlockNumberOrHashWithNumber(1))
		require.Error(t, err)
		require.Empty(t, res)
	})
}

func TestPublicBlockChainAPI_GetHeaderByHash(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicBlockChainAPI(b)
	require.NotPanics(t, func() {
		b.EXPECT().CalcLogsBloom().
			Return(false).
			Times(1)
		b.EXPECT().HeaderByHash(ctx, gomock.Any()).
			Return(&evmcore.EvmHeader{
				Number: big.NewInt(1),
			}, nil).
			Times(1)
		res := api.GetHeaderByHash(ctx, common.HexToHash("0x1"))
		require.NotEmpty(t, res)
	})
	require.NotPanics(t, func() {
		b.EXPECT().CalcLogsBloom().
			Return(false).
			Times(1)
		b.EXPECT().HeaderByHash(ctx, gomock.Any()).
			Return(nil, nil).
			Times(1)
		res := api.GetHeaderByHash(ctx, common.HexToHash("0x1"))
		require.Empty(t, res)
	})
}

func TestPublicBlockChainAPI_GetHeaderByNumber(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicBlockChainAPI(b)
	require.NotPanics(t, func() {
		b.EXPECT().CalcLogsBloom().
			Return(false).
			Times(1)
		b.EXPECT().HeaderByNumber(ctx, gomock.Any()).
			Return(&evmcore.EvmHeader{
				Number: big.NewInt(1),
			}, nil).
			Times(1)
		res, err := api.GetHeaderByNumber(ctx, rpc.BlockNumber(rpc.PendingBlockNumber))
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
	require.NotPanics(t, func() {
		b.EXPECT().CalcLogsBloom().
			Return(false).
			Times(1)
		b.EXPECT().HeaderByNumber(gomock.Any(), gomock.Any()).
			Return(nil, ErrBackendTest).
			Times(1)
		_, err := api.GetHeaderByNumber(ctx, rpc.BlockNumber(rpc.PendingBlockNumber))
		require.Error(t, err)
	})
}

func TestPublicBlockChainAPI_GetProof(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicBlockChainAPI(b)
	require.NotPanics(t, func() {
		res, err := api.GetProof(ctx, common.Address{1}, []string{"1"}, rpc.BlockNumberOrHashWithNumber(1))
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
	require.NotPanics(t, func() {
		b.EXPECT().StateAndHeaderByNumberOrHash(gomock.Any(), gomock.Any()).
			Return(nil, nil, ErrBackendTest).
			Times(1)
		_, err := api.GetProof(ctx, common.Address{1}, []string{"1"}, rpc.BlockNumberOrHashWithNumber(1))
		require.Error(t, err)
	})
}

func TestPublicBlockChainAPI_GetStorageAt(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicBlockChainAPI(b)
	require.NotPanics(t, func() {
		res, err := api.GetStorageAt(ctx, common.Address{1}, "1", rpc.BlockNumberOrHashWithNumber(1))
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
	require.NotPanics(t, func() {
		b.EXPECT().StateAndHeaderByNumberOrHash(gomock.Any(), gomock.Any()).
			Return(nil, nil, ErrBackendTest).
			Times(1)
		res, err := api.GetStorageAt(ctx, common.Address{1}, "1", rpc.BlockNumberOrHashWithNumber(1))
		require.Error(t, err)
		require.Empty(t, res)
	})
}

func TestPublicBlockChainAPI_GetUncleByBlockHashAndIndex(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicBlockChainAPI(b)
	require.NotPanics(t, func() {
		_, err := api.GetUncleByBlockHashAndIndex(ctx, common.Hash{1}, hexutil.Uint(1))
		require.NoError(t, err)
	})

	require.NotPanics(t, func() {
		b.EXPECT().BlockByHash(ctx, gomock.Any()).
			Return(nil, ErrBackendTest).
			Times(1)
		_, err := api.GetUncleByBlockHashAndIndex(ctx, common.Hash{1}, hexutil.Uint(1))
		require.NoError(t, err)
	})
}

func TestPublicBlockChainAPI_GetUncleByBlockNumberAndIndex(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicBlockChainAPI(b)
	require.NotPanics(t, func() {
		_, err := api.GetUncleByBlockNumberAndIndex(ctx, rpc.BlockNumber(1), hexutil.Uint(1))
		require.NoError(t, err)
	})
	require.NotPanics(t, func() {
		b.EXPECT().BlockByNumber(ctx, gomock.Any()).
			Return(nil, ErrBackendTest).
			Times(1)
		_, err := api.GetUncleByBlockNumberAndIndex(ctx, rpc.BlockNumber(1), hexutil.Uint(1))
		require.NoError(t, err)
	})
}

func TestPublicBlockChainAPI_GetUncleCountByBlockHash(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicBlockChainAPI(b)

	require.NotPanics(t, func() {
		b.EXPECT().BlockByHash(ctx, gomock.Any()).
			Return(nil, nil).
			Times(1)
		api.GetUncleCountByBlockHash(ctx, common.Hash{1})
	})

	require.NotPanics(t, func() {
		b.EXPECT().BlockByHash(ctx, gomock.Any()).
			Return(nil, ErrBackendTest).
			Times(1)
		api.GetUncleCountByBlockHash(ctx, common.Hash{1})
	})
}

func TestPublicBlockChainAPI_GetUncleCountByBlockNumber(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	api := NewPublicBlockChainAPI(b)

	require.NotPanics(t, func() {
		b.EXPECT().BlockByNumber(gomock.Any(), gomock.Any()).
			Return(nil, nil).
			Times(1)
		api.GetUncleCountByBlockNumber(ctx, rpc.BlockNumber(1))
	})

	require.NotPanics(t, func() {
		b.EXPECT().BlockByNumber(gomock.Any(), gomock.Any()).
			Return(nil, ErrBackendTest).
			Times(1)
		api.GetUncleCountByBlockNumber(ctx, rpc.BlockNumber(1))
	})
}
