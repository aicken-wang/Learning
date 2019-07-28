package main

import "fmt"

func main (){
	for i := 0; i < 10;i++{
		fmt.Println(i)
	}
	age := 10;
	for age < 100 {
		age++;
		fmt.Println(age)
	}
	age = 10;
	for ;age < 100; {
		age++;
		fmt.Println(age)
	}
	expression := 5;
	//fallthrough 穿透到下一个case,会执行下一个case
	switch expression {
	case 1,2,3,4,5:
		fmt.Println("in the array [1,2,3,4,5]")
	default:
		fmt.Println("other... ")
	}
	
}
