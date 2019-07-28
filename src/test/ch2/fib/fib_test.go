package fib

import (
	"fmt"
	"testing"
)
//全局变量声明
//var a int
func TestFibList(t *testing.T)  {
	//a = 1  //全局变量赋值
	//定义变量方法一
	//var a int = 1
	//var b int = 1
	//定义变量方法二
	//var(
	//	a int = 1
	//	b int = 1
	//)
	// 自动推断变量类型
	a := 1
	b := 1
	fmt.Print(a)
	for i :=0; i<5; i++ {
		//fmt.Print(" ", b)
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()

}
func TestSwap(t *testing.T) {
	a := 1
	b := 2
	//tmp := a
	//a = b
	//b = tmp
	//对多个变量赋值
	a, b = b, a
	t.Log(a, b)
}