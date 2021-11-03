package main

import "fmt"

type speaker interface {
	speak()
}

type dog struct{}

func (d *dog) speak() {
	fmt.Println("汪汪汪汪")
}

type cat struct{}

func (c *cat) speak() {
	fmt.Println("喵喵喵喵")
}

func hit(x speaker) {
	x.speak()
}

func main() {
	c := cat{}
	d := dog{}

	hit(&c)
	hit(&d)
}
