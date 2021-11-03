package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func fileWrite1() {
	file, err := os.OpenFile("file.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file error, %v\n", err)
		return
	}
	defer file.Close()
	file.Write([]byte("哈哈哈sunday\n"))

	file.WriteString("星期天O(∩_∩)O哈哈~\n")
}

func bufioWriteFile() {
	file, err := os.OpenFile("file.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file error, %v\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("嘻嘻嘻呵呵呵\n")
	writer.Flush()
}

func ioutilWriteFile() {
	err := ioutil.WriteFile("test.txt", []byte("sssssfa\nssssshhhhhh"), 0644)
	if err != nil {
		fmt.Printf("open file error, %v\n", err)
		return
	}
}

//ioutil.WriteFile简化成了os.WriteFile
func osWriteFile() {
	err := os.WriteFile("test.txt", []byte("起的隆冬器a\nssssshhhhhh"), 0644)
	if err != nil {
		fmt.Printf("open file error, %v\n", err)
		return
	}
}

func main() {
	// fileWrite1()
	// bufioWriteFile()
	// ioutilWriteFile()
	osWriteFile()
}
