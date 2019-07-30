package main

import "fmt"

func (a *animal) run() {
	fmt.Printf("%s 在跑\n",a.name)
}
type animal struct {
	name string
}

type dog struct {
	food string
	//将结构体装在dog类里就继承了animal的属性和方法,伪继承
	//duck-type 
	*animal
}
func (d *dog) speak(){
	fmt.Printf("%s, 叫汪汪汪\n",d.name)
}
func (d *dog) eat(){
	fmt.Printf("%s 在啃 %s\n", d.name, d.food)
}
func main() {
	var d = dog{
		food:"骨头",
		animal: &animal {
			name:"小黄",
		},
	}
	d.speak()
	d.run()
	d.eat()
}