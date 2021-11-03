package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	fmt.Println(f1()) //  5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //5

	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

//1: x := 1
//2: y := 2
//3: defer calc("AA", 1, calc("A", 1, 2))
//4: calc("A", 1, 2)  // A 1 2 3
//5: defer calc("AA", 1, 3)
//6: x = 10
//7: defer calc("BB", 10, calc("B", 10, 2))
//8: calc("B", 10, 2) // B 10 2 12
//9: defer calc("BB", 10, 12)
//10: y = 20
//11: calc("BB", 10, 12) // BB 10 12 22
//12: calc("AA", 1, 3)   // AA 1 3 4