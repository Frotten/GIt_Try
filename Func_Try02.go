package main

import (
	"fmt"
	"math/rand"
	"time"
)

func AnyFuncFuncTry02(a, b int, c func(int, int) int) int {
	return c(a, b)
}

func AddFuncTry02(a, b int) int {
	return a + b
}

func SubFuncTry02(a, b int) int {
	return a - b
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	a, b := rand.Intn(100)+1, rand.Intn(100)+1
	fmt.Println("a:", a, "b:", b)
	fmt.Println("Add:", AnyFuncFuncTry02(a, b, AddFuncTry02))
	fmt.Println("Sub:", AnyFuncFuncTry02(a, b, SubFuncTry02))
}
