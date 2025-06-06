package main

import (
	"fmt"
	"math/rand"
)

func randIntA() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Intn(100) + 1
		}
	}()
	return ch
}
func randIntB() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Intn(100) + 1
		}
	}()
	return ch
}
func randIntTry04(done chan bool) chan int {
	ch := make(chan int, 10)
	go func() {
	Label:
		for {
			select {
			case ch <- <-randIntA():
			case ch <- <-randIntB():
			case <-done:
				break Label
			}
		}
		close(ch)
	}()
	return ch
}
func main() {
	done := make(chan bool)
	ch := randIntTry04(done)
	fmt.Println("RandIntA：", <-ch)
	fmt.Println("RandIntB：", <-ch)
	close(done)
	for v := range ch {
		fmt.Println("RandIntA or RandIntB：", v)
	}
}
