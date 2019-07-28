package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func Test_update(t *testing.T){
	fmt.Println("更新数据库操作")
	/*
	   第一个参数:连接什么数据库
	   第二个参数:连接字符串
	   语法:数据库登录用户名:密码@tcp(mysql主机ip:端口)/database名称
	*/
	//打开连接
	db,err := sql.Open("mysql","root:123456@tcp(localhost:3306)/goDB")

	if err != nil {
		fmt.Println("数据库打开失败", err)
		return
	}
	//关闭连接
	defer func() {
		if db != nil {
			db.Close()
		}
		fmt.Println("关闭连接")
	}()
	/*
	   准备处理SQL语句
	   支持占位符,防止SQL注入
	*/
	Stmt,err := db.Prepare("update people set name=?,address=? where id=?")
	//错误处理
	if err != nil {
		fmt.Println("预处理失败", err)
	}
	//关闭对象
	defer func() {
		if Stmt != nil {
			Stmt.Close()
			fmt.Println("Stmt关闭")
		}
	}()
	//Exec() 参数为不定项参数，对应占位符?个数
	res,err := Stmt.Exec("小王","湖南",2)
	//错误处理
	if err != nil {
		fmt.Println("执行SQL错误", err)
	}
	//受影响行数
	count,err := res.RowsAffected()
	if err != nil {
		fmt.Println("获取结果失败")
	}
	//如果修改前和修改后的值相同,RowsAffected()返回0
	if count > 0 {
		fmt.Println("id = 2 修改成功")
	} else {
		fmt.Println("修改失败")
	}
}
