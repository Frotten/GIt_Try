package main

//func sum() {
//	sum := 0
//	for i := 0; i < 100000; i++ {
//		sum++
//	}
//	fmt.Println("sum02 : ", sum)
//	time.Sleep(time.Second * 1)
//}
//func main() {
//	go func() {
//		sum := 0
//		for i := 0; i < 100000; i++ {
//			sum++
//		}
//		fmt.Println("sum01 : ", sum)
//		time.Sleep(1 * time.Second)
//	}()
//	go sum()
//	fmt.Println("当前程序的goroutine数目:", runtime.NumGoroutine())
//	time.Sleep(3 * time.Second)
//}
