package ethapi

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-lachesis/evmcore"
	"github.com/Fantom-foundation/go-lachesis/inter"
)

// TODO: for no error
func TestDoCall(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	nonce := hexutil.Uint64(1)
	gas := hexutil.Uint64(0)
	gasPrice := hexutil.Big(*big.NewInt(0))
	value := hexutil.Big(*big.NewInt(0))
	data := hexutil.Bytes([]byte{1, 2, 3})
	code := hexutil.Bytes([]byte{1, 2, 3})
	balance := &hexutil.Big{}

	require.NotPanics(t, func() {
		DoCall(ctx, b, CallArgs{
			Gas:      &gas,
			GasPrice: &gasPrice,
			Value:    &value,
			Data:     &data,
		}, rpc.BlockNumberOrHashWithNumber(1), map[common.Address]account{
			common.HexToAddress("0x0"): account{
				Nonce:   &nonce,
				Code:    &code,
				Balance: &balance,
				StateDiff: &map[common.Hash]common.Hash{
					common.Hash{1}: {1},
				},
			},
		}, vm.Config{}, 100*time.Second, globalGasCap)
		// require.NoError(t, err)
	})

	require.NotPanics(t, func() {
		b.EXPECT().StateAndHeaderByNumberOrHash(gomock.Any(), gomock.Any()).
			Return(nil, nil, ErrBackendTest).
			Times(1)

		DoCall(ctx, b, CallArgs{
			Gas:      &gas,
			GasPrice: &gasPrice,
			Value:    &value,
			Data:     &data,
		}, rpc.BlockNumberOrHashWithNumber(1), map[common.Address]account{
			common.HexToAddress("0x0"): account{
				Nonce:   &nonce,
				Code:    &code,
				Balance: &balance,
				StateDiff: &map[common.Hash]common.Hash{
					common.Hash{1}: {1},
				},
			},
		}, vm.Config{}, 100*time.Second, globalGasCap)
		// require.NoError(t, err)
	})

	require.NotPanics(t, func() {
		initTestBackend(t, b)
		DoCall(ctx, b, CallArgs{
			Gas:      &gas,
			GasPrice: &gasPrice,
			Value:    &value,
			Data:     &data,
		}, rpc.BlockNumberOrHashWithNumber(1), map[common.Address]account{
			common.HexToAddress("0x0"): account{
				Nonce:   &nonce,
				Code:    &code,
				Balance: &balance,
				State: &map[common.Hash]common.Hash{
					common.Hash{1}: {1},
				},
				StateDiff: &map[common.Hash]common.Hash{
					common.Hash{1}: {1},
				},
			},
		}, vm.Config{}, 100*time.Second, globalGasCap)
	})

	require.NotPanics(t, func() {
		initTestBackend(t, b)
		DoCall(ctx, b, CallArgs{
			Gas:      &gas,
			GasPrice: &gasPrice,
			Value:    &value,
			Data:     &data,
		}, rpc.BlockNumberOrHashWithNumber(1), map[common.Address]account{
			common.HexToAddress("0x0"): account{
				Nonce:   &nonce,
				Code:    &code,
				Balance: &balance,
				StateDiff: &map[common.Hash]common.Hash{
					common.Hash{1}: {1},
				},
				State: &map[common.Hash]common.Hash{},
			},
		}, vm.Config{}, 100*time.Second, globalGasCap)
	})

	require.NotPanics(t, func() {
		initTestBackend(t, b)
		gas := hexutil.Uint64(1000)
		DoCall(ctx, b, CallArgs{
			Gas:      &gas,
			GasPrice: &gasPrice,
			Value:    &value,
			Data:     &data,
		}, rpc.BlockNumberOrHashWithNumber(1), map[common.Address]account{
			common.HexToAddress("0x0"): account{
				Nonce:   &nonce,
				Code:    &code,
				Balance: &balance,
				StateDiff: &map[common.Hash]common.Hash{
					common.Hash{1}: {1},
				},
			},
		}, vm.Config{}, 100*time.Second, globalGasCap/1000)
	})

	require.NotPanics(t, func() {
		initTestBackend(t, b)
		b.EXPECT().GetEVM(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, nil, ErrBackendTest).
			Times(1)
		DoCall(ctx, b, CallArgs{
			Gas:      &gas,
			GasPrice: &gasPrice,
			Value:    &value,
			Data:     &data,
		}, rpc.BlockNumberOrHashWithNumber(1), map[common.Address]account{
			common.HexToAddress("0x0"): account{
				Nonce:   &nonce,
				Code:    &code,
				Balance: &balance,
				StateDiff: &map[common.Hash]common.Hash{
					common.Hash{1}: {1},
				},
			},
		}, vm.Config{}, 100*time.Second, globalGasCap)
	})

	require.NotPanics(t, func() {
		initTestBackend(t, b)
		DoCall(ctx, b, CallArgs{
			Gas:      &gas,
			GasPrice: &gasPrice,
			Value:    &value,
			Data:     &data,
		}, rpc.BlockNumberOrHashWithNumber(1), map[common.Address]account{
			common.HexToAddress("0x0"): account{
				Nonce:   &nonce,
				Code:    &code,
				Balance: &balance,
				State: &map[common.Hash]common.Hash{
					common.Hash{1}: {1},
				},
			},
		}, vm.Config{}, 100*time.Second, globalGasCap)
	})
}

// TODO: for no error
func TestDoEstimateGas(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	require.NotPanics(t, func() {
		gas := hexutil.Uint64(21000)
		DoEstimateGas(ctx, b, CallArgs{Gas: &gas}, rpc.BlockNumberOrHashWithNumber(1), 22000)
		// require.NoError(t, err)
	})
}

func TestFormatLogs(t *testing.T) {
	require.NotPanics(t, func() {
		stack := make([]*big.Int, 0, 1024)
		stack = append(stack, big.NewInt(1))
		mem := make([]byte, 32, 32)
		storage := make(map[common.Hash]common.Hash, 3)
		log := vm.StructLog{
			Pc:            1,
			Op:            2,
			Gas:           3,
			GasCost:       4,
			Memory:        mem,
			MemorySize:    5,
			Stack:         stack,
			Storage:       storage,
			Depth:         6,
			RefundCounter: 7,
			Err:           nil,
		}
		res := FormatLogs([]vm.StructLog{log})
		require.NotEmpty(t, res)
	})
}

func TestRPCMarshalBlock(t *testing.T) {
	require.NotPanics(t, func() {
		res, err := RPCMarshalBlock(&evmcore.EvmBlock{
			EvmHeader: evmcore.EvmHeader{
				Number:     big.NewInt(1),
				Hash:       common.Hash{2},
				ParentHash: common.Hash{1},
				Root:       common.Hash{0},
				TxHash:     common.Hash{1},
				Time:       1,
				Coinbase:   common.Address{0},
				GasLimit:   0,
				GasUsed:    0,
			},
			Transactions: types.Transactions{
				types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
				types.NewTransaction(2, common.Address{2}, big.NewInt(2), 2, big.NewInt(0), []byte{}),
			},
		}, types.Bloom{}, true, true)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestRPCMarshalEvent(t *testing.T) {
	require.NotPanics(t, func() {
		res, err := RPCMarshalEvent(
			&inter.Event{
				EventHeader: inter.EventHeader{
					EventHeaderData: inter.EventHeaderData{
						Version:       0,
						Epoch:         0,
						Seq:           0,
						Frame:         0,
						IsRoot:        false,
						Creator:       0,
						PrevEpochHash: common.Hash{},
						Parents:       nil,
						GasPowerLeft:  inter.GasPowerLeft{},
						GasPowerUsed:  0,
						Lamport:       0,
						ClaimedTime:   0,
						MedianTime:    0,
						TxHash:        common.Hash{},
						Extra:         nil,
					},
					Sig: nil,
				},
				Transactions: types.Transactions{
					types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
					types.NewTransaction(2, common.Address{2}, big.NewInt(2), 2, big.NewInt(0), []byte{}),
				},
			},
			true, true,
		)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestRPCMarshalEventHeader(t *testing.T) {
	require.NotPanics(t, func() {
		res := RPCMarshalEventHeader(
			&inter.EventHeaderData{
				Version:       0,
				Epoch:         0,
				Seq:           0,
				Frame:         0,
				IsRoot:        false,
				Creator:       0,
				PrevEpochHash: common.Hash{},
				Parents:       nil,
				GasPowerLeft:  inter.GasPowerLeft{},
				GasPowerUsed:  0,
				Lamport:       0,
				ClaimedTime:   0,
				MedianTime:    0,
				TxHash:        common.Hash{},
				Extra:         nil,
			},
		)
		require.NotEmpty(t, res)
	})
}

func TestRPCMarshalHeader(t *testing.T) {
	require.NotPanics(t, func() {
		res := RPCMarshalHeader(
			&evmcore.EvmHeader{
				Number:     big.NewInt(1),
				Hash:       common.Hash{2},
				ParentHash: common.Hash{3},
				Root:       common.Hash{4},
				TxHash:     common.Hash{5},
				Time:       6,
				Coinbase:   common.Address{7},
				GasLimit:   8,
				GasUsed:    9,
			}, types.Bloom{},
		)
		require.NotEmpty(t, res)
	})
}

func TestSendTxArgs_setDefaults(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		args := SendTxArgs{
			From:     common.Address{1},
			To:       &common.Address{2},
			Gas:      &gas,
			GasPrice: nil,
			Value:    nil,
			Nonce:    nil,
			Data:     nil,
			Input:    nil,
		}
		b.EXPECT().GetPoolNonce(ctx, gomock.Any()).
			Return(uint64(1), nil).
			Times(1)

		err := args.setDefaults(ctx, b)
		require.NoError(t, err)
	})

	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		data := hexutil.Bytes([]byte{1, 2, 3})
		args := SendTxArgs{
			From:     common.Address{1},
			To:       nil,
			Gas:      &gas,
			GasPrice: nil,
			Value:    nil,
			Nonce:    nil,
			Data:     &data,
			Input:    nil,
		}
		b.EXPECT().GetPoolNonce(ctx, gomock.Any()).
			Return(uint64(1), nil).
			Times(1)

		err := args.setDefaults(ctx, b)
		require.NoError(t, err)
	})

	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		data := hexutil.Bytes([]byte{1, 2, 3})
		args := SendTxArgs{
			From:     common.Address{1},
			To:       nil,
			Gas:      &gas,
			GasPrice: nil,
			Value:    nil,
			Nonce:    nil,
			Data:     nil,
			Input:    &data,
		}

		b.EXPECT().GetPoolNonce(ctx, gomock.Any()).
			Return(uint64(1), nil).
			Times(1)
		err := args.setDefaults(ctx, b)
		require.NoError(t, err)
	})

	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		data := hexutil.Bytes([]byte{})
		args := SendTxArgs{
			From:     common.Address{1},
			To:       nil,
			Gas:      &gas,
			GasPrice: nil,
			Value:    nil,
			Nonce:    nil,
			Data:     nil,
			Input:    &data,
		}

		b.EXPECT().GetPoolNonce(ctx, gomock.Any()).
			Return(uint64(1), nil).
			Times(1)

		err := args.setDefaults(ctx, b)
		require.Error(t, err)
	})

	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		data := hexutil.Bytes([]byte{1, 2, 3})
		input := hexutil.Bytes([]byte{3, 2, 1})
		args := SendTxArgs{
			From:     common.Address{1},
			To:       nil,
			Gas:      &gas,
			GasPrice: nil,
			Value:    nil,
			Nonce:    nil,
			Data:     &data,
			Input:    &input,
		}

		b.EXPECT().GetPoolNonce(ctx, gomock.Any()).
			Return(uint64(1), nil).
			Times(1)

		err := args.setDefaults(ctx, b)
		require.Error(t, err)
	})

	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		args := SendTxArgs{
			From:     common.Address{1},
			To:       &common.Address{2},
			Gas:      &gas,
			GasPrice: nil,
			Value:    nil,
			Nonce:    nil,
			Data:     nil,
			Input:    nil,
		}
		b.EXPECT().GetPoolNonce(ctx, gomock.Any()).
			Return(uint64(1), ErrBackendTest).
			Times(1)

		err := args.setDefaults(ctx, b)
		require.Error(t, err)
	})

	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		args := SendTxArgs{
			From:     common.Address{1},
			To:       &common.Address{2},
			Gas:      &gas,
			GasPrice: nil,
			Value:    nil,
			Nonce:    nil,
			Data:     nil,
			Input:    nil,
		}

		b.EXPECT().GetPoolNonce(ctx, gomock.Any()).
			Return(uint64(1), ErrBackendTest).
			Times(1)
		err := args.setDefaults(ctx, b)
		require.Error(t, err)
	})

	require.NotPanics(t, func() {
		initTestBackend(t, b)
		data := hexutil.Bytes([]byte{1, 2, 3})
		args := SendTxArgs{
			From:     common.Address{1},
			To:       &common.Address{2},
			Gas:      nil,
			GasPrice: nil,
			Value:    nil,
			Nonce:    nil,
			Data:     &data,
			Input:    nil,
		}

		b.EXPECT().GetPoolNonce(ctx, gomock.Any()).
			Return(uint64(1), nil).
			Times(1)

		err := args.setDefaults(ctx, b)
		require.Error(t, err)
	})
}

func TestSendTxArgs_toTransaction(t *testing.T) {
	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		nonce := hexutil.Uint64(1)
		value := hexutil.Big(*big.NewInt(1))
		gasPrice := hexutil.Big(*big.NewInt(1))
		args := SendTxArgs{
			From:     common.Address{1},
			To:       &common.Address{2},
			Gas:      &gas,
			GasPrice: &gasPrice,
			Value:    &value,
			Nonce:    &nonce,
			Data:     nil,
			Input:    nil,
		}

		res := args.toTransaction()
		require.NotEmpty(t, res)
	})
	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		nonce := hexutil.Uint64(1)
		value := hexutil.Big(*big.NewInt(1))
		gasPrice := hexutil.Big(*big.NewInt(1))
		input := hexutil.Bytes([]byte{1, 2, 3})
		args := SendTxArgs{
			From:     common.Address{1},
			To:       &common.Address{2},
			Gas:      &gas,
			GasPrice: &gasPrice,
			Value:    &value,
			Nonce:    &nonce,
			Data:     nil,
			Input:    &input,
		}

		res := args.toTransaction()
		require.NotEmpty(t, res)
	})

	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		nonce := hexutil.Uint64(1)
		value := hexutil.Big(*big.NewInt(1))
		gasPrice := hexutil.Big(*big.NewInt(1))
		data := hexutil.Bytes([]byte{1, 2, 3})
		args := SendTxArgs{
			From:     common.Address{1},
			To:       &common.Address{2},
			Gas:      &gas,
			GasPrice: &gasPrice,
			Value:    &value,
			Nonce:    &nonce,
			Data:     &data,
			Input:    nil,
		}

		res := args.toTransaction()
		require.NotEmpty(t, res)
	})

	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		nonce := hexutil.Uint64(1)
		value := hexutil.Big(*big.NewInt(1))
		gasPrice := hexutil.Big(*big.NewInt(1))
		data := hexutil.Bytes([]byte{1, 2, 3})
		args := SendTxArgs{
			From:     common.Address{1},
			To:       nil,
			Gas:      &gas,
			GasPrice: &gasPrice,
			Value:    &value,
			Nonce:    &nonce,
			Data:     &data,
			Input:    nil,
		}

		res := args.toTransaction()
		require.NotEmpty(t, res)
	})
}
