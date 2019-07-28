package objects_test

import (
	"fmt"
	"testing"
	"unsafe"
)

/*
对数据的封装和行为
结构体的定义
type Employe struct {
 Id string
 Name string
 Age int
}
实例创建及初始化
e := Employee { "0", "Bob", 20}
e1 := Employee { Name: "Mike", Age: 30}
e2 := new(Employee) //注意这里返回的引用/指针，相当于 e := &Employee{}
e2.Id = "2" //与其他主要编程语言的差异，通过实例的指针访问成员不需要使用->
e2.Age = 22
e2.Name = "Aicken"
*/
type Employee struct {
	Id string
	Name string
	Age int
}
func (e Employee) String() string {
	//Address is c0000766a0
	fmt.Printf("Address is %x \n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf(" ID: %s-Name:%s-Age:%d",e.Id, e.Name, e.Age)
}
func (e *Employee) StringPtr() string {
	//Address is c000076610
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}
func TestCreateEmployeeObj( t *testing.T)  {
	e := Employee{ "0","Aicken",20 }
	e1 := Employee{"1", "Mike", 30}
	e2 := new (Employee) //返回指针
	e2.Id = "2"
	e2.Name = "King"
	e2.Age = 18
	t.Log(e)
	t.Log(e1)
	t.Log(e.Name)
	t.Log(e2)
	//%T 是实例的类型
	t.Logf("e is %T", e)
	//加上 &e is *objects_test.Employee
	t.Logf("&e is %T", &e)
	//返回指针类型 e2 is *objects_test.Employee
	t.Logf("e2 is %T", e2)
}
/**
行为（方法）定义
与其他主要编程语言的差异
结构体的定义
type Employe struct {
 Id string
 Name string
 Age int
}
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//第一种定义方式在实例对应方法被调用时，实例的成员会进行值复制
func (e Employee) String() string {
	return fmt.Sprintf(" ID: %s-Name:%s-Age:%d",e.Id, e.Name, e.Age)
}
//通常情况下为了避免内存拷贝我们使用第二种定义方式
func (e *Employee) String() string {
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d, e.Id, e.Name, e.Age")
}
*/

func TestStructOperations(t *testing.T)  {

	e := Employee{ "0","wang",27}
	//1 Address is c000076670
	fmt.Printf(" 1 Address is %x \n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
	
}
func TestStructOperationsPtr(t *testing.T)  {
	e := Employee{ "0","wang",27}
	//Address is c000076610
	fmt.Printf("Address is %x \n", unsafe.Pointer(&e.Name))
	t.Log(e.StringPtr())

}
