package main

import "fmt"

//匿名字段
type emp struct {
	id int
	name string
	string
}
type address struct {
	province string
	city string
}
//结构体的嵌套
type people struct {
	name string
	age int
	addr address
}
type man struct {
	name string
	age int
	address
}

func main() {
	var e = emp{
		id:001,
		name:"娜扎",
	}
	fmt.Println(e.name)
	fmt.Println(e.string)

	var p = people {
		name:"aicken",
		age :18,
		addr: address{
			province:"广东",
			city:"深圳",
		},
	}
	fmt.Println(p)
	var m = man {
		name:"aicken",
		age :18,
		address:address {
			province:"广东",
			city:"深圳",
		},
	}
	fmt.Println(m.province)
	fmt.Println(m.address.province)
}