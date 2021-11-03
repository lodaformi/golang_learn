package main

import "fmt"

//匿名字段，即结构体中字段的生命不要名字
//带来的问题：因为字段不能相同，所以不能出现相同的数据类型，局限性较大
type person1 struct {
	string
	int
	// int
}

//嵌套结构体
type workPlace struct {
	province string
	city     string
}

type bornPlace struct {
	province string
	city     string
}

type person2 struct {
	name string
	age  int
	// wp   workPlace
	workPlace
	bornPlace
}

func main() {
	p := person1{"tom", 18}
	fmt.Println(p)

	p2 := person2{"哈哈", 23, workPlace{"北京", "丰台"}, bornPlace{"湖南", "长沙"}}
	fmt.Println(p2)
	fmt.Println(p2.wp.province)
	fmt.Println(p2.wp.city)
	fmt.Println(p2.province)
	fmt.Println(p2.city)
	fmt.Println(p2.bornPlace.province)
	fmt.Println(p2.bornPlace.city)

}
