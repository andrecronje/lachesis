package ethapi

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-lachesis/evmcore"
)

// PublicTransactionPoolAPI

func TestPublicTransactionPoolAPI_FillTransaction(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPublicTransactionPoolAPI(b, nonceLock)
	require.NotPanics(t, func() {
		b.EXPECT().GetPoolNonce(ctx, gomock.Any()).
			Return(uint64(1), nil).
			AnyTimes()

		gas := hexutil.Uint64(1)
		res, err := api.FillTransaction(ctx, SendTxArgs{
			From: common.Address{1},
			To:   &common.Address{2},
			Gas:  &gas,
		})
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPublicTransactionPoolAPI_GetBlockTransactionCountByHash(t *testing.T) {
	ctx := context.TODO()
	ctrl := gomock.NewController(t)
	b := NewMockBackend(ctrl)

	nonceLock := new(AddrLocker)
	api := NewPublicTransactionPoolAPI(b, nonceLock)
	require.NotPanics(t, func() {
		b.EXPECT().BlockByHash(gomock.Any(), gomock.Any()).
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

		res := api.GetBlockTransactionCountByHash(ctx, common.Hash{1})
		require.NotNil(t, res)
		require.Equal(t, hexutil.Uint(1), *res)
	})

	require.NotPanics(t, func() {
		b.EXPECT().BlockByHash(gomock.Any(), gomock.Any()).
			Return(nil, nil).
			Times(1)

		res := api.GetBlockTransactionCountByHash(ctx, common.Hash{1})
		require.Nil(t, res)
	})
}

func TestPublicTransactionPoolAPI_GetBlockTransactionCountByNumber(t *testing.T) {
	ctx := context.TODO()
	ctrl := gomock.NewController(t)
	b := NewMockBackend(ctrl)

	nonceLock := new(AddrLocker)
	api := NewPublicTransactionPoolAPI(b, nonceLock)
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
		res := api.GetBlockTransactionCountByNumber(ctx, rpc.BlockNumber(1))
		require.NotNil(t, res)
	})
	require.NotPanics(t, func() {
		b.EXPECT().BlockByNumber(gomock.Any(), gomock.Any()).
			Return(nil, nil).
			Times(1)
		res := api.GetBlockTransactionCountByNumber(ctx, rpc.BlockNumber(1))
		require.Nil(t, res)
	})
}

func TestPublicTransactionPoolAPI_GetRawTransactionByBlockHashAndIndex(t *testing.T) {
	ctx := context.TODO()
	ctrl := gomock.NewController(t)
	b := NewMockBackend(ctrl)

	nonceLock := new(AddrLocker)
	api := NewPublicTransactionPoolAPI(b, nonceLock)
	require.NotPanics(t, func() {
		b.EXPECT().BlockByHash(gomock.Any(), gomock.Any()).
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
		res := api.GetRawTransactionByBlockHashAndIndex(ctx, common.Hash{1}, hexutil.Uint(0))
		require.NotNil(t, res)
	})

	require.NotPanics(t, func() {
		b.EXPECT().BlockByHash(gomock.Any(), gomock.Any()).
			Return(nil, nil).
			Times(1)
		res := api.GetRawTransactionByBlockHashAndIndex(ctx, common.Hash{1}, hexutil.Uint(0))
		require.Nil(t, res)
	})
}

func TestPublicTransactionPoolAPI_GetRawTransactionByBlockNumberAndIndex(t *testing.T) {
	ctx := context.TODO()
	ctrl := gomock.NewController(t)
	b := NewMockBackend(ctrl)

	nonceLock := new(AddrLocker)
	api := NewPublicTransactionPoolAPI(b, nonceLock)
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
		res := api.GetRawTransactionByBlockNumberAndIndex(ctx, rpc.BlockNumber(1), hexutil.Uint(0))
		require.NotNil(t, res)
	})

	require.NotPanics(t, func() {
		b.EXPECT().BlockByNumber(gomock.Any(), gomock.Any()).
			Return(nil, nil).
			Times(1)
		res := api.GetRawTransactionByBlockNumberAndIndex(ctx, rpc.BlockNumber(1), hexutil.Uint(0))
		require.Nil(t, res)
	})
}

func TestPublicTransactionPoolAPI_GetRawTransactionByHash(t *testing.T) {
	ctx := context.TODO()
	ctrl := gomock.NewController(t)
	b := NewMockBackend(ctrl)

	nonceLock := new(AddrLocker)
	api := NewPublicTransactionPoolAPI(b, nonceLock)

	require.NotPanics(t, func() {
		b.EXPECT().GetTransaction(ctx, gomock.Any()).
			Return(
				types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
				uint64(1), uint64(1), nil,
			).
			Times(1)
		b.EXPECT().GetPoolTransaction(gomock.Any()).Return(
			types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
		).
			Times(1)
		res, err := api.GetRawTransactionByHash(ctx, common.Hash{1})
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	require.NotPanics(t, func() {
		b.EXPECT().GetTransaction(ctx, gomock.Any()).
			Return(nil, uint64(0), uint64(0), nil).
			Times(1)
		b.EXPECT().GetPoolTransaction(gomock.Any()).Return(
			types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
		).
			Times(1)
		res, err := api.GetRawTransactionByHash(ctx, common.Hash{1})
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	require.NotPanics(t, func() {
		b.EXPECT().GetTransaction(ctx, gomock.Any()).
			Return(nil, uint64(0), uint64(0), ErrBackendTest).
			Times(1)
		b.EXPECT().GetPoolTransaction(gomock.Any()).
			Return(nil).
			Times(1)
		res, err := api.GetRawTransactionByHash(ctx, common.Hash{1})
		require.Error(t, err)
		require.Empty(t, res)
	})
}

func TestPublicTransactionPoolAPI_GetTransactionByBlockHashAndIndex(t *testing.T) {
	ctx := context.TODO()
	ctrl := gomock.NewController(t)
	b := NewMockBackend(ctrl)

	nonceLock := new(AddrLocker)
	api := NewPublicTransactionPoolAPI(b, nonceLock)
	require.NotPanics(t, func() {
		b.EXPECT().BlockByHash(ctx, gomock.Any()).
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

		res := api.GetTransactionByBlockHashAndIndex(ctx, common.Hash{1}, hexutil.Uint(0))
		require.NotEmpty(t, res)
	})

	require.NotPanics(t, func() {
		b.EXPECT().BlockByHash(gomock.Any(), gomock.Any()).
			Return(nil, nil).
			Times(1)
		res := api.GetTransactionByBlockHashAndIndex(ctx, common.Hash{1}, hexutil.Uint(0))
		require.Empty(t, res)
	})
}

func TestPublicTransactionPoolAPI_GetTransactionByBlockNumberAndIndex(t *testing.T) {
	ctx := context.TODO()
	ctrl := gomock.NewController(t)
	b := NewMockBackend(ctrl)

	nonceLock := new(AddrLocker)
	api := NewPublicTransactionPoolAPI(b, nonceLock)

	require.NotPanics(t, func() {
		b.EXPECT().BlockByNumber(ctx, gomock.Any()).
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

		res := api.GetTransactionByBlockNumberAndIndex(ctx, rpc.BlockNumber(1), hexutil.Uint(0))
		require.NotEmpty(t, res)
	})

	require.NotPanics(t, func() {
		b.EXPECT().BlockByNumber(ctx, gomock.Any()).
			Return(nil, nil).
			Times(1)
		res := api.GetTransactionByBlockNumberAndIndex(ctx, rpc.BlockNumber(1), hexutil.Uint(0))
		require.Empty(t, res)
	})
}

func TestPublicTransactionPoolAPI_GetTransactionByHash(t *testing.T) {
	ctx := context.TODO()
	ctrl := gomock.NewController(t)
	b := NewMockBackend(ctrl)

	b.EXPECT().HeaderByNumber(ctx, gomock.Any()).
		Return(&evmcore.EvmHeader{
			Number: big.NewInt(1),
		}, nil).
		AnyTimes()

	nonceLock := new(AddrLocker)
	api := NewPublicTransactionPoolAPI(b, nonceLock)
	require.NotPanics(t, func() {
		b.EXPECT().GetTransaction(ctx, gomock.Any()).
			Return(
				types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
				uint64(1), uint64(1), nil,
			).
			Times(1)
		b.EXPECT().GetPoolTransaction(gomock.Any()).Return(
			types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
		).
			Times(1)

		res, err := api.GetTransactionByHash(ctx, common.Hash{1})
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	require.NotPanics(t, func() {
		b.EXPECT().GetTransaction(ctx, gomock.Any()).
			Return(nil, uint64(0), uint64(0), nil).
			Times(1)
		b.EXPECT().GetPoolTransaction(gomock.Any()).Return(
			types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
		).
			Times(1)
		res, err := api.GetTransactionByHash(ctx, common.Hash{1})
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	require.NotPanics(t, func() {
		b.EXPECT().GetTransaction(ctx, gomock.Any()).
			Return(nil, uint64(0), uint64(0), ErrBackendTest).
			Times(1)
		b.EXPECT().GetPoolTransaction(gomock.Any()).Return(
			nil,
		).
			Times(1)
		res, err := api.GetTransactionByHash(ctx, common.Hash{1})
		require.Error(t, err)
		require.Empty(t, res)
	})
}

func TestPublicTransactionPoolAPI_GetTransactionCount(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPublicTransactionPoolAPI(b, nonceLock)

	b.EXPECT().GetPoolNonce(gomock.Any(), gomock.Any()).
		Return(uint64(1), nil).
		AnyTimes()

	require.NotPanics(t, func() {
		res, err := api.GetTransactionCount(ctx, common.Address{1}, rpc.BlockNumberOrHashWithNumber(1))
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
	require.NotPanics(t, func() {
		res, err := api.GetTransactionCount(ctx, common.Address{1}, rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber))
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
	require.NotPanics(t, func() {
		b.EXPECT().StateAndHeaderByNumberOrHash(gomock.Any(), gomock.Any()).
			Return(nil, nil, ErrBackendTest).
			Times(1)
		res, err := api.GetTransactionCount(ctx, common.Address{1}, rpc.BlockNumberOrHashWithNumber(1))
		require.Error(t, err)
		require.Empty(t, res)
	})
}

func TestPublicTransactionPoolAPI_GetTransactionReceipt(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	setHeaderOn := func() {
		b.EXPECT().HeaderByNumber(gomock.Any(), gomock.Any()).
			Return(&evmcore.EvmHeader{
				Number: big.NewInt(1),
			}, nil).
			Times(1)
	}

	setReceiptsOn := func() {
		rec1 := types.NewReceipt([]byte{}, false, 100)
		rec1.PostState = []byte{1, 2, 3}
		rec1.ContractAddress = common.Address{1}
		rec2 := types.NewReceipt([]byte{}, false, 100)
		b.EXPECT().GetReceiptsByNumber(gomock.Any(), gomock.Any()).
			Return(types.Receipts{
				rec1,
				rec2,
			}, nil).
			Times(1)
	}

	b.EXPECT().GetTransaction(gomock.Any(), gomock.Any()).
		Return(
			types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
			uint64(1), uint64(1), nil,
		).
		Times(3)

	nonceLock := new(AddrLocker)
	api := NewPublicTransactionPoolAPI(b, nonceLock)

	require.NotPanics(t, func() {
		setHeaderOn()
		setReceiptsOn()
		res, err := api.GetTransactionReceipt(ctx, common.Hash{1})
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	require.NotPanics(t, func() {
		setHeaderOn()
		b.EXPECT().GetReceiptsByNumber(ctx, gomock.Any()).
			Return(nil, ErrBackendTest).
			Times(1)
		res, err := api.GetTransactionReceipt(ctx, common.Hash{1})
		require.Error(t, err)
		require.Empty(t, res)
	})

	require.NotPanics(t, func() {
		setReceiptsOn()
		b.EXPECT().HeaderByNumber(ctx, gomock.Any()).
			Return(nil, ErrBackendTest).
			Times(1)

		res, err := api.GetTransactionReceipt(ctx, common.Hash{1})
		require.Error(t, err)
		require.Empty(t, res)
	})

	require.NotPanics(t, func() {
		setHeaderOn()
		b.EXPECT().GetReceiptsByNumber(ctx, gomock.Any()).
			Return(nil, nil).
			Times(1)
		res, err := api.GetTransactionReceipt(ctx, common.Hash{1})
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPublicTransactionPoolAPI_PendingTransactions(t *testing.T) {
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPublicTransactionPoolAPI(b, nonceLock)
	require.NotPanics(t, func() {
		_, err := api.PendingTransactions()
		require.NoError(t, err)
	})
}

func TestPublicTransactionPoolAPI_Resend(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPublicTransactionPoolAPI(b, nonceLock)
	require.NotPanics(t, func() {
		gasPrice := hexutil.Big(*big.NewInt(1))
		gasLimit := hexutil.Uint64(1)
		nonce := hexutil.Uint64(1)
		gas := hexutil.Uint64(0)
		_, _ = api.Resend(ctx, SendTxArgs{
			From:  common.Address{1},
			To:    &common.Address{2},
			Nonce: &nonce,
			Gas:   &gas,
		}, &gasPrice, &gasLimit)
	})
}

func TestPublicTransactionPoolAPI_SignTransaction(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPublicTransactionPoolAPI(b, nonceLock)
	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		gasPrice := hexutil.Big(*big.NewInt(0))
		nonce := hexutil.Uint64(1)
		res, err := api.SignTransaction(ctx, SendTxArgs{
			From:     common.Address{1},
			To:       &common.Address{2},
			Gas:      &gas,
			GasPrice: &gasPrice,
			Nonce:    &nonce,
		})
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
	require.NotPanics(t, func() {
		gasPrice := hexutil.Big(*big.NewInt(0))
		nonce := hexutil.Uint64(1)
		res, err := api.SignTransaction(ctx, SendTxArgs{
			From:     common.Address{1},
			To:       &common.Address{2},
			Gas:      nil,
			GasPrice: &gasPrice,
			Nonce:    &nonce,
		})
		require.Error(t, err)
		require.Empty(t, res)
	})
	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		nonce := hexutil.Uint64(1)
		res, err := api.SignTransaction(ctx, SendTxArgs{
			From:     common.Address{1},
			To:       &common.Address{2},
			Gas:      &gas,
			GasPrice: nil,
			Nonce:    &nonce,
		})
		require.Error(t, err)
		require.Empty(t, res)
	})
	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		gasPrice := hexutil.Big(*big.NewInt(0))
		res, err := api.SignTransaction(ctx, SendTxArgs{
			From:     common.Address{1},
			To:       &common.Address{2},
			Gas:      &gas,
			GasPrice: &gasPrice,
			Nonce:    nil,
		})
		require.Error(t, err)
		require.Empty(t, res)
	})
}

func TestPublicTransactionPoolAPI_SendTransaction(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPublicTransactionPoolAPI(b, nonceLock)

	api2 := NewPrivateAccountAPI(b, nonceLock)

	addr, _ := api2.NewAccount("1234")
	key, _ := api2.ImportRawKey("11223344556677889900aabbccddff0011223344556677889900aabbccddff00", "1234")
	d := uint64(1)
	_, _ = api2.UnlockAccount(ctx, key, "1234", &d)

	b.EXPECT().GetPoolNonce(gomock.Any(), gomock.Any()).
		Return(uint64(1), nil).
		AnyTimes()
	b.EXPECT().SendTx(ctx, gomock.Any()).
		Return(nil).
		Times(2)

	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		gasPrice := hexutil.Big(*big.NewInt(0))
		nonce := hexutil.Uint64(1)
		res, err := api.SendTransaction(ctx, SendTxArgs{
			From:     addr,
			To:       &common.Address{2},
			Gas:      &gas,
			GasPrice: &gasPrice,
			Nonce:    &nonce,
		})
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		gasPrice := hexutil.Big(*big.NewInt(0))
		res, err := api.SendTransaction(ctx, SendTxArgs{
			From:     addr,
			To:       &common.Address{2},
			Gas:      &gas,
			GasPrice: &gasPrice,
			Nonce:    nil,
		})
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPublicTransactionPoolAPI_SendRawTransaction(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	b.EXPECT().SendTx(ctx, gomock.Any()).
		Return(nil).
		Times(1)

	nonceLock := new(AddrLocker)
	api := NewPublicTransactionPoolAPI(b, nonceLock)
	require.NotPanics(t, func() {
		trx := types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{})
		data, _ := rlp.EncodeToBytes(trx)
		res, err := api.SendRawTransaction(ctx, data)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	require.NotPanics(t, func() {
		data := hexutil.Bytes([]byte{1, 2, 3})
		res, err := api.SendRawTransaction(ctx, data)
		require.Error(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPublicTransactionPoolAPI_Sign(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	apiAM := NewPrivateAccountAPI(b, nonceLock)

	addr, _ := apiAM.NewAccount("1234")
	key, _ := apiAM.ImportRawKey("11223344556677889900aabbccddff0011223344556677889900aabbccddff00", "1234")
	d := uint64(1)
	_, _ = apiAM.UnlockAccount(ctx, key, "1234", &d)

	api := NewPublicTransactionPoolAPI(b, nonceLock)
	require.NotPanics(t, func() {
		trx := types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{})
		data, _ := rlp.EncodeToBytes(trx)
		res, err := api.Sign(addr, data)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}
