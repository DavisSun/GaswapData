package internal

type Transaction struct {
	From          string
	Nonce         uint64
	GasPrice      uint64
	FirstGasPrice uint64
}

type Transactions []Transaction

func (trans Transactions) Len() int {
	return len(trans)
}

func (trans Transactions) Less(i, j int) bool {
	if trans[i].FirstGasPrice != trans[j].FirstGasPrice {
		if trans[i].Nonce != trans[j].Nonce {
			return trans[i].Nonce < trans[j].Nonce
		}
		return trans[i].From < trans[j].From
	}
	return trans[i].FirstGasPrice > trans[j].FirstGasPrice
}

func (trans Transactions) Swap(i, j int) {
	trans[i], trans[j] = trans[j], trans[i]
}
