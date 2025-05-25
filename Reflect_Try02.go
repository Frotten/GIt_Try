package main

//type Person struct {
//	Name  string
//	Age   int
//	Skill []string
//}
//
//func (p Person) Printinfo() {
//	fmt.Println("Name:", p.Name)
//	fmt.Println("Age:", p.Age)
//	fmt.Println("Skill:", p.Skill)
//}
//
//func (p Person) None(A []string, B int, C bool) {
//	if C == true {
//		fmt.Println(A, B)
//	} else {
//		fmt.Println(len(A), B, A)
//	}
//}
//
//func Func01() {
//	fmt.Println("纯正无参函数")
//}
//func Func02(a int, b string) {
//	fmt.Println("纯正有参函数")
//}
//func Func03(a int, b string) string {
//	fmt.Println("纯正有参有返回值函数")
//	return b + strconv.Itoa(a)
//}
//func main() {
//	a := 1.123
//	V := reflect.ValueOf(&a)
//	ele := V.Elem()
//	fmt.Println("Type:", ele.Type(), "Change ? :", ele.CanSet())
//	ele.SetFloat(3.1415926535)
//	fmt.Println(a)
//	fmt.Println("----------------")
//
//	Bob := Person{
//		Name: "Bob",
//		Age:  999,
//		Skill: []string{
//			"Shoot",
//			"Run",
//			"Uppercut",
//		},
//	}
//	EB := reflect.ValueOf(&Bob).Elem()
//	fmt.Println("Can change ? :", EB.CanSet())
//	name := EB.FieldByName("Name")
//	name.SetString("Doomsday Iron Fist")
//	Age := EB.FieldByName("Age")
//	Age.SetInt(18)
//	skill := EB.FieldByName("Skill")
//	skill.Set(reflect.ValueOf([]string{"Uppercut", "Jump", "Fly"}))
//	fmt.Println(Bob)
//	fmt.Println("----------------")
//
//	Anna := Person{
//		Name: "Anna",
//		Age:  60,
//		Skill: []string{
//			"Aim",
//			"Strengthen",
//			"Decrease",
//		},
//	}
//	EA := reflect.ValueOf(Anna)
//	MEA := EA.MethodByName("Printinfo")
//	MEA.Call(nil)
//	NEA := EA.MethodByName("None")
//	NEA.Call([]reflect.Value{reflect.ValueOf([]string{"It's Life"}), reflect.ValueOf(1), reflect.ValueOf(true)})
//	fmt.Println("----------------")
//
//	FEA := reflect.ValueOf(Func01)
//	fmt.Println(FEA.Type(), FEA.Kind())
//	FEA.Call(nil)
//
//	FEB := reflect.ValueOf(Func02)
//	fmt.Println(FEB.Type(), FEB.Kind())
//	FEB.Call([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf("Ni hao")})
//
//	FEC := reflect.ValueOf(Func03)
//	fmt.Println(FEC.Type(), FEC.Kind())
//	fmt.Println(FEC.Type().Name())                                                   //这里什么也不会输出，因为这里要输出的是函数类型的名称，而函数类型没有名称
//	TR := FEC.Call([]reflect.Value{reflect.ValueOf(10086), reflect.ValueOf("欢迎致电")}) //Call的返回值是一个value切片类型
//	fmt.Printf("%T\n%T\n", TR, TR[0])
//	fmt.Println(TR[0].Interface())
//}
