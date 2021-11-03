package main

import (
	"fmt"
	"strings"
)

func main() {
	n := strings.Count("a:b:c:c", "")
	fmt.Println(n)
}
