package main

import "fmt"

// var name string
// var age int
// var isOk bool

var (
	name string
	age  int
	isOk bool = false
)

func main() {
	var school string = "bj"
	name = "tom"
	age = 20
	isOk = true

	fmt.Printf("name : %s", name)
	fmt.Println(age)
	fmt.Print(isOk)
	fmt.Println(school)
}
