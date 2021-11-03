package main

import "fmt"

func main() {
	const (
		n1 = 3
		n2
		n3
	)

	const (
		a1 = iota //0
		a2        //1
		a3 = 34   //34
		a4        //34
		a5        //34
	)

	const (
		b1 = iota //0
		b2        //1
		_         //2
		b3        //3
	)

	const (
		c1 = iota //0
		c2 = 6    //6
		c3 = iota //2
		c4        //3
	)

	const (
		d1, d2 = iota + 1, iota + 2 //d1=1 d2=2
		d3, d4 = iota + 1, iota + 2 //d3=2 d4=3
	)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)
}
