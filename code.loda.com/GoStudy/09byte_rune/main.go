package main

import "fmt"

func main() {
	str1 := "hello星期五"
	// for i := 0; i < len(str1); i++ {
	// 	// fmt.Println(str1[i])		//输出是数字
	// 	// fmt.Printf("%c ", str1[i]) //有乱码
	// 	// fmt.Printf("%v ", str1[i]) //输出是数字
	// }

	// for i, c := range str1 {
	// 	fmt.Printf("%d, %c\n", i, c)
	// }
	// fmt.Println()

	for _, c := range str1 {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}
