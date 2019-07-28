package main

import (
	"fmt"
	"strings"
)

//定义一个函数 返回值也是一个函数
//函数作为 返回值
func cb() func() {
	return func() {
		fmt.Println("娜扎")
	}
}

//闭包
func callback(name *string) func() {
	//name := "娜扎"
	return func() {
		fmt.Printf("hello %s ！\n", *name)
	}
}

//闭包2 名字后面添加后缀
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func calc(sum int) (func(int) int, func(int) int) {
	add := func(i int) int {
		sum += i
		return sum
	}
	sub := func(i int) int {
		sum -= i
		return sum
	}
	return add, sub
}
func main() {
	f := func() {
		fmt.Println("匿名函数")
	}
	f()
	//直接调用匿名函数
	func() {
		fmt.Println("匿名函数")
	}()
	//闭包 =
	r := cb()
	cb()() //直接调用匿名函数
	r()    //执行cb函数内部的匿名函数，
	//闭包 = 函数 + 外层变量的引用
	str := "古力娜扎"
	callback(&str)() //callback() 此时就是一个闭包
	//前面是 suffix name
	fmt.Println(makeSuffixFunc("-wang")("0 aicken"))
	fmt.Println(makeSuffixFunc("-wang")("1 aicken-wang"))
	fmt.Println(makeSuffixFunc(".mp3")("等你下课"))
	fmt.Println(makeSuffixFunc(".com")("https://studygolang"))
	x, y := calc(1)
	//在原来的基础上加10
	fmt.Println(x(10))
	//上一次的基础上减1
	fmt.Println(y(1))
	fmt.Println(y(10))
	/*
		内置函数介绍
		close 主要用来关闭 channel
		len 用来求长度比如string array slice map channel
		new 用来分配内存，主要用来分配值类型，比如int struct 返回的是指针
		make 用来分配内存，主要用来分配引用类型比如chan map slice
		append 用来追加元素到数组，slice中
		panic/recover 用来做错误处理
	*/ 

}
