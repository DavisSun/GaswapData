package runner

import (
	"GaswapData/internal"
	"github.com/go-co-op/gocron"
	"log"
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

}

func convertTxPoolContentToTranscations(response *internal.TxpoolContentResponse) {

}
