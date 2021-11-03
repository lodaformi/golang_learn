package main

import "fmt"

type person struct {
	name  string
	age   int
	hobby []string
}

func main() {
	//结构体初始化方式一
	var p person
	p.name = "哈哈"
	p.age = 25
	p.hobby = []string{"football", "swimming"}

	fmt.Println(p)

	//结构体初始化方式二
	p2 := person{
		name:  "混合",
		age:   65,
		hobby: []string{"running"},
	}
	fmt.Println(p2)

	//结构体初始化方式三
	p3 := person{
		"tom",
		32,
		[]string{"yoga", "basketball", "baseball"},
	}
	fmt.Println(p3)

	//匿名结构体
	var p4 struct {
		name string
		age  int
	}
	p4.name = "mike"
	p4.age = 35
	fmt.Println(p4)

	//结构体指针
	p5 := new(person)
	(*p5).name = "jerry"
	p5.age = 22 //语法糖，不用明确的写明解地址操作
	fmt.Printf("%T\n", p5)
	fmt.Printf("%v\n", p5)  //p5存放的是新建对象的地址
	fmt.Printf("%p\n", &p5) //p5本身的地址

	p6 := &person{
		name:  "老年人",
		age:   65,
		hobby: []string{"running"},
	}
	fmt.Printf("%T\n", p6)
	fmt.Printf("%#v\n", p6)
}
