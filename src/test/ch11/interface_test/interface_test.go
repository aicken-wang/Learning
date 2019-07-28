package interface_test

import "testing"

/**
接口定义
type Programmer interface {
	WriteHelloWorld() Code
}
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
接口实现
type GoProgrammer struct {

}
func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\" Hello World'!\")"
}
*/
//定义接口
type Programmer interface {
	WriteHelloWorld() string
}
//实现接口
type GoProgrammer struct {

}
//定义方法 函数签名WriteHelloWorld() string 保持一致
//decker
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World!\")"
}
func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
/*
Go 接口
与其它主要编程语言的差异
1.接口为非入侵性，实现不依赖于接口定义
2.所以接口的定义可以包含在接口使用者包内
*/

/*
// prog 接口变量
var prog Coder = &GoProgrammer {} // 接口
//类型
type GoProgrammer struct {
}
//数据
&GoProgrammer{}
*/
/*
自定义类型
1.type IntConvertionFn func(n int) int
2.type MyPoint int
*/