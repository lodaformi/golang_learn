package main

import (
	"fmt"
	"net"
	"strings"
)

func procMsg(conn net.UDPConn) {
	fmt.Println("in procMsg func")
	var recMsg [1024]byte
	for {
		//接收消息
		n, udpAddr, err := conn.ReadFrom(recMsg[:])
		if err != nil {
			fmt.Println("server read error, ", err)
			return
		}
		fmt.Println(string(recMsg[:n]))

		//回消息
		senMsg := strings.ToUpper(string(recMsg[:n]))
		conn.WriteTo([]byte(senMsg), udpAddr)
	}
}

func server() {
	// net.Listen()
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 50050,
		Zone: "ipv4",
	})
	if err != nil {
		fmt.Println("listen udp error, ", err)
		return
	}
	// defer conn.Close()
	// go procMsg(*conn) //为什么开启goroutine就不行了？？？
	procMsg(*conn)
}

func main() {
	server()
}
