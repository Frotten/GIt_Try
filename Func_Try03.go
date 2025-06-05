package main

import "fmt"

func ClosePackage(Num int) (func(int) int, func(int) int) {
	f1 := func(i int) int {
		Num += i
		return Num
	}
	f2 := func(i int) int {
		Num -= i
		return Num
	}
	return f1, f2
}

func ClosePackage01(Num int) (func(int) int, func(int) int) {
	f1 := func(i int) int {
		Num += i
		return Num
	}
	f2 := func(i int) int {
		Num -= i
		return Num
	}
	return f1, f2
}
func main() {
	Num := 1
	f1, f2 := ClosePackage(Num)   //这里f1,f2共享一个Num变量，和f3,f4不同
	f3, f4 := ClosePackage01(Num) //这里f3,f4共享一个Num变量,和f1,f2不同
	fmt.Println("f1:", f1(2))
	fmt.Println("Num:", Num) // Num的值没有变化，因为这里的Num不是被捕获的Num，而f1,f2和f3,f4共享的Num变量不同
	fmt.Println("f2:", f2(1))
	fmt.Println("Num:", Num)
	fmt.Println("f3:", f3(2))
	fmt.Println("Num:", Num)
	fmt.Println("f4:", f4(1))
	fmt.Println("Num:", Num)
	fmt.Println("f1:", f1(2))
	fmt.Println("Num:", Num)
	fmt.Println("f2:", f2(1))
	fmt.Println("Num:", Num)
	fmt.Println("f3:", f3(2))
	fmt.Println("Num:", Num)
	fmt.Println("f4:", f4(1))
	fmt.Println("Num:", Num)
}
