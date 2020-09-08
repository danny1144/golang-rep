package main

import (
	"database/sql"
	"day09/sqlxMysql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var err error
var db *sql.DB

func initDb() {

	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal("open url format invalid ", err)
		return
	}
	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatal("connect mysql failed ", err)
		return
	}
	log.Println("连接mysql 成功")
}

func findOne(id int64) {

	sql := `select id,name ,price from product where id=?`
	row := db.QueryRow(sql, id)
	var (
		ids, price int
		name       string
	)
	err = row.Scan(&ids, &name, &price)
	if err != nil {
		log.Fatal("select one error ", err)
	}
	fmt.Println("product", ids, name, price)
}
func findMore(limit int) {
	sql := `select id,name ,price from product limit ?`

	rows, err := db.Query(sql, limit)
	if err != nil {
		log.Fatal("select more error ", err)
	}
	for rows.Next() {
		var (
			ids, price int
			name       string
		)
		rows.Scan(&ids, &name, &price)
		if err != nil {
			log.Fatal("select one error ", err)
		}
		fmt.Println("product", ids, name, price)
	}
}
func main() {
	//initDb()
	//findMore(10)
	sqlxMysql.InitDb()
	sqlxMysql.FindOne(1)
	sqlxMysql.FindMore(10)
}
