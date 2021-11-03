package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func show(in int) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println("hello", in)
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go show(i)
	}
	wg.Wait()
}
