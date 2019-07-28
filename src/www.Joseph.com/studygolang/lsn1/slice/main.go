package main

import "fmt"

func main() {
	var a =[3]int{1, 2, 3} //数组
	var b = []int{1, 2, 3} //切片
	fmt.Printf("a：%T, b: %T\n", a, b)
	//声明切片方式2 :从数组得到切片
	var c []int
	//[:]切片左包含又不包含 从下标的角度去理解 
	c = a[0: 2]
	fmt.Printf("c:%T \n",c)
	//切片的大小 len(c)
	//切片的容量 cap(c)
	//实际切片的长度
	fmt.Println(len(c))
	//该切片对应的底层数组的长度
	fmt.Println(cap(c))
}