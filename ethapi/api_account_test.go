package ethapi

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// PublicAccountAPI

func TestPublicAccountAPI_Accounts(t *testing.T) {
	b := newTestBackend(t)

	api := NewPublicAccountAPI(
		b.AccountManager())
	require.NotPanics(t, func() {
		api.Accounts()
	})
}

// PrivateAccountAPI

func TestPrivateAccountAPI_DeriveAccount(t *testing.T) {
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPrivateAccountAPI(b, nonceLock)
	require.NotPanics(t, func() {
		api.DeriveAccount("https://test.ru", "/test", nil)
	})
}

func TestPrivateAccountAPI_ImportRawKey(t *testing.T) {
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPrivateAccountAPI(b, nonceLock)

	require.NotPanics(t, func() {
		res, err := api.ImportRawKey("11223344556677889900aabbccddff0011223344556677889900aabbccddff00", "1234")
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPrivateAccountAPI_ListAccounts(t *testing.T) {
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPrivateAccountAPI(b, nonceLock)

	require.NotPanics(t, func() {
		api.ListAccounts()
	})
}

func TestPrivateAccountAPI_ListWallets(t *testing.T) {
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPrivateAccountAPI(b, nonceLock)

	require.NotPanics(t, func() {
		api.ListWallets()
	})
}

func TestPrivateAccountAPI_NewAccount(t *testing.T) {
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPrivateAccountAPI(b, nonceLock)

	require.NotPanics(t, func() {
		res, err := api.NewAccount("1234")
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPrivateAccountAPI_UnlockAccount(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPrivateAccountAPI(b, nonceLock)

	addr, _ := api.NewAccount("1234")
	api.ImportRawKey("11223344556677889900aabbccddff0011223344556677889900aabbccddff00", "1234")

	require.NotPanics(t, func() {
		d := uint64(1)
		res, err := api.UnlockAccount(ctx, addr, "1234", &d)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
	require.NotPanics(t, func() {
		res, err := api.UnlockAccount(ctx, addr, "1234", nil)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
	require.NotPanics(t, func() {
		d := uint64(time.Duration(math.MaxInt64)/time.Second + 1)
		_, err := api.UnlockAccount(ctx, addr, "1234", &d)
		require.Error(t, err)
	})
}

func TestPrivateAccountAPI_LockAccount(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPrivateAccountAPI(b, nonceLock)

	addr, _ := api.NewAccount("1234")
	_, _ = api.ImportRawKey("11223344556677889900aabbccddff0011223344556677889900aabbccddff00", "1234")
	d := uint64(1)
	_, _ = api.UnlockAccount(ctx, addr, "1234", &d)

	require.NotPanics(t, func() {
		api.LockAccount(addr)
	})
}

func TestPrivateAccountAPI_Sign(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPrivateAccountAPI(b, nonceLock)

	addr, _ := api.NewAccount("1234")
	key, _ := api.ImportRawKey("11223344556677889900aabbccddff0011223344556677889900aabbccddff00", "1234")
	d := uint64(1)
	_, _ = api.UnlockAccount(ctx, key, "1234", &d)

	require.NotPanics(t, func() {
		res, err := api.Sign(ctx, hexutil.Bytes([]byte{1, 2, 3}), addr, "1234")
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPrivateAccountAPI_SignTransaction(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPrivateAccountAPI(b, nonceLock)

	addr, _ := api.NewAccount("1234")
	key, _ := api.ImportRawKey("11223344556677889900aabbccddff0011223344556677889900aabbccddff00", "1234")
	d := uint64(1)
	_, _ = api.UnlockAccount(ctx, key, "1234", &d)

	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		gasPrice := hexutil.Big(*big.NewInt(0))
		nonce := hexutil.Uint64(1)
		res, err := api.SignTransaction(ctx, SendTxArgs{
			From:     addr,
			To:       &common.Address{1},
			Gas:      &gas,
			GasPrice: &gasPrice,
			Nonce:    &nonce,
		}, "1234")
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
	require.NotPanics(t, func() {
		gasPrice := hexutil.Big(*big.NewInt(0))
		nonce := hexutil.Uint64(1)
		_, err := api.SignTransaction(ctx, SendTxArgs{
			From:     addr,
			To:       &common.Address{1},
			Gas:      nil,
			GasPrice: &gasPrice,
			Nonce:    &nonce,
		}, "1234")
		require.Error(t, err)
	})
	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		nonce := hexutil.Uint64(1)
		_, err := api.SignTransaction(ctx, SendTxArgs{
			From:     addr,
			To:       &common.Address{1},
			Gas:      &gas,
			GasPrice: nil,
			Nonce:    &nonce,
		}, "1234")
		require.Error(t, err)
	})
	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		gasPrice := hexutil.Big(*big.NewInt(0))
		_, err := api.SignTransaction(ctx, SendTxArgs{
			From:     addr,
			To:       &common.Address{1},
			Gas:      &gas,
			GasPrice: &gasPrice,
			Nonce:    nil,
		}, "1234")
		require.Error(t, err)
	})
}

func TestPrivateAccountAPI_SignAndSendTransaction(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPrivateAccountAPI(b, nonceLock)

	addr, _ := api.NewAccount("1234")
	key, _ := api.ImportRawKey("11223344556677889900aabbccddff0011223344556677889900aabbccddff00", "1234")
	d := uint64(1)
	_, _ = api.UnlockAccount(ctx, key, "1234", &d)

	b.EXPECT().SendTx(ctx, gomock.Any()).
		Return(nil).
		Times(1)

	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		gasPrice := hexutil.Big(*big.NewInt(0))
		nonce := hexutil.Uint64(1)
		res, err := api.SignAndSendTransaction(ctx, SendTxArgs{
			From:     addr,
			To:       &common.Address{1},
			Gas:      &gas,
			GasPrice: &gasPrice,
			Nonce:    &nonce,
		}, "1234")
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPrivateAccountAPI_SendTransaction(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPrivateAccountAPI(b, nonceLock)

	addr, _ := api.NewAccount("1234")
	key, _ := api.ImportRawKey("11223344556677889900aabbccddff0011223344556677889900aabbccddff00", "1234")
	d := uint64(1)
	_, _ = api.UnlockAccount(ctx, key, "1234", &d)

	b.EXPECT().GetPoolNonce(ctx, gomock.Any()).
		Return(uint64(1), nil).
		AnyTimes()
	b.EXPECT().SendTx(ctx, gomock.Any()).
		Return(nil).
		Times(2) // TODO: why 2 ?

	require.NotPanics(t, func() {
		gas := hexutil.Uint64(0)
		gasPrice := hexutil.Big(*big.NewInt(0))
		res, err := api.SendTransaction(ctx, SendTxArgs{
			From:     addr,
			To:       &common.Address{1},
			Gas:      &gas,
			GasPrice: &gasPrice,
		}, "1234")
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})
}

func TestPrivateAccountAPI_EcRecover(t *testing.T) {
	ctx := context.TODO()
	b := newTestBackend(t)

	nonceLock := new(AddrLocker)
	api := NewPrivateAccountAPI(b, nonceLock)

	require.NotPanics(t, func() {
		sig := hexutil.Bytes([]byte{})
		data := hexutil.Bytes([]byte{})
		_, err := api.EcRecover(ctx, data, sig)
		require.Error(t, err)
	})
	require.NotPanics(t, func() {
		sig := hexutil.Bytes(make([]byte, crypto.SignatureLength, crypto.SignatureLength))
		data := hexutil.Bytes([]byte{})
		_, err := api.EcRecover(ctx, data, sig)
		require.Error(t, err)
	})
	require.NotPanics(t, func() {
		sig := hexutil.Bytes(make([]byte, crypto.SignatureLength, crypto.SignatureLength))
		sig[crypto.RecoveryIDOffset] = 27
		data := hexutil.Bytes([]byte{})
		_, err := api.EcRecover(ctx, data, sig)
		require.Error(t, err)
	})
}
