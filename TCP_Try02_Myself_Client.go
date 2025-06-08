package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.56.1:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		Reader := bufio.NewReader(os.Stdin)
		TempString, err := Reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		TempString = strings.TrimSpace(TempString)
		if strings.ToUpper(TempString) == "EXIT" {
			break
		}
		_, err = conn.Write([]byte(TempString))
		if err != nil {
			panic(err)
		}
		response := make([]byte, 1024)
		n, err := conn.Read(response)
		if err != nil {
			fmt.Println("Read Error", err)
			break
		}
		fmt.Println("收到回复：", string(response[:n]))
	}
}
