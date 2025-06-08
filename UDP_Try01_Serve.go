package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "192.168.56.1:8080")
	if err != nil {
		panic(err)
	}
	listen, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Listen Fielded")
		return
	}
	defer listen.Close()
	for {
		buf := make([]byte, 1024)
		n, destination, err := listen.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("Received from", destination.String(), " Data:", string(buf[:n]))
		_, err = listen.WriteToUDP(buf[:n], destination)
		if err != nil {
			fmt.Println("Receive Failed:", err)
		}
	}
}
