package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randIntTry03() chan int {
	ch := make(chan int)
	go func() {
		for {
			ch <- rand.Intn(100) + 1
		}
	}()
	return ch
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	ch1 := randIntTry03()
	fmt.Println("随机数01:", <-ch1)
	fmt.Println("随机数02:", <-ch1)
}
