package db

import (
	//"main/common/db"

	"main/common/eclient"
	"main/common/table"

	"github.com/jinzhu/gorm"
)

func StoreLog(datab *gorm.DB, txHash string) bool {
	datab.AutoMigrate(&table.Log{})
	//for _, txhash := range txHash {
	logs := eclient.GetTxlogs(txHash)
	//fmt.Println(logs)
	if len(logs) > 0 {
		InsertLog(datab, logs)
	}
	//}
	return true
}
