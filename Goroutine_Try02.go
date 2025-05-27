package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var Htt = []string{
	"https://pkg.go.dev/",
	"https://qq.com/",
}

func main() {
	for _, value := range Htt {
		wg.Add(1)
		go func(value string) {
			defer wg.Done()
			rest, err := http.Get(value)
			if err == nil {
				fmt.Println(value, ":", rest.Status)
			}

		}(value)
	}
	wg.Wait()
}
