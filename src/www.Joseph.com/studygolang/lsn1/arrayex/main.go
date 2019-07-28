package main

import "fmt"

func main() {
	var arr = [...]int{1,2,3,4,5,6,7,8,9}
	elem := 8
	length := len(arr)
	for i1,v1 := range arr {
		other := elem - v1
		for i2 := i1 + 1; i2 < length; i2++ {
			if (arr[i2] == other) {
				fmt.Printf("arr[%d], arr[%d] => V[%d, %d] \n",i1, i2, arr[i1], arr[i2])
			}
		}
	}
}