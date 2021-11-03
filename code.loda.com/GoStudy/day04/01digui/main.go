package main

import "fmt"

//阶乘
func f1(n uint32) uint32 {
	if n == 1 {
		return 1
	}
	return n * f1(n-1)
}

//what
//上台阶
func f2(n uint32) uint32 {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return f2(n-1) + f2(n-2)
}

func main() {
	fmt.Println(f1(5))

	fmt.Println(f2(11))
}
