package main

import "fmt"

type animal struct {
	name string
	age  int
}

func newAnimal(n string, a int) *animal {
	return &animal{
		name: n,
		age:  a,
	}
}

func (a *animal) wang() {
	fmt.Printf("%s:汪汪汪\n", a.name)
}

func (a *animal) ageIncrese() {
	a.age++
}

type myInt int

func (m *myInt) sayHello() {
	fmt.Printf("%d - hi\n", *m)
}

func main() {
	a1 := newAnimal("啤酒", 3)
	fmt.Printf("%T\n", a1)
	a1.wang()
	fmt.Println(a1.age)
	a1.ageIncrese()
	fmt.Println(a1.age)

	m := myInt(20)
	m.sayHello()
}
