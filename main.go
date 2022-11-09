package main

import (
	"GaswapData/crawlers"
	"GaswapData/runner"
)

func main() {
	runner.Predict(crawlers.GetLatestBlockNumber())

	//crawlers.GetTxPoolContent()
	//println(crawlers.GetLatestBlockNumber())
	//crawlers.WatchTxPool()

	//r := gin.Default()
	//mainRouter.Setup(r)
	//
	//err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//runner.Start()

}
