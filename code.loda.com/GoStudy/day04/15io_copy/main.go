package main

import (
	"fmt"
	"io"
	"os"
)

func myCopy() {
	src_file, err := os.OpenFile("src.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("open src file error, %v\n", err)
		return
	}
	defer src_file.Close()

	dst_file, err := os.OpenFile("dst.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open dst file error, %v\n", err)
		return
	}
	defer dst_file.Close()

	_, err = io.Copy(dst_file, src_file)
	if err != nil {
		fmt.Printf("copy file error, %v\n", err)
		return
	}
}
func main() {
	myCopy()
}
