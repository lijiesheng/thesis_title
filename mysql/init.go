package mysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	db, err := sqlx.Connect("mysql", "root:ljs024816@tcp(101.42.97.221:3306)/ljs_test?parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	if err := db.Ping(); err != nil {
		fmt.Println("数据库失败")
		return
	} else {
		fmt.Println("数据库成功")
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	Db = db
}
