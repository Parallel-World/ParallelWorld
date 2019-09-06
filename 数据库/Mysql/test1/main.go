package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//用户名：密码@[连接方式](主机名：端口号)/数据库名
	dsn := "root: @tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()
}
