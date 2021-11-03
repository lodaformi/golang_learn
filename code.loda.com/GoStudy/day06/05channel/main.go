package main

import "fmt"

func main() {
	ch1 := make(chan int, 5)

	for i := 10; i < 15; i++ {
		ch1 <- i
	}
	close(ch1)

	// for {

	// 	ret, ok := <-ch1
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(ret)
	// }
	fmt.Println("---------------")
	for i := range ch1 {
		fmt.Println("first ", i)
		// fmt.Println("second ", i)
		one := <-ch1
		fmt.Println("one ", one)
		two := <-ch1
		fmt.Println("two ", two)

	}

	ret, ok := <-ch1
	fmt.Println(ret, ok)
}
