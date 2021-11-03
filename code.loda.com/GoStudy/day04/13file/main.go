package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFilefun1() {
	fileObj, err := os.Open("main.go")
	if err != nil {
		fmt.Printf("open file error, %v\n", err)
		return
	}
	defer fileObj.Close()

	var content []byte
	b := make([]byte, 128)
	for {
		n, err := fileObj.Read(b)
		if err == io.EOF {
			fmt.Println("读取到文件末尾")
			break
		}
		if err != nil {
			fmt.Printf("read file error, %v\n", err)
			return
		}
		fmt.Printf("读取了%d个字符\n", n)

		// fmt.Println(string(b[:n])) //对b进行切片是因为当n小于128时
		content = append(content, b[:n]...)
		// if n < 128 {
		// 	fmt.Printf("读取了%d个字符\n", n)
		// 	fmt.Println(string(b))
		// 	return
		// }
	}
	fmt.Println(string(content))
}

func bufioReadFile() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file error, %v\n", err)
		return
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读取到末尾")
			break
		}
		if err != nil {
			fmt.Printf("读取错误,%v\n", err)
			return
		}
		fmt.Print(line)
	}
}

func ioutilReadFile() {
	content, err := ioutil.ReadFile("main.go")
	if err != nil {
		fmt.Printf("open file error, %v\n", err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	// readFilefun1()
	// bufioReadFile()
	ioutilReadFile()
}
