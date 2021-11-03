package main

import "fmt"

func main() {
	m := make(map[string]int, 5)
	fmt.Println(m)
	fmt.Println(m["中国"])

	//元素为map的切片
	s1 := make([]map[string]int, 2, 5)
	//长度要不为0，否则panic: runtime error: index out of range [0] with length 0
	s1[0] = make(map[string]int, 2) //要对其中的map进行初始化，否则panic: assignment to entry in nil map
	s1[0]["长沙"] = 2021
	s1[0]["岳麓区"] = 985
	fmt.Println(s1)

	//元素为切片的map
	m2 := make(map[string][]int, 2)
	m2["橘子洲"] = []int{1, 3, 4, 5}
	fmt.Println(m2)

}
