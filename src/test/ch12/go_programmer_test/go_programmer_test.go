package go_programmer_test

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	WriteHellWorld() Code
}
type GoProgrammer struct {
	
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}
type CProgrammer struct {

}
func (p *CProgrammer) WriteHelloWorld() Code {
return "printf(\"Hello World \n\")"
}

func writeFirstProgram(p Programmer) {
	//实例的类型 %T
	fmt.Printf("%T %v", p, p.WriteHellWorld())
}
func TestPolymorphism( t *testing.T)  {
	 og := &GoProgrammer{}
	//CProg := new(CProgrammer)
	writeFirstProgram(og)
	//writeFirstProgram(CProg)
}

/*
1空接口可以表示任何类型
2.通过断言来将空接口转换为制定类型
v,ok := p.(int) //ok 为true 转换成功
*/