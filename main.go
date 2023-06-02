package main

import (
	"main/common/db"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// flag.StringVar(&address, "address", "", "Get all transactions of this address")
	// flag.IntVar(&startBlockHeight, "fromblockNumber", 0, "Target the beginning block number,default:0")
	// flag.IntVar(&endBlockHeight, "endblockNumber", 0, "Target the end block number, default:latest")
	// flag.Parse()
	// if address == "" || !addresscheck.CheckAddress(address) {
	// 	fmt.Println("Provide address or an invalid address")
	// 	return
	// }
	address := "0xA94bb988cE6Be550A45d9E5cE055B1044559ABf3"
	txDb := db.Buildconnect()
	defer txDb.Close()
	db.StoreTx(txDb, address)
}
