package main

import (
	"fmt"
	"strings"
)

var (
	coins =50
	users = []string{"Matthew", "sorry","Augustus"}
	distribution = make(map[string]int, len(users))
)
func dispatchCoin() int {
	//1 遍历名字的切片
	sum := 0
	for _, name := range users {
		//使用switch实现
		//根据规则分发金币
		if strings.Contains(name, "e") || strings.Contains(name, "E") {
			distribution[name] = distribution[name] + 1
			sum++
		} else if strings.Contains(name, "i") ||strings.Contains(name, "I") {
			distribution[name] = distribution[name] + 2
			sum += 2
		} else if strings.Contains(name, "o") ||strings.Contains(name, "O") {
			distribution[name] = distribution[name] + 3
			sum += 3
		}
	}
	ret := coins - sum 
	return ret
	
}
func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	for k, v := range distribution {
		fmt.Println(k,v)
	}
}