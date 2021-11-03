package main

import "fmt"

func main() {

	var s1 = []string{"qq", "ww", "rrr"}
	fmt.Println(s1, len(s1), cap(s1))

	var s2 = []int{2, 3}
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s2 == nil) //false

	var s3 = []bool{}
	fmt.Println(s3, len(s3), cap(s3))
	fmt.Println(s3 == nil) //false

	var s4 []string
	fmt.Println(s4, len(s4), cap(s4))
	fmt.Println(s4 == nil) //true

	a := [5]float64{9, 8, 7, 5, 6}
	s5 := a[1:3]
	fmt.Println(s5, len(s5), cap(s5))

	s6 := a[1:4] // 8, 7, 5      3 4
	fmt.Println(s6, len(s6), cap(s6))

	//切片再切
	s7 := s6[1:4] //7 5 6    3 3 // 索引的上限是cap(s)而不是len(s)
	fmt.Println(s7, len(s7), cap(s7))

	//make
	//如果不指定第三个参数cap，则第三个参数和第二个参数一样
	//使用make创建出来的切片，默认不为nil，
	//如果长度不为0，则元素全部初始化为零值
	s8 := make([]float64, 0, 4)
	fmt.Println(s8, len(s8), cap(s8))
	fmt.Println(s8 == nil) //false
	//长度为0但是并不为空指针
	fmt.Println(len(s8) == 0) //true
	//判断一个切片是不是空，要用len函数，而不是指针是否为nil

}
