package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func Test_select(t *testing.T){
	fmt.Println("查询数据库操作")
	//打开连接
	DB,err := sql.Open("mysql","root:123456@tcp(localhost:3306)/goDB")
	DB.Ping()
	// 打开数据库失败
	if err != nil {
		fmt.Println("打开数据库失败",err)
		return
	}
	//结束后关闭连接
	defer func() {
		if DB != nil {
			DB.Close()
			fmt.Println("关闭连接")
		}
	}()
	//预处理SQL
	Stmt,err := DB.Prepare("select * from people")
	//关闭Stmt对象
	defer func() {
		if Stmt != nil {
			Stmt.Close()
			fmt.Println("关闭 Stmt")
		}
	}()
	//预处理错误
	if err != nil {
		fmt.Println("预处理执行错误",err)
	}
	//执行SQL
	rows,err := Stmt.Query()
	if err != nil {
		fmt.Println("查询失败",err)
	}
	//循环遍历结果
	for rows.Next() {
		var id int
		var name, address string
		//把行内值赋给变量
		rows.Scan(&id, &name, &address)
		fmt.Println(id, name, address)
	}
	defer func() {
		if rows != nil {
			rows.Close()
			fmt.Println("关闭结果集")
		}
	}()
	//获取结果
}