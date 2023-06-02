package db

import (

	//"main/common/db"
	"fmt"
	"main/common/eclient"
	"main/common/table"
	"main/config"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	//startBlockHeight, endBlockHeight int
	latestblockNum int64
	endblockNum    = int64(0)
	count          = 0
)

func StoreTx(datab *gorm.DB, address string) {
	datab.AutoMigrate(&table.TxResult{})
	//if endBlockHeight == 0 {
	latestblockNum = eclient.GetLatestBlockNum()
	// } else {
	// 	latestblockNum = int64(endBlockHeight)
	// }
	startblock, _ := eclient.GetTxInfoFromEtherScan(
		address,
		0,
		latestblockNum,
		1,
		100,
		"asc",
		config.ApiKey)
	startblockNum := startblock.Result[0].BlockNumber

	fmt.Printf("startNum:%s\n", startblockNum)

	// deal with transactions less than 100
	if len(startblock.Result) < 100 {
		fmt.Println("less than 100")
		Insert(datab, startblock)
		return
	}
	_startblockNum, _ := strconv.Atoi(startblockNum)

	// if startBlockHeight > _startblockNum {
	// 	_startblockNum = startBlockHeight
	// }

	for _startblockNum < int(latestblockNum) {
		endblockNum = int64(_startblockNum) + 50000
		if endblockNum > latestblockNum {
			endblockNum = latestblockNum
		}

		fmt.Printf("endNum:%d\n", endblockNum)

		lastblock, _ := eclient.GetTxInfoFromEtherScan(
			address,
			int64(_startblockNum),
			endblockNum,
			100,
			100,
			"asc",
			config.ApiKey)
		if len(lastblock.Result) == 100 {
			number, _ := strconv.Atoi(lastblock.Result[99].BlockNumber)
			endblockNum = int64(number) - 1
		}

		if parseTxArray(datab, int64(_startblockNum), endblockNum, address) {
			_startblockNum = int(endblockNum)
		}
		//txHash = append(txHash, res...)
		//fmt.Println(txHash)
	}
	//return txHash
}

func parseTxArray(datab *gorm.DB, start, end int64, address string) bool {
	//temresTxhash := make([]string, 0)
	page := 1
	for page <= 100 {
		transactionsArray, _ := eclient.GetTxInfoFromEtherScan(
			address,
			start,
			end,
			page,
			100,
			"asc",
			config.ApiKey)
		count++
		if count == 5 {
			time.Sleep(1 * time.Second)
			count = 0
		}

		if len(transactionsArray.Result) > 0 {
			Insert(datab, transactionsArray)
			//temresTxhash = append(temresTxhash, resTxhash...)
		}

		if len(transactionsArray.Result) == 100 {
			page += 1
		} else {
			break
		}
	}
	return true
}
