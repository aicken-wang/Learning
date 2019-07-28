package main

import "fmt"

func main() {
	/*切片三要素
	地址 切片中第一个元素指向的内存空间
	大小 切片中目前元素的个数 len()
	容量 底层数组最大能存放的元素的个数 cap()
	切片支持自动扩容 方式 vector 一样
	append()函数是往切片中追加元s素
	a = append(a, elem)
	copy() 函数复制切片
	*/

	var a = []int{}
	fmt.Printf("a:%v len %d cap:%d ptr:%p \n", a, len(a), cap(a), &a)
	a = append(a, 1)
	fmt.Printf("a:%v len %d cap:%d ptr:%p \n", a, len(a), cap(a), &a)

	var b = [...]int{1, 2, 3}
	a = append(a, 2)
	fmt.Printf("a:%v len %d cap:%d ptr:%p \n", a, len(a), cap(a), &a)
	a = append(a, 3)
	fmt.Printf("a:%v len %d cap:%d ptr:%p \n", a, len(a), cap(a), &a)
	a = append(a, 4)
	var c = [...]int{1, 2, 3}
	fmt.Println(b, c)
	fmt.Printf("a:%v len %d cap:%d ptr:%p \n", a, len(a), cap(a), &a)
	//切片是引用类型
	s1 := []int{1, 2, 3}
	s2 := s1
	var s3 []int
	//分配空间
	s3 = make([]int, len(s1), cap(s1))
	//var s3 = []int{} //空间为0
	copy(s3, s1) //拷贝一份新的数据 深拷贝,返回拷贝数据的len
	fmt.Printf("修改切片前 s1 = %v \t s2 = %v \n", s1, s2)
	s1[0] = 10
	fmt.Printf("修改切片后 s1 = %v \t s2 = %v \n", s1, s2)
	fmt.Printf("s3 = %v S1ptr:%p S3ptr:%p\n", s3, &s1, &s3)
	city := []string{"深圳", "北京", "上海"}
	fmt.Println(city)
	city = append(city, "杭州")
	city = append(city[:1], city[2:]...)
	fmt.Println(city)
	//切片的总容量已经超过1024在追加就0.25倍增加
	//make:用来给引用类型做初始化（申请内存空间）
	//new 用来创建值类型
	//切片声明未初始化不能使用 如 var s []int  s[0] = 100; 执行会报错
	//s = make([]int, 3, 3)
}
