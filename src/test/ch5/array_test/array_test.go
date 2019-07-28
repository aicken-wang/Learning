package array_test

import (
	"testing"
)

/*
数组的声明
var a [3]int //声明并初始化为默认为零值
a[0] = 1
b := [3]int{1,2,3} //声明同时初始化
c := [2][2]int{{1,2},{3,4}} //多维数组初始化
*/

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1,2,3,4}
	arr2 := [...]int{1,2,3,4,5}
	arr1[1] = 5
	t.Log(arr[1])
	t.Log(arr1,arr2)
}
func TestArrayTranvel(t *testing.T)  {
	arr2 := [...]int{1,2,3,4,5}
	for i :=0; i < len( arr2 );i++  {
		t.Log(arr2[i])
	}
	for idx, e := range arr2 {
		t.Log(idx,e)
	}
	//_ 占位
	for _, e := range arr2 {
		t.Log(e)
	}
}
/*
数组截取
a[开始索引（不包含），结束索引（包含）] 左开右闭
a := [...]int{1,2,3,4,5}
a[1:2]//2
a[1:3]//2,3
a[1:len(a)] // 2, 3, 4, 5
a[1:]//2,3,4,5
a[:3]//1,2,3
Go不支持负数
*/
func TestArraySection (t *testing.T) {
	arr := [...]int{1,2,3,4,5}
	arr_sec := arr[:]
	t.Log(arr_sec)
}