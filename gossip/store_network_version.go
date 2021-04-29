package gossip

import (
	"math/big"
)

const (
	nvKey = "v"
	diKey = "d"
)

// SetNetworkVersion stores network version.
func (s *Store) SetNetworkVersion(v *big.Int) {
	err := s.table.NetworkVersion.Put([]byte(nvKey), v.Bytes())
	if err != nil {
		s.Log.Crit("Failed to put key", "err", err)
	}
}

// GetNetworkVersion returns stored network version.
func (s *Store) GetNetworkVersion() *big.Int {
	valBytes, err := s.table.NetworkVersion.Get([]byte(nvKey))
	if err != nil {
		s.Log.Crit("Failed to get key", "err", err)
	}
	if valBytes == nil {
		return big.NewInt(0)
	}

	return new(big.Int).SetBytes(valBytes)
}

// SetNonSupportedUpgrade stores non-supported network upgrade.
func (s *Store) SetNonSupportedUpgrade(v *big.Int) {
	err := s.table.NetworkVersion.Put([]byte(diKey), v.Bytes())
	if err != nil {
		s.Log.Crit("Failed to put key", "err", err)
	}
}

// GetNonSupportedUpgrade returns stored non-supported network upgrade.
func (s *Store) GetNonSupportedUpgrade() *big.Int {
	valBytes, err := s.table.NetworkVersion.Get([]byte(diKey))
	if err != nil {
		s.Log.Crit("Failed to get key", "err", err)
	}
	if valBytes == nil {
		return big.NewInt(0)
	}

	return new(big.Int).SetBytes(valBytes)
}
