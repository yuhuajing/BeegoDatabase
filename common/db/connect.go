package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type MysqlCon struct {
	//Name            string
	Addr            string
	port            int
	Db              string
	Username        string
	Password        string
	MaxIdealConn    int
	MaxOpenConn     int
	ConnMaxLifetime int
}

var mysqlCon = MysqlCon{
	//"FilNest",
	"127.0.0.1",
	3306,
	"testUser",
	"root",
	"123456",
	10,
	256,
	600,
}

func Buildconnect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		mysqlCon.Username, mysqlCon.Password, mysqlCon.Addr, mysqlCon.port, mysqlCon.Db, "10s")
	//mysql connection
	dba, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Connect error:%s\n", err)
	}
	return dba
}
