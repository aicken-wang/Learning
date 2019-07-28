package main

import (
	"database/sql"
	"fmt"
	"testing"
	//添加驱动，由于不使用所以需要空导入,在前面添加_
	_ "github.com/go-sql-driver/mysql"
)

func Test_mysql_del(t *testing.T)  {
	fmt.Println("删除数据库操作")
	//打开连接
	db,err := sql.Open("mysql","root:123456@tcp(localhost:3306)/goDB")
	if err != nil {
		fmt.Println("数据库打开失败",err)
		return
	}
	//关闭数据库
	defer func() {
		if db != nil {
			db.Close()
		}
		fmt.Println("关闭连接")
	}()
	//预处理SQL
	Stmt,err := db.Prepare("delete  from people where id =?")
	if err != nil {
		fmt.Println("预处理失败", err)
	}
	//关闭对象
	defer func() {
		if Stmt != nil {
			Stmt.Close()
			fmt.Println("关闭Stmt")
		}
	}()
	// Exec() 参数为不定项参数，对应占位符?个数 传入ID
	//获取结果
	res,err := Stmt.Exec(1)
	//受影响的行数
	count,err := res.RowsAffected()
	if err != nil {
		fmt.Println("获取结果失败",err)
	}
	if count > 0 {
		fmt.Println("删除成功")
	} else {
		fmt.Println("删除失败")
	}

}
