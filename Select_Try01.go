package main

//func main() {
//	c1 := make(chan int)
//	c2 := make(chan int)
//	c3 := make(chan bool)
//	go func() {
//		fmt.Println("c1:", <-c1)
//		c3 <- true
//	}()
//	go func() {
//		c2 <- 10
//		fmt.Println("c2<-10")
//		c3 <- true
//	}()
//	select {
//	case c1 <- 1:
//		fmt.Println("c1:", c1)
//	case Tempnum := <-c2:
//		fmt.Println("c2:", Tempnum)
//	case num := <-time.After(3 * time.Second):
//		fmt.Println("timeout", num)
//	}
//	<-c3
//}
