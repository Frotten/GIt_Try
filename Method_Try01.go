package main

import "fmt"

type Int int

func (a Int) PrinMethod() {
	fmt.Println(a)
}

func (a *Int) SetMethod(b int) {
	*a = Int(b)
}

func main() {
	A := Int(1)
	B := Int(2)
	A.PrinMethod()
	(&A).SetMethod(3)
	A.PrinMethod()
	B.PrinMethod()
	B.SetMethod(4)
	(&B).PrinMethod()
	Int.PrinMethod(A)
	(*Int).SetMethod(&B, 5)
	Int.PrinMethod(B)
}
