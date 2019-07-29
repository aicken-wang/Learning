package main

//结构体
//创建新的类型要使用type关键字
import "fmt"

type emp struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var P1 = emp{
		name:   "aicken",
		age:    18,
		gender: "男",
		hobby:  []string{"coding", "run", "乒乒球", "双色球"},
	}
	fmt.Println(P1)
	//实例化一个joseph,struct 是值类型的，如果初始化之后没有给属性赋值，struct 是默认值
	var joseph = emp{}
	fmt.Println(joseph.age)
	fmt.Println(joseph.hobby)
	//实例化方法2 new(T) T 类型 或者结构体
	var king = new(emp)
	king.name = "king"
	king.age = 18
	king.hobby = []string{"play Game"}
	fmt.Println(*king)
	// 实例化方法3
	var nazha = &emp{}
	nazha.name = "娜扎"
	nazha.age = 18
	nazha.hobby = []string{"唱歌", "colsplay"}
	fmt.Println(*nazha)
}
