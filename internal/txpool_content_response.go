package internal

type TxpoolContentResponse struct {
	Pending map[string]map[int]*RPCTransaction `json:"pending"`
	Queued  map[string]map[int]*RPCTransaction `json:"queued"`
}
