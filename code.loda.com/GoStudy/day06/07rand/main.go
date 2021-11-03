package main

import (
	"fmt"
	"math/rand"
)

func randFunc() {
	for i := 0; i < 5; i++ {
		// fmt.Println(rand.Intn(10))
		fmt.Println(rand.Int63())
	}
}

func main() {
	randFunc()
}
