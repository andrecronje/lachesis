package params

type DagConfig struct {
	DagID uint64 `json:"chainId"` // chainId identifies the current chain and is used for replay protection
}