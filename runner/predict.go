package runner

import (
	"GaswapData/crawlers"
	"GaswapData/internal"
	"github.com/go-co-op/gocron"
	"log"
	"sort"
	"time"
)

func Start() {
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Every(5).Seconds().Do(func() {
		log.Println("111111111")
	})
	if err != nil {
		panic(err)
	}
	scheduler.StartAsync()
}

func Predict(blockNumber uint64) {
	transacions := convertTxPoolContentToTransactions(crawlers.GetTxPoolContent())
	for _, tx := range transacions {
		println(tx.FirstGasPrice)
	}
}

// Only obtain pending
func convertTxPoolContentToTransactions(response *internal.TxpoolContentResponse) internal.Transactions {
	pending := response.Pending
	result := internal.Transactions{}
	for from, poolTxs := range pending {
		var firstGasPrice uint64
		for nonce, txDetail := range poolTxs {
			if firstGasPrice == 0 {
				firstGasPrice = txDetail.GasPrice.ToInt().Uint64()
			}
			tx := internal.Transaction{
				From:          from,
				Nonce:         nonce,
				GasPrice:      txDetail.GasPrice.ToInt().Uint64(),
				FirstGasPrice: firstGasPrice,
			}
			result = append(result, tx)
		}
	}
	sort.Sort(result)
	return result
}
