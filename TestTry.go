package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := "Hell World"
	b := a[:]
	b = "Fantastic World"
	fmt.Println(a, b)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))
}
