package main

import "fmt"

type ElementType int32

type EleType = int32

func main() {
	var a ElementType = 100
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var b EleType = 200
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
