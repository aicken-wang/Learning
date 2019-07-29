package main

import "fmt"

//设计一个程序，存储学员信息：id 年龄 薪水 等级
//根据id获取员工信息

func main() {
	empMap := make(map[string]map[string]string, 10)
	//数据初始化内层map
	empMap["娜扎"] = make(map[string]string, 4)
	//初始化4个元素
	empMap["娜扎"]["id"] = "1"
	empMap["娜扎"]["age"] = "18"
	empMap["娜扎"]["salary"] = "30k"
	empMap["娜扎"]["level"] = "T3"

	empMap["小王子"] = make(map[string]string, 4)
	//初始化4个元素
	empMap["小王子"]["id"] = "2"
	empMap["小王子"]["age"] = "20"
	empMap["小王子"]["salary"] = "20k"
	empMap["小王子"]["level"] = "T2"

	for k, v := range empMap {
		//打印部门
		fmt.Println(k)
		//遍历部门下的员工
		for k1, v1 := range v {
			fmt.Println(k1, v1)
		}
	}
	//根据ID获取员工的信息
	fmt.Println("——————————————————————————")
	for _, v := range empMap {
		fmt.Println(v)
		id, ok := v["id"]
		if ok {
			fmt.Println("2222 ", id)
			if id == "2" {
				for k1, v1 := range v {
					fmt.Println(k1, v1)
				}
			}
		} else {
			fmt.Println("id 字段未配置")
		}
	}
}
