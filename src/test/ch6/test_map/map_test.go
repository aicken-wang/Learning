package test_map

import "testing"

/*
Map 声明
m := map[string]int{"one":1,"two":2, "three":3}
m1 := map[string]int{}
m1["one"] = 1
m2 := make(map[string]int,10)// Initial capacity
//为啥不初始化len?
*/
func TestInitMap(t *testing.T)  {
	m := map[string]int{"one":1,"two":2, "three":3}
	t.Log(m["one"])
	t.Logf("Len = %d",len(m))
	m["one"] = 1
	m0 := map[int]int{}
	t.Logf("len m0 = %d ", len(m0))
	m1 := make(map[string]int,10)
	t.Logf("Len m1 = %d",len(m1))
}
func TestAccessNotExistingKey( t *testing.T)  {
	m := map[int]int{}
	t.Log(m[2])
	m[2] = 0
	t.Log(m[2])
	if v, ok := m[3]; ok{
		t.Logf( "Key 3's value is %d", v)
	} else {
		t.Log("Key 3's not existing.")
	}
}
/**
	Map 元素的访问
与其他主要编程语言的差异
在访问的Key不存在时，仍会返回零值，不能通过返回nil来判断元素是否存在
*/
//遍历 Map
func TestTravelMap(t *testing.T) {
	m := map[string]int{"one":1,"two":2, "three":3}
	for k, v := range m {
		t.Log(k, v)
	}
}