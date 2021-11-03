package main

import "fmt"

func main() {
	num := 15
	ptr := &num

	var ptr2 **int = &ptr

	fmt.Printf("%T - %p\n", ptr, ptr)
	fmt.Println(*ptr)

	fmt.Printf("%T - %p\n", ptr2, ptr2)
	fmt.Println(*ptr2)
	fmt.Println(**ptr2)

	ptr3 := new(*int)
	fmt.Println(ptr3 == nil)
	*ptr3 = ptr
	fmt.Printf("%T - %p\n", ptr3, ptr3)
	fmt.Println(*ptr3)
	fmt.Println(**ptr3)

}
