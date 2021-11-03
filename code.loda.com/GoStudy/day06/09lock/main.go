package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var lock sync.Mutex
var x int64

func add1() {
	defer wg.Done()
	for i := 0; i < 5000000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	add1()
	add1()
	wg.Wait()

	fmt.Println(x)
}
