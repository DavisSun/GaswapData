package crawlers

import (
	common "GaswapData/common"
	"GaswapData/internal"
	"context"
	"fmt"
	ethereumCommon "github.com/ethereum/go-ethereum/common"
	"log"
	"strconv"
)

func GetLatestBlockNumber() uint64 {
	client := common.Ethclient()
	blockNumber, err := client.BlockNumber(context.TODO())
	if err != nil {
		panic(err)
	}
	return blockNumber
}

func GetNextBlockBaseGasFee() float64 {
	client := common.Ethclient()
	latestBlock, err := client.BlockByNumber(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	baseGasFee := latestBlock.BaseFee()
	gasUsed := latestBlock.GasUsed()
	gasLimit := uint64(30000000)
	gasTarget, _ := strconv.ParseFloat(fmt.Sprintf("%.9f", float64(gasUsed)/float64(gasLimit)), 64)
	gasTarget = (gasTarget - 0.5) * 2
	var nextBaseGas = float64(baseGasFee.Uint64()) * (0.125*gasTarget + 1)
	return nextBaseGas

}

func WatchTxPool() {
	log.Printf("Start")
	client := common.Ethclient()
	gethWssClient := common.Gethclient()
	transctionsCh := make(chan ethereumCommon.Hash, 100)
	_, err := gethWssClient.SubscribePendingTransactions(context.Background(), transctionsCh)
	if err != nil {
		log.Printf("failed to SubscribePendingTransactions: %v", err)
		return
	}
	for {
		select {
		case hash := <-transctionsCh:
			tx, _, err := client.TransactionByHash(context.Background(), hash)
			if err != nil {
				continue
			}
			data, _ := tx.MarshalJSON()
			log.Printf("tx: %v", string(data))
			println("tx: %v", string(data))
		}
	}
}

func GetTxPoolContent() *internal.TxpoolContentResponse {
	println("Start")
	client := common.RPCClient()
	defer client.Close()
	var result internal.TxpoolContentResponse
	if err := client.Call(&result, "txpool_content"); err != nil {
		println("[GetTxPoolContent] Failed")
		log.Fatal(err)
	}
	return &result
}
