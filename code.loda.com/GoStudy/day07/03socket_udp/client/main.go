package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func client() {
	conn, err := net.Dial("udp", "127.0.0.1:50050")
	if err != nil {
		fmt.Println("dial udp error, ", err)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	var cont [1024]byte
	for {
		//写消息
		fmt.Println("please input: ")
		msg, err := reader.ReadString('\n')
		msg = strings.Trim(msg, "\r\n")
		if err != nil {
			fmt.Println("reader read msg error, ", err)
			return
		}

		conn.Write([]byte(msg))

		//回消息
		n, err := conn.Read(cont[:])
		if err != nil {
			fmt.Println("client read msg error, ", err)
		}
		fmt.Println(string(cont[:n]))
	}

}

func main() {
	client()
}
