package main

import (
	"fmt"
	"encoding/json"
)
//json序列化
//首字母大写相当于public属性
//首字母小写相当于private属性
//emp 员工结构体
//若不首字母大写 导致转换产生 空[]byte
//定义元信息：json tag
type Emp struct {
	ID int `json:"id"`
	Gender string `json:"gender"`
	Name string `json:"name"`
	Salary float32 `json:"salary"`
}
func main(){
	var emp1 = Emp{
		ID:1,
		Gender:"女",
		Name:"娜扎",
		Salary:998,
	}
	fmt.Println(emp1)
	//序列化，把编程语言里的数据转换成JSON 格式字符串
	v, err := json.Marshal(emp1)
	if err != nil {
		fmt.Printf("转换json失败 %s\n", err)
	}
	fmt.Println(v)
	fmt.Println(string(v))//把[]byte转换成string
	var emp2 = new(Emp)
	str := string("{\"ID\":1,\"Gender\":\"男\",\"Name\":\"aicken\",\"Salary\":998}")
	//反序列化：把满足JSON格式的字符串转换成，当前编程语言里面的对象
	json.Unmarshal([]byte(str),emp2)
	fmt.Println(*emp2)
	fmt.Printf("%#v\n",emp2)
}