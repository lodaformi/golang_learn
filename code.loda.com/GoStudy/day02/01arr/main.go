package main

import "fmt"

func main() {
	//定义时初始化
	var arr1 [5]int16 = [5]int16{9, 8, 7, 3, 6}

	//如果长度已定，可以使用索引初始化指定位置的值
	//未指定的位置初始化为默认零指
	var arr2 = [5]int16{0: 9, 3: 6}

	//如果长度不定可以使用...代替数组长度，由编译器推断
	arr3 := [...]int16{3, 4, 5, 2, 65, 7, 0, 9, 12, 8, 342, 53, 4, 25}

	//先定义，后初始化
	var arr4 [3]float32
	arr4 = [3]float32{3, 4, 7}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	sum := 0
	arr5 := [...]int64{1, 3, 5, 7, 8}
	for _, v := range arr5 {
		sum += int(v)
	}

	fmt.Printf("sum is %d\n", sum)

	for i, v := range arr5 {
		for j := i + 1; j < len(arr5); j++ {
			if (v + arr5[j]) == 8 {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
}
