package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 不要忘了导入数据库驱动
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:182598@yp@tcp(127.0.0.1:3306)/dev?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("connect db failed, err:%v\n", err)
		return
	}
	//queryRowDemo()
	//insertUserDemo()
	//queryMultiRowDemo()
	//namedQuery()
	//users := []*user{
	//	{Name: "Tom", Age: 18},
	//	{Name: "Jack", Age: 20},
	//	{Name: "Alice", Age: 22},
	//}
	//BatchInsertUsers3(users)

	ids := []int{2, 1, 3}
	res, err := QueryAndOrderByIDs(ids)
	if err != nil {
		return
	}
	fmt.Println(res)
}
