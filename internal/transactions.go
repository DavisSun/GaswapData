package internal

type Transaction struct {
	From          string
	Nonce         uint32
	GasPrice      uint64
	FirstGasPrice uint64
}
