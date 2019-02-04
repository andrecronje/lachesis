package utils

import (
	"encoding/hex"
	"net"
	"strconv"
	"testing"
)

func GetUnusedNetAddr(n int, t testing.TB) []string {
	// addresses 1-1024 are reserved for non-root users;
	// so we start with default lachesis port ang going up until one free found
	idx := int(0)
	addresses := make([]string, n)
	for i := 12000; i < 65536; i++ {
		addrStr := "127.0.0.1:" + strconv.Itoa(i)
		addr, err := net.ResolveTCPAddr("tcp", addrStr)
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		l, err := net.ListenTCP("tcp", addr)
		if err != nil {
			continue
		}
		defer l.Close()
		t.Logf("Unused port %s is chosen", addrStr)
		addresses[idx] = addrStr
		idx++
		if idx == n {
			return addresses
		}
	}
	t.Fatalf("No free port left!!!")
	return addresses
}

// HashFromHex converts hex string to bytes.
func HashFromHex(s string) []byte {
	if len(s) > 1 {
		if s[0:2] == "0x" || s[0:2] == "0X" {
			s = s[2:]
		}
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	h, _ := hex.DecodeString(s)
	return h
}

// FreePort gets free network port on host.
func FreePort(network string) (port uint16) {
	addr, err := net.ResolveTCPAddr(network, "localhost:0")
	if err != nil {
		panic(err)
	}

	l, err := net.ListenTCP(network, addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	return uint16(l.Addr().(*net.TCPAddr).Port)
}
