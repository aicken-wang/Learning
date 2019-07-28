package _501

import (
	"fmt"
	"testing"
)

func Test_time(t *testing.T) {
	////声明时间变量
	//var curTime time.Time
	//fmt.Println(curTime)
	//t.Log(curTime)
	//c_t := time.Now()
	//t.Log(c_t)
	////google输入法
	////var arr [3][3]int = [3][3]int{{1, 2, 3}, { 4, 5, 6}, {7, 8, 9}}
	//arr := [3][3]int{
	//	{2, 2, 3},
	//	{4, 5, 6},
	//	{7, 8, 9},
	//}
	//arr0 := arr[0]
	//fmt.Println(arr)
	//fmt.Println(arr0[0], arr0[1], arr0[2])
	//fmt.Println(arr[0][0])
	s1 := []int {1, 2}
	s2 := []int {2,3,3}
	copy(s1,s2)
	fmt.Println(s1)
}
