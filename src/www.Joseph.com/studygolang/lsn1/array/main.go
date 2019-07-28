package main

import "fmt"

func main() {
	var a[3] int 
	a =  [3] int{1, 2, 3}
	fmt.Println(a)

	var b[3][2]int
	b = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
	}
	fmt.Println(b)
	//声明同时初始化
	var c = [3][2]int{
		{1, 2},
		{3,4},
	}
	fmt.Println(c)
	//注意事项，多维数组除了第一次其它的都不能用...
	fmt.Println(c[1][1])

	for index := 0; index < len(c); index++ {
		for i :=0; i < len(c[index]); i++ {
			fmt.Println(c[index][i])
		}	
	}
	fmt.Println("使用range遍历 ")
	for _,v1 := range c {
		for _,v2 :=range v1{
			fmt.Println(v2)
		}
	}
	fmt.Println("数组赋值是值传递 ")
	ta := [...]int{1, 2}
	tb := ta
	tb[0] = 100
	fmt.Println(ta)
	fmt.Println(tb)
}