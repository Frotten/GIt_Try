package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[:3]
	fmt.Println("lenA :", len(a), "capA :", cap(a)) //lenA : 6 capA : 6
	fmt.Println("lenB :", len(b), "capB :", cap(b)) //lenB : 3 capB : 6
	b = append(b, 10, 11, 12)
	fmt.Println(a)            // [1 2 3 10 11 12] --b的append操作会影响到a，因为b是a的切片
	fmt.Println(b)            // [1 2 3 10 11 12]
	b = append(b, 13, 14)     //此时b超出容量，系统会重新分配内存，ab的底层数组不再共享
	fmt.Println(a)            //[1 2 3 10 11 12]
	fmt.Println(b)            //[1 2 3 10 11 12 13 14]
	a = append(a, 4, 5, 6, 7) //可以发现a，b的底层数组不再共享
	fmt.Println(a)            //[1 2 3 10 11 12 4 5 6 7]
	fmt.Println(b)            //[1 2 3 10 11 12 13 14]
}
