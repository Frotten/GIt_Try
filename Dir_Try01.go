package main

import (
	"fmt"
	"os"
)

func main() {
	location := "D:\\TryDir"
	err := os.Mkdir(location, 0777)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("目标已存在")
		}
	}
}
