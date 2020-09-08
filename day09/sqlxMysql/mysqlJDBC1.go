package sqlxMysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var err error
var db *sqlx.DB

func InitDb() {

	db, err = sqlx.Open("mysql", "root:root@tcp(localhost:3306)/test?charset=utf8")
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

type product struct {
	Id    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Price int64  `json:"price" db:"price"`
}

func FindOne(id int64) {

	sql := "select id,name ,price from product where id=?"
	var p product
	err := db.Get(&p, sql, id)
	if err != nil {
		log.Fatal("select one error ", err)
	}
	fmt.Printf("查找id为%d的product:%v \n", id, p)
}
func FindMore(limit int) {
	sql := `select id,name ,price from product limit ?`

	var plist []product
	err := db.Select(&plist, sql, limit)
	if err != nil {
		log.Fatal("select more error ", err)
	}
	fmt.Println("product:", plist)

}
