package main

import (
	"fmt"
)

func show(in interface{}) {
	val, ok := in.(string)
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("guess type error")
	}
}

func show2(in interface{}) {
	switch val := in.(type) {
	case string:
		fmt.Println(val)
	case int:
		fmt.Println(val)
	case bool:
		fmt.Println(val)
	case float32:
		fmt.Println(val)
	default:
		fmt.Println("guess type error")
	}
}

func main() {

	show2("ccc")
	show2(false)
}
