package main

//func safeMayPanic() {
//	defer func() {
//		if r := recover(); r != nil {
//			fmt.Println("Recover Completely")
//		}
//	}()
//	fmt.Println("Starting")
//	panic("Something wrong")
//	//这里不会被执行
//}
//
//func main() {
//	fmt.Println("Starting")
//	safeMayPanic()
//	fmt.Println("Ending")
//}
