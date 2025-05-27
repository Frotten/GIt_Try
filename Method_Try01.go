package main

import "fmt"

type Int int

func (a Int) Prin() {
	fmt.Println(a)
}

func (a *Int) Set(b int) {
	*a = Int(b)
}

func main() {
	A := Int(1)
	B := Int(2)
	A.Prin()
	(&A).Set(3)
	A.Prin()
	B.Prin()
	B.Set(4)
	(&B).Prin()
}
