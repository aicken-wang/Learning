package string_test

import (
	"testing"
	"unsafe"
)

/*
字符串
与其他主要编程语言的差异
1.string 是数据类型，不是引用或指针类型
2. string是只读 byte slice ,len 函数是它所包含的byte数
3.string的byte数组可以存放任何数据
*/
func TestString(t *testing.T){
	var s string
	t.Log(s) //初始化默认零值""
	s = "hello"
	t.Log(len(s))
	//s[1] = "3" //string是不可变byte slice
	s = "\347\216\213"//可以存储任何二进制数据
	//s = "\xE4\xB8\xA5"
	//s ="\xE4\xB8\xA5\xFF"
	t.Log(s)
	s = "中"
	t.Log(len(s)) // byte数
	c :=[]rune(s)
	t.Log("rune size:",unsafe.Sizeof(c[0]))
	t.Logf("中 Unicode %x", c[0])
	t.Logf("中 UTF %x",s)

}
/*
strings 包 https://golang.org/pkg/strings/
strconv 包 https://golang.org/pkg/strconv/
*/
func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range  s {
		t.Logf(" %[1]c %[1]d",c)
	}
}