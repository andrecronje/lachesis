package ethapi

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestAddrLocker_LockAddr(t *testing.T) {
	assert.NotPanics(t, func() {
		l := AddrLocker{}
		l.LockAddr(common.Address{1})
		l.UnlockAddr(common.Address{1})
	})
}
