package main

import (
	"net"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.56.1:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := "Hell World" + strconv.Itoa(i) + "\n"
		conn.Write([]byte(msg))
	}
}
