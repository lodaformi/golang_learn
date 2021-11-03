package main

import "fmt"

func main() {

	//变量
	var num int = 5
	var num2 = 10
	num3 := 30
	fmt.Println(num)
	fmt.Println(num2)
	fmt.Println(num3)

	//一维数组
	var arr1 [3]int = [3]int{4, 56, 7}
	var arr2 = [3]int{9, 8, 1}
	arr3 := [5]int{2, 3}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	//二维数组
	var darr = [4][3]float32{
		{1, 2},
		{2, 3},
	}
	fmt.Println(darr)

	for _, i := range darr {
		for _, j := range i {
			fmt.Printf("%v ", j)
		}
		fmt.Println()
	}

	fmt.Println("hello world")
}
