package main

import (
	"fmt"
)

//什么是预处理
//普通SQL语句执行过程：
//
//客户端对SQL语句进行占位符替换得到完整的SQL语句。
//客户端发送完整SQL语句到MySQL服务端
//MySQL服务端执行完整的SQL语句并将结果返回给客户端。
//预处理执行过程：
//
//把SQL语句分成两部分，命令部分与数据部分。
//先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
//然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
//MySQL服务端执行完整的SQL语句并将结果返回给客户端。

//预处理的工作原理：
//预编译 SQL 语句：首先，SQL 语句只会被解析和编译一次，MySQL 将这个 SQL 语句的执行计划存储在服务器内部。
//参数绑定：之后，在执行该 SQL 语句时，只需要绑定不同的参数，而不需要重新编译 SQL 语句。
//执行 SQL：执行时，MySQL 只处理不同的输入参数，利用已编译好的执行计划直接执行 SQL 查询。

//优势：
//提高性能：每次执行 SQL 时，不需要重复编译，节省了 CPU 和内存资源。
//减少 SQL 注入风险：使用预处理时，SQL 语句和数据分离，数据作为参数传递，避免了 SQL 注入攻击。

// 预处理查询示例
func prepareQueryDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 预处理插入示例
func prepareInsertDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("小王子", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = stmt.Exec("沙河娜扎", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	fmt.Println("insert success.")
}
