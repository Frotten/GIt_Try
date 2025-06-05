package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) //捕获到panic，输出错误信息
		}
	}()
	defer func() {
		panic("First time?") //重复触发
	}()
	defer func() {
		panic("Second time?") //触发到又一个panic，继续寻找recover
	}()
	panic("This time") //Panic触发，开始沿着栈向上搜寻recover
}
