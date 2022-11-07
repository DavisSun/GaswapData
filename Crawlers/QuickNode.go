package Crawlers

import (
	common "GaswapData/Common"
	"context"
	ethereumCommon "github.com/ethereum/go-ethereum/common"
	"log"
)

func GetLatestBlockNumber() uint64 {
	client := common.Ethclient()
	blockNumber, err := client.BlockNumber(context.TODO())
	if err != nil {
		panic(err)
	}
	return blockNumber
}

func GetNextBlockBaseGasFee() uint64 {
	client := common.Ethclient()
	latestBlock, err := client.BlockByNumber(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	baseGasFee := latestBlock.Header().BaseFee
	//gasLimit := latestBlock.Header().GasLimit
	//gasUsed := latestBlock.Header().GasUsed

	return baseGasFee.Uint64()
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

func GetTxPoolContent() {
	println("Start")
	client := common.RPCClient()
	//defer client.Close()

	type resultx struct {
		pending []string `json:"pending"`
		queued  []string `json:"queued"`
	}

	type test struct {
		jsonrpc string
		id      int
		result  resultx
	}
	//var result string
	var result2 resultx

	//type request struct {
	//	id int,
	//	jsonrpc string,
	//}
	if err := client.Call(&result2, "txpool_content"); err != nil {
		println("Failed")
		log.Fatal(err)
	}

	for i := 0; i < len(result2.pending); i += 1 {
		tran := result2.pending[i]

		println(tran)
	}

	//return &result
}

//func GetTxPoolTest() {
//	type message struct {
//		Version string          `json:"jsonrpc,omitempty"`
//		ID      int             `json:"id,omitempty"`
//		Method  string          `json:"method,omitempty"`
//		Params  json.RawMessage `json:"params,omitempty"`
//		Result  json.RawMessage `json:"result,omitempty"`
//	}
//
//	client := httpNative.Client{
//		Timeout: 5 * time.Minute,
//	}
//
//}
