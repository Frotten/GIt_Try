package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	Addr, err := net.ResolveUDPAddr("udp", "192.168.56.1:9090")
	if err != nil {
		log.Fatal(err)
	}
	Conn, err := net.DialUDP("udp", nil, Addr)
	if err != nil {
		log.Fatal(err)
	}
	defer Conn.Close()
	for {
		Temp, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Read Error:", err)
			continue
		}
		Temp = strings.TrimSpace(Temp)
		if Temp == "" || strings.ToUpper(Temp) == "EXIT" {
			fmt.Println("Read Over")
			break
		}
		_, err = Conn.Write([]byte(Temp))
		if err != nil {
			fmt.Println("Write Error:", err)
			continue
		}
		Res := make([]byte, 1024)
		n, err := Conn.Read(Res)
		if err != nil {
			fmt.Println("Read Error:", err)
			continue
		}
		fmt.Println("回响：", string(Res[:n]))
	}
}
