package db

import (
	"fmt"
	//"main/common/db"
	"main/common/table"
	//"main/core"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jinzhu/gorm"
)

// 插入
func Insert(dba *gorm.DB, txdata *table.ResponseTxData) {
	for _, tx := range txdata.Result {
		res := Read(dba, &table.TxResult{}, tx.Hash)
		fmt.Println(tx.Hash)
		fmt.Printf("insertTx:%d\n", res.RowsAffected)
		logDb := Buildconnect()
		defer logDb.Close()
		if res.RowsAffected == 0 {
			//	str = append(str, tx.Hash)
			dba.Create(&table.TxResult{
				BlockNumber:       tx.BlockNumber,
				TimeStamp:         tx.TimeStamp,
				Hash:              tx.Hash,
				Nonce:             tx.Nonce,
				BlockHash:         tx.BlockHash,
				TransactionIndex:  tx.TransactionIndex,
				From:              tx.From,
				To:                tx.To,
				Value:             tx.Value,
				Gas:               tx.Gas,
				GasPrice:          tx.GasPrice,
				IsError:           tx.IsError,
				TxReceiptStatus:   tx.TxReceiptStatus,
				Input:             tx.Input,
				ContractAddress:   tx.ContractAddress,
				CumulativeGasUsed: tx.CumulativeGasUsed,
				GasUsed:           tx.GasUsed,
				Confirmations:     tx.Confirmations,
				MethodId:          tx.MethodId,
				FunctionName:      tx.FunctionName})
		}
		StoreLog(logDb, tx.Hash)
	}
}

func InsertLog(dba *gorm.DB, txdata []*types.Log) bool {
	for _, tx := range txdata {
		res := Readlogs(dba, &table.Log{}, tx.TxHash.Hex(), int(tx.Index))
		//fmt.Println(res.RowsAffected)
		fmt.Println(tx.TxHash.Hex())
		fmt.Printf("insertLogs:%d\n", res.RowsAffected)
		if res.RowsAffected == 0 {
			dba.Create(&table.Log{
				Address:     tx.Address.Hex(),
				Topics:      parsearray(tx.Topics),
				Data:        parseByteToString(tx.Data),
				BlockNumber: tx.BlockNumber,
				TxHash:      tx.TxHash.Hex(),
				TxIndex:     tx.Index,
				BlockHash:   tx.BlockHash.Hex(),
				Index:       tx.TxIndex,
				Removed:     tx.Removed,
			})
		}
	}
	return true
}

func parsearray(topics []common.Hash) string {
	str := ""
	for _, top := range topics {
		str += top.Hex()
	}
	return str
}

func parseByteToString(b []byte) string {
	str := ""
	//fmt.Println(len(b))
	for i := 0; i < len(b)/32; i++ {
		num := big.NewInt(0)
		num.SetBytes(b[i*32 : i*32+32])
		resInt := fmt.Sprintf("%d", num)
		str += "data" + strconv.Itoa(i) + ": " + resInt + " "
		//num, err := strconv.ParseInt(_b[i:i+64], 16, 32)
	}
	return str
}

// 查
func Read(dba *gorm.DB, tx *table.TxResult, hash string) *gorm.DB {
	//db.First(product, 1) // 查询id为1的product
	//db.First(product, "code = ?", "L1212")

	result := dba.Select("hash").Find(tx, "hash = ?", hash)

	// SELECT name, age FROM users;
	//db.Select([]string{"name", "age"}).Find(&users)

	//result := db.Find(tx, "hash = ?", hash)
	return result
}

func Readlogs(dba *gorm.DB, tx *table.Log, hash string, index int) *gorm.DB {
	//db.First(product, 1) // 查询id为1的product
	//db.First(product, "code = ?", "L1212")
	result := dba.Where("tx_hash = ? AND tx_index = ?", hash, index).Find(&table.Log{})
	//result := dba.Where(table.Log{TxHash: hash, Index: uint(index)}).Find(&table.Log{})
	return result
}

// 更新 - 更新product的price为2000
func Modify(dba *gorm.DB, product *table.Product) {
	dba.Model(product).Update("Price", 2000)
}

// 删除指定条件
func DeleteRow(dba *gorm.DB, str string, id uint) {
	dba.Unscoped().Where(str, id).Delete(&table.Product{})
}

// 删除 - 删除表
func Delete(dba *gorm.DB, product *table.Product) {
	dba.DropTable(table.Product{})
}
