package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func client() {
	//连接server端
	conn, err := net.Dial("tcp", "127.0.0.1:50050")
	if err != nil {
		fmt.Println("connection failed, ", err)
		return
	}
	defer conn.Close()

	//持续向连接中写消息
	var msg [123]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		//向连接写入消息
		fmt.Print("please input msg: ")
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\r\n") //windows在读取字符后，会加上\r\n，最好去掉，这样读取到的字符个数才是你输入的
		// fmt.Println("msg ", msg)
		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Println("write error, ", err)
		}
		// fmt.Println("n ", n)

		n, err := conn.Read(msg[:])
		// fmt.Println("read n, ", n)
		if err != nil {
			fmt.Println("read failed, ", err)
			return
		}
		fmt.Println(string(msg[:n]))
	}
}

func main() {
	client()
}
