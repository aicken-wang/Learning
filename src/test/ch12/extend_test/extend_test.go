package extend_test

import (
	"fmt"
	"testing"
)

/*
LSP 子类交互原则，所有父类定义的场合子类都放进去。通过重载的方式
*/
type Pet struct {
	
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo (host string) {
	p.Speak()
	fmt.Println(" ",host)
}
//复用
//type Dog struct {
//	p *Pet
//}
//func (d *Dog) Speak() {
//	fmt.Println("wang! wang! wang!")
//}
//
//func (d *Dog) SpeakTo (host string) {
//	d.Speak()
//	fmt.Println(" ",host)
//}
//
//func TestDog(t *testing.T) {
//	dog := new(Dog)
//	dog.SpeakTo("i Go i Go")
//}
//匿名复用
type Dog struct {
	Pet
}
func (d *Dog) Speak() {
	fmt.Println("wang! wang! wang!")
}

//func (d *Dog) SpeakTo (host string) {
//	d.Speak()
//	fmt.Println(" ",host)
//}

func TestDog(t *testing.T) {
	//不支持隐式类型转换
	//var dog  Pet := new(Dog)
	//无法支持 LSP
	//var dog  Dog := new(Dog)
	//var p = *Pet(dog)
	//p.SpeakTo("i Go i Go")
	//dog.SpeakTo("Golang")
	//var dog  Dog := new(Dog)
	dog := new (Dog)
	dog.SpeakTo("i Go i Go")
}