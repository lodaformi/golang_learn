package main

import (
	"encoding/json"
	"fmt"
)

//字段名必须大写，因为若是小写，在json的包中无法获取到相应字段
//可以在后面加上tag，注意是反引号``，跟上解析为相应格式时的字段名字
type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	//结构体 -> json
	p1 := person{Name: "张飞", Age: 35}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("error %v\n", err)
	}
	// fmt.Printf("%#v\n", b) //[]byte{0x7b,.... 0x7d}
	// fmt.Printf("%#v\n", string(b))
	fmt.Printf("%v\n", string(b)) //返回值b为[]byte，可以强转为string

	//json -> 结构体
	s := `{"name":"赵子龙", "age":26}` //使用反引号定义json格式的字符串
	var p2 person
	json.Unmarshal([]byte(s), &p2) //将字符串强转为[]byte，传入的是结构体指针
	fmt.Println(p2)
}
