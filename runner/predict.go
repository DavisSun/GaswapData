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

func Predict(blockNumber uint64) []*internal.GasPricePredictionRecord {
	transacions := convertTxPoolContentToTransactions(crawlers.GetTxPoolContent())
	result := make([]*internal.GasPricePredictionRecord, 0)
	var totalGasLimit uint64
	var gasLimit uint64
	blockHeight := blockNumber + 1
	transactionsSumCount := 0
	for _, tx := range transacions {

		println(tx.FirstGasPrice)
	}

}

// Only obtain pending
func convertTxPoolContentToTransactions(response *internal.TxpoolContentResponse, nextBlockBaseFee uint64) internal.Transactions {
	pending := response.Pending
	result := internal.Transactions{}
	for from, poolTxs := range pending {
		var firstGasPrice uint64
		for nonce, txDetail := range poolTxs {
			gasPrice := txDetail.GasTipCap.ToInt().Uint64()

			if txDetail.Type == 2 {
				gasPrice += nextBlockBaseFee

			}
			if firstGasPrice == 0 {
				firstGasPrice = gasPrice
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
