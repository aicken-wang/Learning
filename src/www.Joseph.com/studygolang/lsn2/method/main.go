package main

import "fmt"
//可以给任意类型追加方法
//不能给别的包定义的类型添加方法
type MyInt int 
//当一个函数 有了类型定义的接收者 就是一个方法
func (m *MyInt) sayHello() {
	fmt.Println("Hello MyInt")
}
//cannot define new methods on non-local type int
// func (i *int) sayHi() {
// 	fmt.Println(" err ")
// }
func main() {
	var a MyInt
	fmt.Println(a)
	//方法必须要注册到实例上
	a.sayHello()
	
}