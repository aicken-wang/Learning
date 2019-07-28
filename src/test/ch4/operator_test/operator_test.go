package operator_test

import (
	"fmt"
	"testing"
)

func TestCompareArray(t *testing.T)  {
	a := [...]int{1,2,3,4}
	b := [...]int{1,2,3,4}
	//c := [...]int{1,2,3,4,5}
	//
	t.Log(a == b)
	//t.Log(c)
	//t.Log( a == c)
}
/*
循环
Go 语言仅支持循环关键字 for
for j :=7; j <=9; j++
*/
func TestForLoopT(t *testing.T)  {
	for j := 0; j <= 10; j++ {
		t.Log("执行第" ,j , "次")
	}
}
/*
while 条件循环
while (n < 5)
*/
func TestFor(t *testing.T){
	n := 0
	for n < 5 {
		n++
		fmt.Println(n)
	}
}
/*
while 条件循环
while (true)
*/
func TestForLoop(t *testing.T){
	//n := 0
	//for {
	//	n++
	//	fmt.Print("死循环")
	//}
}
func TestSwitchMultiCase(t *testing.T) {
	for i :=0; i <5; i++  {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1,3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")


		}
	}
}
func TestSwitchCaseCondition(t *testing.T) {
	for i :=0; i <5; i++  {
		switch {
		case i%2==0:
			t.Log("Even")
		case i%2==1:
			t.Log("Odd")
		default:
			t.Log(" 不会走这个分支")
		}
	}
}
