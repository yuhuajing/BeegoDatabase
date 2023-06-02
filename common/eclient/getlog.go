package eclient

import (
	"context"
	"main/config"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTxlogs(txhash string) []*types.Log {

	nclient, _ := ethclient.Dial(config.Ethserver)
	//https://cool-muddy-butterfly.discover.quiknode.pro/0e41f42d5a7c9611f30ef800444bfcb93d3ae9a6/
	// if err != nil {
	// 	log.Fatal(err)
	// }
	logres := make([]*types.Log, 0)
	receipt, _ := nclient.TransactionReceipt(context.Background(), common.HexToHash(txhash))
	for _, log := range receipt.Logs {
		logres = append(logres, log)
	}
	return logres
}
