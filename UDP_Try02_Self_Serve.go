package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	Addr, err := net.ResolveUDPAddr("udp", "192.168.56.1:9090")
	if err != nil {
		log.Fatal(err)
		return
	}
	listener, err := net.ListenUDP("udp", Addr)
	if err != nil {
		fmt.Println("Listen Error:", err)
		return
	}
	defer listener.Close()
	for {
		buf := make([]byte, 1024)
		n, des, err := listener.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Read Error:", err)
			continue
		}
		if n == 0 {
			fmt.Println("Read Over")
			continue
		}
		fmt.Println("Received:", string(buf[:n]), " from:", des)
		num, err := listener.WriteToUDP(buf[:n], des)
		if err != nil || num != n {
			fmt.Println("Write Error:", err)
		} else {
			fmt.Println("Write Success:", num)
		}
	}
}
