package main

import "fmt"
//panic 致命的错误 
//使用panic("这是一个异常 TEST")
func f1() {
	defer func() {
		err := recover() //尝试将函数从当前的异常状态恢复过来
		//err  => runtime error: index out of range
		fmt.Printf("recover 捕获到了panic的异常错误信息 %s\n",err)
	}()
	panic("panic异常 TEST")
}
func main()  {
	f1()
	fmt.Println("恢复执行")
	//var a []int
	//a[0] = 100 //panic: runtime error: index out of range
}