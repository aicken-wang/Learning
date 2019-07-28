package variablelen_test

import (
	"fmt"
	"testing"
)

/*
可变长参数及defer
func sum(ops ... int) int {
	s := 0
	for _, op := range ops {
		s += op
	}
	return s
}
*/
func Sum(ops ...int) int {
	s := 0
	for _, op := range ops {
		s += op
	}
	return s
}
func TestVariableLen(t *testing.T) {
	t.Log(Sum(1,2,3,4))
	t.Log(Sum(1,2,3,4,5))
}

func Clear()  {
	fmt.Println("Clear resource.")

}
//defer 函数
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Print("Start\n")
	defer func() {
		t.Log("我是一个闭包callback")
	}()
	//panic 是无法修复的错误
	panic("err") //defer 仍会执行
	fmt.Println("panic是一个致命的错误，触发它这行代码执行不到")
}