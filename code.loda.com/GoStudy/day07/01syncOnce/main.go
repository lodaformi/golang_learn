package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func f_rec(ch chan<- int) {
	// defer wg.Done()
	for i := 0; i < 100; i++ {
		ch <- i
	}
	// close(ch)
}

func f_send(ch_r <-chan int, ch_w chan<- int) {
	// defer wg.Done()
	for {
		i, ok := <-ch_r // 通道关闭后再取值ok=false
		if !ok {
			break
		}
		ch_w <- i * i
	}
	// close(ch_w)
	// once.Do(func() { close(ch_w) })
}

// channel 练习
func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	// 开启goroutine将0~100的数发送到ch1中
	// wg.Add(1)
	go f_rec(ch1)
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	// wg.Add(1)
	go f_send(ch1, ch2)
	// go f_send(ch1, ch2)

	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}

	// for {
	// 	v, ok := <-ch2
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }
	// wg.Wait()

	fmt.Println("main end")
}
