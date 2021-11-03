package main

import "fmt"

func main() {
	var dArr1 [3][2]float64 = [3][2]float64{
		{1, 2},
		{3, 4},
	}

	dArr2 := [3][2]string{
		{"bj", "sh"},
		{"cs", "zz"},
		{"hlj", "db"},
	}

	var dArr3 [3][2]string

	dArr3 = [3][2]string{
		{"bj", "sh"},
		{"cs", "zz"},
		{"hlj", "db"},
	}

	var dArr4 = [3][2]string{
		{"bj", "sh"},
		{"cs", "zz"},
		{"hlj", "db"},
	}

	var dArr5 = [...][2]string{
		{"bj", "sh"},
		{"cs", "zz"},
		{"hlj", "db"},
	}

	fmt.Println(dArr1)
	fmt.Println(dArr2)
	fmt.Println(dArr3)
	fmt.Println(dArr4)
	fmt.Println(dArr5)

	for _, v1 := range dArr5 {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Println()
	}
}
