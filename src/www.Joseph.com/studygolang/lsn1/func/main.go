package main

import "fmt"

//函数传值
func youSay(name string) {
	fmt.Printf("youSay: %s  ptr = %p \n", name, &name)
}

//传指针
func youSayPtr(p *string) {
	fmt.Printf("youSayPtr %s . ptr = %p \n", *p, p)
}

//Go语言中函数参数简写
func add(x, y int) (ret int) {
	//方式一
	//ret := x + y
	//return ret
	//方式二
	//return x + y
	//方式三  前两个可以用func add(x, y int) int 代替 func add(x, y int) (ret int)
	ret = x + y
	return
}

//函数接收可变参数 ...表示可变参数
//可变参数在函数体中时切片类型
func funcArgs(args ...int) int {
	fmt.Println(args)
	fmt.Printf("%T \n", args)
	ret := 0
	for _, arg := range args {
		ret = ret + arg
	}
	return ret
}

//固定参数和可变参数同事出现时，可变参数要放在最后
func argOrder(arg1 int, args ...int) {
	fmt.Println(arg1, "- 分割线 -", args)
}

//定义具有多个返回值的函数
func calc(a, b int) (add, sub int) {
	add = a + b
	sub = a - b
	return
}

//go 语言中没有默认参数
func main() {
	nameStr := "娜扎"
	fmt.Printf("%p \n", &nameStr)
	youSay(nameStr)
	fmt.Println("add(1, 2) = ", add(1, 2))
	youSayPtr(&nameStr)
	fmt.Println(funcArgs(1, 2, 3, 4, 5))
	argOrder(1, 2, 3, 4, 5, 6)
	x, y := calc(10, 20)
	fmt.Println(x, y)
}
