package main

import (
	"fmt"
	"unicode"
)

var (
	num2   int64   = 567
	f2     float64 = 56.99
	status bool    = false
	str2   string  = "星期五Friday"
)

func main() {
	// num := 15
	// f1 := 13.56
	// isOk := true
	// str := "hello"

	// fmt.Printf("%v %T\n", num, num)
	// fmt.Printf("%v %T\n", f1, f1)
	// fmt.Printf("%v %T\n", isOk, isOk)
	// fmt.Printf("%v %T\n", str, str)

	// fmt.Printf("%d %T\n", num2, num2)
	// fmt.Printf("%f %T\n", f2, f2)
	// fmt.Printf("%v %T\n", status, status)
	// fmt.Printf("%s %T\n", str2, str2)

	// const (
	// 	c1 = iota //0
	// 	c2        //1
	// 	c3 = 66
	// 	c4
	// )
	// fmt.Printf("%d %T\n", c1, c1)
	// fmt.Printf("%d %T\n", c2, c2)
	// fmt.Printf("%d %T\n", c3, c3)
	// fmt.Printf("%d %T\n", c4, c4)
	// fmt.Printf("%d %T\n", iota, iota)

	var str3 string = "hello小王子哈哈哈"
	var chCount int64 = 0
	// var asCount int64 = 0
	for _, c := range str3 {
		if unicode.Is(unicode.Han, c) {
			chCount++
		}
		// if unicode.Is(unicode.ASCII_Hex_Digit, c) {
		// 	asCount++
		// }
		// fmt.Printf("%v, %T\n", reflect.TypeOf(c), reflect.TypeOf(c))
	}
	fmt.Printf("汉字有%d个\n", chCount)
	// fmt.Printf("汉字有%d个，ASCII字符有%d个\n", chCount, asCount)

	//九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
