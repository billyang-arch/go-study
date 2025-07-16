package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:182598@yp@tcp(127.0.0.1:3306)/dev?charset=utf8mb4&parseTime=True"
	//sql_demo.Open() 会返回一个 *sql_demo.DB 对象，但它不会立刻与数据库建立实际连接。
	//它仅仅会根据提供的 DSN 信息（如数据库类型、地址、端口、用户名、密码等）返回一个连接池的实例。
	//连接池本身并不代表一个活跃的连接，它是一个用来管理多个数据库连接的池子。
	//sql_demo.Open() 本质上是一个延迟连接的操作，直到真正需要执行查询时，连接才会被建立。通过连接池的管理，程序可以有效地复用数据库连接，减少连接的开销。
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	// 调用查询单条记录的函数
	//queryRowDemo()
	//插入数据
	//insertRowDemo()
	//更新数据
	//updateRowDemo()
	//删除数据
	//deleteRowDemo()
	// 调用查询多条记录的函数
	//queryMultiRowDemo()
	//预处理插入
	//prepareInsertDemo()
	//预处理查询
	//prepareQueryDemo()

	//事务处理
	transactionDemo()

	//sql注入示例
	//sqlInjectDemo("xxx' or 1=1#")
	//sqlInjectDemo("xxx' union select * from user #")
	//sqlInjectDemo("xxx' and (select count(*) from user) <10 #")
}

// sql注入示例
func sqlInjectDemo(name string) {
	//我们任何时候都不应该自己拼接SQL语句！
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var u user
	err := db.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", u)
}
