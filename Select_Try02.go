package main

//
//func randInt(ch1 chan bool) chan int {
//	ch2 := make(chan int)
//	go func() {
//	Lable:
//		for {
//			select {
//			case <-ch1:
//				break Lable
//			case ch2 <- rand.Intn(100) + 1:
//			}
//		}
//		close(ch2)
//	}()
//	return ch2
//}
//func main() {
//	rand.New(rand.NewSource(time.Now().UnixNano()))
//	ch1 := make(chan bool)
//	fmt.Println("Rand1:", <-randInt(ch1))
//	fmt.Println("Rand2:", <-randInt(ch1))
//	close(ch1)
//	fmt.Println("Rand3:", <-randInt(ch1))
//	fmt.Println("Rand4:", <-randInt(ch1))
//}
