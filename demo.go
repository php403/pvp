package main

import (
	"fmt"
	"net"
)

func main()  {
	listener, _ := net.Listen("tcp", "localhost:8888")
	for  {
		conn, _ := listener.Accept()
		readBuf(conn)
	}
}

func readBuf(coon net.Conn) {
	buf := make([]byte, 512)
	len, err := coon.Read(buf)
	if err != nil {
		fmt.Println("Error reading", err.Error())
		return //终止程序
	}
	fmt.Printf("Received data: %v\n", string(buf[:len]))
}

