package main

import "fmt"

var num = 10

func testGlobal() {
	fmt.Println("全局变量 num %d \n", num)
}

//参数: 参数1 参数2 是同类型的int 参数3是一个返回值为int的函数，接收两个形参
//返回值:int
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

func main() {
	testGlobal()
	f := testGlobal
	fmt.Printf("%T\n", f)
	fmt.Println(calc(10, 20, add))
	fmt.Println(calc(10, 20, sub))
}
