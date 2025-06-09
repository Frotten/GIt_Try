package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func Encode(message string) ([]byte, error) {
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		fmt.Println("binary.Write Fielded")
		return nil, err
	}
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		fmt.Println("binary.Write Fielded")
		return nil, err
	}
	return pkg.Bytes(), nil
}
func main() {
	Listen, err := net.Dial("tcp", "192.168.56.1:8080")
	if err != nil {
		panic(err)
	}
	defer Listen.Close()
	for {
		String := make([]string, 1)
		fmt.Scan(&String[0])
		TarByte, err := Encode(String[0])
		if err != nil {
			fmt.Println("Error Reading:", err)
		}
		_, err = Listen.Write(TarByte)
		if err != nil {
			fmt.Println("Write Error:", err)
			return
		}
		response := make([]byte, 1024)
		n, err := Listen.Read(response)
		if err != nil {
			fmt.Println("Read Error:", err)
			continue
		}
		fmt.Println("Read Response:", string(response[:n]))
	}
}
