package internal

type TxpoolContentResponse struct {
	Pending map[string]map[uint64]*RPCTransaction `json:"pending"`
	Queued  map[string]map[uint64]*RPCTransaction `json:"queued"`
}
