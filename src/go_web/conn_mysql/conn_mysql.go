package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	fmt.Println("Hello MySQL")
	//1. 打开连接
	db,err := sql.Open("mysql","root:123456@tcp(localhost:3306)/godb")
	db.Ping()
	if err!=nil {
		fmt.Println("打开数据库失败！")
		return
	}
	fmt.Println("数据库连接成功")
	defer func() {
		if db!=nil{
			db.Close()
		}
		fmt.Println("关闭连接")
	}()
	//2.预处理SQL
	//？表示占位符
	Stmt,err :=db.Prepare("insert into people values (default,?,?)")
	if err!=nil {
		fmt.Println("预处理失败！")
		return
	}
	//关闭对象
	defer func() {
		if Stmt !=nil {
			Stmt.Close()
			fmt.Println("Stmt关闭")
		}
	}()
	//参数和占位符对应
	res,err := Stmt.Exec("王小二","深圳")
	if err != nil{
		fmt.Println("sql执行失败！")
		return
	}
	//3.获取结果 插入失败返回 0
	count,err := res.RowsAffected()
	if err != nil {
		fmt.Println("获取结果失败")
	}
	//获取主键
	id,err := res.LastInsertId()
	if err !=nil {
		fmt.Println("获取主键失败",err)

	}
	fmt.Printf("受影响的行数 %d 主键ID %d \n",count,id)

	if count > 0 {
		fmt.Println("插入成功！")
	} else {
		fmt.Println("插入失败！")
	}
}