package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

type calculation func(int, int) int

func main() {
	var a calculation
	a = add
	fmt.Println(a(1, 2))
	fmt.Printf("%T\n", a)

	b := add
	fmt.Println(b(1, 2))
	fmt.Printf("%T\n", b)

	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)

	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

}
