package main

import "fmt"

func main() {
	for index := 0; index < 5; index++ {
		for i := 0; i < 3; i++ {
			if index == 1 && i == 1 {
				goto END
			}
		}
	}
	fmt.Println("for for End")
END:
	fmt.Println("End")
}