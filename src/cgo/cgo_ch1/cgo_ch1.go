package main

import (
	"fmt"
)

func main() {
	var a, b = 1,2
	a, b = b, a
	fmt.Println(a, b)
	// 内置支持并发
	go func() {
		println("Hello ")
	}()
}
