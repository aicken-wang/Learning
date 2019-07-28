package type_test

import (
	"testing"
)
type MyInt int64 // 声明一个int64 的别名MyInt
func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a) //显示转换
	// b = a //Err 不能隐式转换
	var c MyInt
	//c = b //Err 同类型的别名也不能隐式转换 严格
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a :=1
	aPtr := &a
	// aPtr = aPtr + 1 //Err 不支持运算
	t.Log(a, aPtr)
	t.Logf(" %T %T", a, aPtr)
}
func TestString(t *testing.T) {
	var s string
	//string是值类型默认是空字符串
	t.Log("*" + s +"*")
	t.Log(len(s))
	if s == ""{
		t.Log("字符串是空")
	}
}