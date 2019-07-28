package main

import "fmt"

//panic和recover
func test1() {
	fmt.Println("func test1")
}
func funcPanic() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover in funcPanic")
		}
	}()
	panic("panic in funcPanic")
}
func test2() {
	fmt.Println("func test02")
}
func main() {
	test1()
	funcPanic()
	test2()

}
