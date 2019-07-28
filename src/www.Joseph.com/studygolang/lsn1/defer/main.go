package main

import "fmt"

func main() {
	fmt.Println("app start....")
	defer fmt.Println("moduleA load successful")
	defer fmt.Println("moduleB load successful")
	fmt.Println("start end...")
}
