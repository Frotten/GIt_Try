package main

//func main() {
//	TL := make(chan int)
//	var wg sync.WaitGroup
//	go func() {
//		wg.Add(1)
//		defer wg.Done()
//		defer close(TL)
//		for i := 1; i <= 10; i++ {
//			TL <- i
//		}
//	}()
//	defer wg.Wait()
//	for value := range TL {
//		fmt.Println("value,敬请见证:", value)
//	}
//}
