package main

import (
	"fmt"
	"reflect"
)

func main() {
	A := "A"
	fmt.Println("TypeOf(A):", reflect.TypeOf(A), " Kind of A :", reflect.TypeOf(A).Kind())
	B := 'B'
	fmt.Println("TypeOf(B):", reflect.TypeOf(B), " Kind of B :", reflect.TypeOf(B).Kind())
	var C rune = 'C'
	fmt.Println("TypeOf(C):", reflect.TypeOf(C), " Kind of C :", reflect.TypeOf(C).Kind())
	VC := reflect.ValueOf(C).Interface()
	if VVC, ok := VC.(int32); ok {
		fmt.Println(int(VVC), string(VVC))
	} else {
		fmt.Println(ok)
	}
}
