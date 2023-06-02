package table

import "github.com/jinzhu/gorm"

type Log struct {
	gorm.Model
	Address     string `json:"address" gencodec:"required"`
	Topics      string `json:"topics" gencodec:"required"`
	Data        string `json:"data" gencodec:"required"` //hexutil.Bytes  `json:"data" gencodec:"required"`
	BlockNumber uint64 `json:"blockNumber"`              //hexutil.Uint64 `json:"blockNumber"`
	TxHash      string `json:"transactionHash" gencodec:"required"`
	TxIndex     uint   `json:"transactionIndex"` // hexutil.Uint   `json:"transactionIndex"`
	BlockHash   string `json:"blockHash"`
	Index       uint   `json:"logIndex"` //  hexutil.Uint   `json:"logIndex"`
	Removed     bool   `json:"removed"`
}
