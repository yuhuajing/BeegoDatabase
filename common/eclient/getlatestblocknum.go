package eclient

import (
	"context"
	"main/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetLatestBlockNum() int64 {
	client, _ := ethclient.Dial(config.Ethserver)
	blockNum, _ := client.BlockNumber(context.Background())
	return int64(blockNum)
}
