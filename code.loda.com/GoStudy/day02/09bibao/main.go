package main

import "fmt"

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

func f3(f func(int, int), a, b int) func() {
	return func() {
		f(a, b)
	}
}

func main() {
	f3(f2, 3, 4)
}
