package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//将处理字符串的功能单独拿出来，方便使用goroutine处理
func procsMsg(conn net.Conn) {
	defer conn.Close()
	var msg [123]byte
	reader := bufio.NewReader(os.Stdin)

	for {
		//读消息
		n, err := conn.Read(msg[:])
		// fmt.Println("read n, ", n)
		if err != nil {
			fmt.Println("read failed, ", err)
			return
		}
		fmt.Println(string(msg[:n]))

		//回消息
		//向连接写入消息
		fmt.Print("reply msg: ")
		msg, _ := reader.ReadString('\n')
		msg = strings.Trim(msg, "\r\n") //windows在读取字符后，会加上\r\n，最好去掉，这样读取到的字符个数才是你输入的
		// fmt.Println("msg ", msg)
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("write error, ", err)
		}
		// fmt.Println("n ", n)
	}
}

func server() {
	listener, err := net.Listen("tcp", "127.0.0.1:50050")
	if err != nil {
		fmt.Println("server failed, ", err)
		return
	}
	defer listener.Close()
	//持续接收不同客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, ", err)
			return
		}
		go procsMsg(conn)
	}

}

func main() {
	server()
}
