package Common

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	httpNative "net/http"
	"time"
)

const (
	url = "https://quick-crimson-bridge.discover.quiknode.pro/8623d2e79360498dde733bcde5b5fe709d81bdaf/"
	wss = "wss://quick-crimson-bridge.discover.quiknode.pro/8623d2e79360498dde733bcde5b5fe709d81bdaf/"
)

func Ethclient() *ethclient.Client {
	//apiKey := "https://wandering-stylish-crater.discover.quiknode.pro/8a7d831183e9fdf3ca6d8a88af7896a0f9dc1c7e/"
	client, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}
	return client
}

func Gethclient() *gethclient.Client {
	rpcCli, err := rpc.Dial(wss)
	if err != nil {
		log.Printf("failed to dial: %v", err)
		return nil
	}
	client := gethclient.New(rpcCli)
	return client
}

func RPCClient() *rpc.Client {
	httpCLient := httpNative.Client{
		Timeout: 5 * time.Minute,
	}

	rpcCli, err := rpc.DialHTTPWithClient(url, &httpCLient)
	if err != nil {
		log.Printf("failed to dial: %v", err)
		return nil
	}
	return rpcCli
}
