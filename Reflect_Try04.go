package main

import (
	"fmt"
	"reflect"
)

type InterfaceA interface {
	MethodA() string
	MethodB() int
}

type InterfaceB interface {
	MethodC() string
}

type structA struct {
	FieldA string
	FieldB int
}

func (s structA) MethodA() string {
	return s.FieldA
}

func (s structA) MethodB() int {
	return s.FieldB
}

func (s structA) MethodC() string {
	return s.FieldA
}
func main() { //只要你实现了对应的方法，你就是那一类东西
	A := structA{
		FieldA: "It's FieldA",
		FieldB: 42,
	}
	var IA InterfaceA
	IA = A
	fmt.Println(IA.MethodA())
	fmt.Println(IA.MethodB())
	VIA := reflect.ValueOf(IA)
	IVIA := VIA.Interface()
	BIVIA, ok := IVIA.(InterfaceB)
	if ok {
		fmt.Println(BIVIA.MethodC())
	} else {
		fmt.Println(ok)
	}
}
