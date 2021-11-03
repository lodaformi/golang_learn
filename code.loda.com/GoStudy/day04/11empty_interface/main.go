package main

import "fmt"

func show(in interface{}) {
	fmt.Printf("type:%T value:%v\n", in, in)
}

func main() {
	//在map中值的类型使用空接口，表示值可以是任意类型
	m1 := make(map[string]interface{}, 10)
	m1["张飞"]=18
	m1["湖南"]="长沙"
	m1["pass"]=false
	fmt.Println(m1)

	//在map中key的类型使用空接口，表示key可以是任意类型
	m2 := make(map[interface{}]int32, 10)
	m2[10010] = 9099
	m2["学号"] = 12005
	m2[false] = 123
	fmt.Println(m2)


	show("国庆人多哪也不去")
	show(100)
}