package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func process02(conn net.Conn) {
	defer conn.Close()
	for {
		Reader := bufio.NewReader(conn)
		Temp := make([]byte, 1024)
		n, err := Reader.Read(Temp)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Read Over")
			} else {
				fmt.Println(err)
			}
			return
		}
		fmt.Println("接收到信息：", string(Temp[:n]))
		_, err = conn.Write(Temp[:n]) //将接收到的信息重新写入，这是为了和客户端进行回复，作为确认
		if err != nil {
			fmt.Println("Write Failed:", err)
			return
		}
	}
}
func main() {
	Listener, err := net.Listen("tcp", "192.168.56.1:8080")
	if err != nil {
		fmt.Println("Listen Fielded")
		panic(err)
	}
	defer Listener.Close()
	for {
		conn, err := Listener.Accept()
		if err != nil {
			fmt.Println("Accept Fielded", err)
			continue
		}
		go process02(conn)
	}
}
