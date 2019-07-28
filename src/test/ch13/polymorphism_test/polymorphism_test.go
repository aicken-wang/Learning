package polymorphism_test

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
