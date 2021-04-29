package app

type (
	// StoreConfig is a config for store db.
	StoreConfig struct {
		// Cache size for Receipts.
		ReceiptsCacheSize int
		// Cache size for Stakers.
		StakersCacheSize int
		// Cache size for Delegations.
		DelegationsCacheSize int
	}
)

// DefaultStoreConfig for product.
func DefaultStoreConfig() StoreConfig {
	return StoreConfig{
		ReceiptsCacheSize:    0,
		DelegationsCacheSize: 1000,
		StakersCacheSize:     400,
	}
}

// LiteStoreConfig is for tests or inmemory.
func LiteStoreConfig() StoreConfig {
	return StoreConfig{
		ReceiptsCacheSize:    0,
		DelegationsCacheSize: 200,
		StakersCacheSize:     80,
	}
}
