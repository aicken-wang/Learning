package function_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
函数是一等公民
与其他主要编程语言的差异
1.可以有多个返回值
2.所有参数都是值传递，slice、map channel会有传引用的错觉
3.函数可以作为变量的值
4.函数可以作为参数和返回值
*/
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}
func timeSpent(inner func(op int) int)func(op int) int{
	return func(n int) int{
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}
func slowFun(op int) int {
	time.Sleep(time.Second*1)
	return op
}
func TestFunc(t *testing.T) {
	//使用占位符 _ 可以忽略该返回值
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

