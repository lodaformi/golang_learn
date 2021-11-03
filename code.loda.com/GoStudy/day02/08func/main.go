package main

import "fmt"

//有参数无返回值
func sum(x int, y int) {
	fmt.Println(x + y)
}

//有参数有无命名返回值
func sum2(x int, y int) int {
	return x + y
}

//有参数有命名返回值
func sum3(x int, y int) (res int) {
	res = x + y
	return
}

//多个返回值
func f1(x int, str string) (int, string) {
	return x, str
}

//多个命名返回值
func f2(x int, str string) (a int, b string) {
	a = x
	b = str
	return
}

//可变长参数
func f3(x int, y ...string) {
	fmt.Println(x)
	fmt.Println(y)
}

//参数类型简写
func f4(x, y int, a, b, c string) {
	fmt.Println(x, y)
	fmt.Println(a, b, c)
}

func main() {
	sum(2, 3)
	num := sum2(4, 5)
	fmt.Println(num)
	num2 := sum3(5, 6)
	fmt.Println(num2)

	a, b := f1(19, "what")
	fmt.Println(a, b)
	_, d := f2(10, "www")

	f2(10, "www")

	fmt.Println(d)
	f3(100)
	f3(100, "cs")
	f3(100, "cs", "长沙", "湖南省会")

	f4(1, 2, "cs", "长沙", "湖南省会")
}
