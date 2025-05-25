package main

//var wg sync.WaitGroup
//
//func ReadOnly(T chan<- int, data int) {
//	defer wg.Done()
//	T <- data
//	fmt.Println("read", data, "complete")
//}
//func WriteOnly(T <-chan int) {
//	defer wg.Done()
//	data := <-T
//	fmt.Println("write", data, "complete")
//}
//func main() {
//	ch := make(chan int)
//	wg.Add(2)
//	go ReadOnly(ch, 5)
//	go WriteOnly(ch)
//	wg.Wait()
//	fmt.Println("main over")
//}
