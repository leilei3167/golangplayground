package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB //DB是一个数据库（操作）句柄，代表一个具有零到多个底层连接的连接池
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:8888@/test")
	if err != nil {
		panic(err.Error())
	}
}
