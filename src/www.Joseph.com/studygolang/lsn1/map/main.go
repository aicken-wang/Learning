package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

const N = 10

func main() {
	//光声明map类型，没有初始化，map默认初始值
	var m1 map[string]string
	fmt.Println(m1 == nil)
	//map的初始化
	m1 = make(map[string]string, 8)
	m1["nazha"] = "娜扎"
	m1["xiaowangzi"] = "小王子"
	fmt.Printf("m1:%#v\n", m1)
	fmt.Printf("type:%T\n", m1)
	//声明map同时初始化
	var m2 = map[string]int{
		"nazha_age": 18,
	}
	fmt.Printf("m:%v  %#v\n", m1, m2)

	var scoreMap = make(map[string]int, 8)
	scoreMap["aicken"] = 90
	scoreMap["Joseph"] = 60
	v, ok := scoreMap["wangxiaoer"]
	if ok {
		fmt.Printf("王小二参加考试 考试成绩 %d\n", v)
	} else {
		fmt.Println("王小二没来考试")
	}
	for k, v := range m1 {
		fmt.Printf("k = %s v = %s\n", k, v)
	}
	for k := range m1 {
		fmt.Printf("k = %s \n", k)
	}
	for _, v := range m1 {
		fmt.Printf(" v = %s\n", v)
	}
	//删除
	delete(m1, "xiaowangzi")

	fmt.Println("遍历m1 map")
	for k, v := range m1 {
		fmt.Printf("k = %s v = %s\n", k, v)
	}

	//按顺序添加k v
	var emp = make(map[string]int, 50)
	//添加50个员工
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("emp%02d", i+1)
		value := rand.Intn(100) // 0~99
		emp[key] = value
	}
	//不是按序遍历
	for k, v := range emp {
		fmt.Println(k, v)
	}
	//按序遍历
	keys := make([]string, 0, 100)
	for k := range emp {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, emp[key])
	}

	//元素类型未map的切片
	var mapSlice = make([]map[string]int, 10, 10) //初始化了切片，元素map未初始化
	//mapSlice[0]["aicken"] = 1 //err
	fmt.Println(mapSlice)
	//对第一个元素初始化
	mapSlice[0] = make(map[string]int, 10) //元素map初始化
	mapSlice[0]["aicken"] = 1
	fmt.Println(mapSlice)
	fmt.Println(mapSlice[0])
	//值为切片类型的map
	var sliceMap = make(map[string][]string, 1) //只是完成map的初始化，元素未初始化
	v1, ok1 := sliceMap["China"]
	if ok1 {
		fmt.Println(v1)
	} else {
		//sliceMap中没有CHINA这个键
		sliceMap["China"] = make([]string, 4, 8)
		sliceMap["China"][0] = "Shenzhen"
		sliceMap["China"][1] = "Beijing"
		sliceMap["China"][2] = "Shanghai"
		sliceMap["China"][3] = "Hangzhou"
	}
	//遍历sliceMap
	for k, v := range sliceMap {
		fmt.Println(k, v)
	}
	//统计一个字符串中每个单词出现的次数
	// Little penguin are you ok xi xi .
	youSay := "Little penguin are you ok, xi xi"
	var wordCnt = make(map[string]int, N)
	//1.字符串中由哪些单词
	words := strings.Split(youSay, " ")
	//遍历单词
	for _, word := range words {
		v, ok := wordCnt[word]
		if ok {
			//统计单词出现的次数
			wordCnt[word] = v + 1
		} else {
			//给没统计的单词赋初值为1，map中没有记录的单词
			wordCnt[word] = 1
		}
	}
	for k, v := range wordCnt {
		fmt.Println(k, v)
	}

}
