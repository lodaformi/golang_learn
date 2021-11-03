package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := []float64{3, 4, 5, 1}
	s2 := s1
	s1[2] = 200
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := []float64{3, 4, 5, 1}
	s4 := make([]float64, 5)
	copy(s4, s3)
	fmt.Println(s3)
	fmt.Println(s4)
	s3[1] = 100
	fmt.Println(s3)
	fmt.Println(s4)

	a := [...]int{4, 23, 1, 7, 8, 67, 9}
	fmt.Println(a)
	sort.Ints(a[:])
	fmt.Println(a)
}
