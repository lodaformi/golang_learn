package main

import "fmt"

func main() {

	s1 := make([]int32, 5, 7)
	s2 := s1 //将s1直接赋值给s2，s1和s2共用一个底层数组

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))

	s1[3] = 30 //底层数组的改变会影响s1和s2切片
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))

	arr := [...]int32{5, 3, 6, 7, 8, 78, 1, 23, 1}
	s3 := arr[2:6]
	var s4 = s3
	fmt.Println(s3, len(s3), cap(s3))
	fmt.Println(s4, len(s4), cap(s4))
	arr[4] = 98
	fmt.Println(s3, len(s3), cap(s3))
	fmt.Println(s4, len(s4), cap(s4))

	//切片的遍历
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%v ", s3[i])
	}
	fmt.Println()

	for i, v := range s4 {
		fmt.Printf("%d - %v\n", i, v)
	}
	fmt.Println()
}
