package main

import "fmt"
//const
const pi = 3.14
//iota 枚举
const (
	n1 = iota //0
	n2   // n2 = iota 1
	n3   // n3 = iota 2
	_ 	 // _  = iota 3
	n5   // n5 = iota 4
)
const (
	_ = iota
	KB = 1 << (10 * iota) //2的10次方
	GB = 1 << (10 * iota) //2的20次方
	TB = 1 << (10 * iota) //2的30次方
	PB = 1 << (10 * iota) //2的40次方
)
/*
0.const声明如果不写，默认就和上一行一样
1.遇到const iota就初始化为0
2.cosnt中每新增一行变量声明iota就递增1
*/
const ()
const (
	a1 = iota //0
	a2 = 100   // a2 = 100 
	a3 = iota   // a3 = iota 2  a2递增了一次
)
func foo() (string, int){
	return "nazha", 18
}
func main() {
    //声明变量必须使用
    var  boy string
    fmt.Println(boy)
    //声明多个变量
    var ( 
        a string
        b bool
        c int
	)
	a = "a"
	b = false
	c = 100
    fmt.Println(a, b, c)
    //声明变量并初始化 err
	// var x string = "Hi" 
    // fmt.Println(x)
    //短变量声明（只能在函数内部使用）
    m := "email"
    fmt.Println(m)
    //类型推导 
    var x = 100
    fmt.Println(x)
	name,age:= foo()
	fmt.Println(name,age)
	n,_ := foo()
	fmt.Println(n)
	fmt.Println("```````````````````")
	fmt.Println(a1, a2, a3)
	fmt.Println(KB, GB, TB, PB)
}