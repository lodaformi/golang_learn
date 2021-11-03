package main

import "fmt"

type animal struct {
	name string
}

func (a *animal) move() {
	fmt.Printf("%s move\n", a.name)
}

type dog struct {
	feet uint8
	animal
}

func (d *dog) wang() {
	fmt.Printf("%s wang\n", d.name)
}

func main() {
	d := dog{feet: 4, animal: animal{name: "啤酒"}}
	d.move()
	d.wang()
}
