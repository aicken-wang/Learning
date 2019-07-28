package main

import "fmt"

func reverse( str string)( string) {
	byteArray := []byte(str)
	strLen := len(byteArray);
	for i := 0; i< (strLen)/2; i++ {
		byteArray[i], byteArray[strLen - 1 - i] = byteArray[strLen - 1- i],byteArray[i]
	}
	return string(byteArray)
}
func main()  {
	s1 := "hello"
	byteArray := []byte(s1)
	s2 := ""
	for i := len(byteArray) -1; i >= 0; i-- {
		s2 = s2 + string(byteArray[i])
	}
	fmt.Println(s2)
	fmt.Println(" reverse('world!') = ",reverse("world!"))
}