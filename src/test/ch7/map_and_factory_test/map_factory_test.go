package map_and_factory_test

import (
	"testing"
)

/*
Map 与工厂模式
Map的value 可以是一个方法
与Go的 Dock type接口方式一起，可以方便的实现单一方法对象的工厂模式
*/
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op  int)int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](1), m[2](2), m[3](2))
}
/*
实现 Set
Go 的内置集合中没有Set的实现 ，可以map[type]bool
1.元素的唯一性
2.基本操作
	1）添加元素
	2）判断元素是否存在
	3）删除元素
	4）元素个数
*/
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d  is existing", n)
	} else {
		t.Logf(" %d is not existing", n)
	}
	t.Log(len(mySet))
	delete(mySet,1)
	t.Log(len(mySet))
	n = 1
	if mySet[n] {
		t.Logf("%d  is existing", n)
	} else {
		t.Logf(" %d is not existing", n)
	}
}