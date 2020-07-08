package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "jcc:jcc@tcp(172.16.0.203:3306)/city_debug"

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("")
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10)
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err :%v\n", err)
	}

	fmt.Println("数据库链接成功")
}