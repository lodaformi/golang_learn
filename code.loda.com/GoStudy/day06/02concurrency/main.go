package main

import (
	"fmt"
	"time"
)

func hello(num int) {
	fmt.Println("hello", num)
}

func main() {
	for i := 0; i < 100; i++ {
		// go hello(i)
		//go routine滞后执行，而外部的i已经变化
		//出现一些线程打印出相同值的情况
		go func() {
			fmt.Println("anonymous", i)
		}()

		go func(in int) {
			fmt.Println("anonymous", in)
		}(i)
	}
	fmt.Println("main func")
	time.Sleep(time.Second)
}
