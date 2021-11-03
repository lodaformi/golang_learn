package main

import "fmt"

func main() {
	var s1 []int32
	s1 = append(s1, 4, 5, 2, 1)
	fmt.Println(s1)

	s2 := []float32{} //这种方式定义切片，相当于初始了零值
	fmt.Println(len(s2), cap(s2), s2 == nil)

	s3 := make([]string, 5) //这种方式定义切片，相当于初始了零值
	fmt.Println(s3, len(s3), cap(s3), s3 == nil)

	s4 := []string{"qq", "河南"}
	s3 = append(s3, s4...)
	fmt.Println(s3, len(s3), cap(s3), s3 == nil)

	var s5 []int32
	for i := 0; i < 10; i++ {
		s5 = append(s5, int32(i))
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", s5, len(s5), cap(s5), s5)
	}

	a := [...]int32{6, 7, 8, 3, 2, 1}
	s6 := a[:]
	fmt.Println(s6, len(s6), cap(s6))
	s6 = append(s6[:2], s6[3:]...)
	fmt.Println(s6, len(s6), cap(s6))
	fmt.Println(a, len(a), cap(a))

	s6[1] = 121
	fmt.Println(s6)
	fmt.Println(a)
}
