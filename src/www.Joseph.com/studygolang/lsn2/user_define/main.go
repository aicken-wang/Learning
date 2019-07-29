package main

import "fmt"

//自定义类型
type w_size32 int

//类型别名:只存在代码编写过程中，代码编译之后根本不存在
// 提高代码的可读性，封装第三方接口使用

type naZha = int

var char uint8

var char byte

func main() {

	var a w_size32

	fmt.Println(a)
	fmt.Printf(" %T \n ", a)

	var b naZha
	fmt.Printf(" %T \n ", b)

	var c byte
	fmt.Printf("%T \n ", c)
}
