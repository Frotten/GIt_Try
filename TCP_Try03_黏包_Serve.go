package main

import (
	"fmt"
	"net"
)

func main() {
	Listener, err := net.Listen("tcp", "192.168.56.1:8080")
	defer Listener.Close()
	if err != nil {
		panic(err)
	}
	for {
		conn, err := Listener.Accept()
		if err != nil {
			panic(err)
		}
		go func(conn net.Conn) {
			Temp := make([]byte, 1024)
			defer conn.Close()
			for {
				n, err := conn.Read(Temp)
				if err != nil {
					fmt.Println("读取数据失败：", err)
					return
				}
				fmt.Println("接收到：", string(Temp[:n]))
			}
		}(conn)
	}
}
