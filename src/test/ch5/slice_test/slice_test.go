package slice_test

import "testing"

func TestSliceInit(t *testing.T)  {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))
	//方式二
	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))
	//方式三
	//[]type ,len,cap 其中len 个元素会被初始化为默认零值，未初始化元素不可以访问
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1],s2[2] ,s2[3])
}
/**
	切片共享存储结构
*/
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++  {
		s = append(s, i)
		t.Log(len(s),cap(s))
	}
}
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May","Jun", "Jul","Aug","Sep","Oct","Nov","Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2),cap(Q2))
	summer := year[5 :8]
	t.Log(summer,len(summer), cap(summer))
	summer[0] = "UnKnow"
	t.Log(Q2)
	t.Log(year)
}
/*
数组 vs 切片
1.容量是否可伸缩
2.是否可以进行比较
结论：
数组不可伸缩 ，可比较
切片可伸缩，不可比较
*/
