package main

import "fmt"

func f1() (r int) { //r的位置在返回值栈区，由于这里的func是闭包引用，所以栈区上的r会修改
	defer func() {
		r++
	}()
	return 0
}
func f2() (r int) { //这里存在一个返回值栈区上的r和func中的局部变量r，由于参数等原因这里defer中的函数操作对象是局部变量的r，栈区上的r不受影响
	r = 0
	defer func(a int) {
		a++
	}(r)
	return r //在这一步将局部变量r赋值给返回值栈区的r，之后执行对局部变量进行操作的匿名函数
}
func f3() (r int) { //这里的原理类似于f2，同样是局部变量和栈区变量的关系
	t := 0
	defer func(a int) {
		a += 5
	}(t)
	return t
}

func main() {
	fmt.Println(f1()) // 1
	fmt.Println(f2()) // 0
	fmt.Println(f3()) // 0
}
