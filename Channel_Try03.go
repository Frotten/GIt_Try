package main

//func Add1(A chan int) chan int {
//	B := make(chan int)
//	go func() {
//		for v := range A {
//			B <- v + 1
//		}
//		close(B)
//	}()
//	return B
//}
//func main() {
//	A := make(chan int)
//	go func() {
//		for i := 1; i <= 10; i++ {
//			A <- i
//		}
//		close(A)
//	}()
//	B := Add1(Add1(Add1(A)))
//	for v := range B {
//		fmt.Printf("第%d个+3后的数是：%d\n", v-3, v)
//	}
//}
