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
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		input = strings.TrimSpace(input)
		if strings.ToUpper(input) == "EXIT" {
			return
		}
		_, err = conn.Write([]byte(input))
		if err != nil {
			panic(err)
		}
		response := make([]byte, 1024)
		n, err := conn.Read(response)
		if err != nil {
			fmt.Println("Read Failed:", err)
			return
		}
		fmt.Println(string(response[:n]))
	}
}
