package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "192.168.56.1:8080")
	if err != nil {
		panic(err)
	}
	listen, err := net.DialUDP("udp", nil, addr) //三个参数分别代表协议类型，本地地址和远程地址
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client Closed")
			} else {
				fmt.Println("Read Error:", err)
				continue
			}
		}
		input = strings.TrimSpace(input)
		if strings.ToUpper(input) == "EXIT" {
			fmt.Println("Exiting Client")
			return
		}
		_, err = listen.Write([]byte(input))
		if err != nil {
			fmt.Println("Write Error:", err)
			continue
		}
		response := make([]byte, 1024)
		n, err := listen.Read(response)
		if err != nil {
			fmt.Println("无应答")
			return
		}
		fmt.Println("回应：", string(response[:n]))
	}
}
