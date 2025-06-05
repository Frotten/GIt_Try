package main

import (
	"fmt"
	"sync"
	"time"
)

func GoroutineTry06Func01(Arr []int, wg *sync.WaitGroup, ch1, ch2 chan int) {
	defer wg.Done()
	for _, value := range Arr {
		ch1 <- value
		if value%2 == 0 {
			ch2 <- value
		}
	}
	close(ch1)
	close(ch2)
}
func ChanFunc01(wg *sync.WaitGroup, ch1 chan int) {
	defer wg.Done()
	for value := range ch1 {
		fmt.Println("1-----", value)
	}
	time.Sleep(100 * time.Millisecond)
}
func ChanFunc02(wg *sync.WaitGroup, ch2 chan int) {
	defer wg.Done()
	for value := range ch2 {
		fmt.Println("2-----", value)
	}
	time.Sleep(1000 * time.Millisecond)
}
func main() {
	A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup
	wg.Add(3)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go GoroutineTry06Func01(A, &wg, ch1, ch2)
	go ChanFunc01(&wg, ch1)
	go ChanFunc02(&wg, ch2)
	wg.Wait()
}
