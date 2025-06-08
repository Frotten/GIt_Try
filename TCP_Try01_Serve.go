package main

import (
	"bufio"
	"fmt"
	"net"
)

func process01(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [1024]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("Error reading :", err)
			return
		}
		TarStr := string(buf[:n])
		fmt.Println("Received:", TarStr)
		conn.Write([]byte(TarStr)) //将信息重新写回到客户端
	}
}

func main() { //TCP服务器端尝试
	listen, err := net.Listen("tcp", "192.168.56.1:8080") //使用Listen在指定IP上监听TCP链接
	if err != nil {
		fmt.Println("Listen Fielded")
		panic(err)
	}
	for {
		conn, err := listen.Accept() //通过Accept来接收链接，如果接收成功则会将连接信息返回给conn
		if err != nil {
			fmt.Println("Accept Fielded", err)
			continue
		}
		go process01(conn)
	}
}
