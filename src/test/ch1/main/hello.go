package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println(os.Args)
	if len(os.Args)>1 {
		fmt.Println("hello world cmd: ",os.Args[1])
	}
	fmt.Print("hello world")
	os.Exit(0)
}