package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func Decode(reader *bufio.Reader) (string, error) {
	LengthOfInformation, _ := reader.Peek(4)             //在不移动指针的情况下查看前n个字节的内容
	LengthBuffer := bytes.NewBuffer(LengthOfInformation) //将前4个字节的内容转换为Buffer类型
	var length int32
	err := binary.Read(LengthBuffer, binary.LittleEndian, &length) //将前4个字节转换为int32类型的长度信息
	if err != nil {
		return "", nil
	}
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	Temp := make([]byte, length+4) //+4是因为前面有4个字节的长度信息
	_, err = reader.Read(Temp)
	if err != nil {
		return "", nil
	}
	return string(Temp[4:]), nil
}

func processTry04(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		Information, err := Decode(reader)
		if err != nil {
			fmt.Println("Decode Error:", err)
			return
		}
		if Information == "" {
			fmt.Println("No Information Received")
			return
		}
		fmt.Println("Received:", Information)
		_, err = conn.Write([]byte(Information))
		if err != nil {
			fmt.Println("Write Error:", err)
			return
		}
	}
}
func main() {
	Listener, err := net.Listen("tcp", "192.168.56.1:8080")
	if err != nil {
		panic(err)
	}
	defer Listener.Close()
	for {
		conn, err := Listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go processTry04(conn)
	}
}
