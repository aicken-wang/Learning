/*
	地址：就是内存地址（用字节来描述的内存地址）
	指针：指针是带类型的
	& 表示去地址
	* 根据地址取值，解引用
*/
package main

import "fmt" 

func main() {
	var a int 
	fmt.Println(a)
	b := &a // 取变量a 的内存地址
	fmt.Printf("b = %v \n",b)
	fmt.Printf("type b : %T \n",b)
	c := "娜扎"
	//b = &c //cannot use &c (type *string) as type *int in assignmen
	fmt.Printf("type &c %T  \t &c = %v \n", &c, &c)
	d := 100
	b = &d //同一类型可以赋值，本质是地址 带有类型
	// * 取地址的值 
	//go语言中对地址只能读取地址内容
	fmt.Println(*b)
	//指针可以做逻辑判断
	fmt.Println(b == &d)
	//指针的应用
	arr := [3]int{1,2,3}
	modifyArray( &arr )
	fmt.Println( arr )
}

func modifyArray(a1 *[3]int) {
	(*a1)[0] =100 //修改a1第一元素的值为100 
	//语法糖因为Go语言中指针不支持修改 a1[0] = 100
}