package main

import "fmt"
func main() {
	//var  a *int //a 是一个int 类型的指针 err 错误的写法
	//*a =10 //panic: runtime error: invalid memory address or nil pointer dereference
	//上述是错误的写法
	//make 给slice、map、channel 引用类型初始化分配内存
	//new 用来初始化值类型的指针
	a := new(int)
	*a = 10
	fmt.Println(a)
	fmt.Println(*a)
	arr := new([3]int)
	fmt.Println(arr)
	arr[0] =1
	fmt.Println(*arr)
}