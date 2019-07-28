package main

import "fmt"

func main() {
	score := 100
	if score >= 60 && score <= 90 {
		fmt.Println("及格就好")
	} else if score <=100 && score >=90{
		fmt.Println("一枝独秀")
	} else {
		fmt.Println("继续努力")
	}
	if age := 18; age ==18 {
		fmt.Println("年轻真好！")
	}
	//err 局部变量不能在外部调用
	//fmt.Println(age)
}