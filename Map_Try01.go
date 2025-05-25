package main

//type Person struct {
//	Name    string
//	Age     int
//	Sex     string
//	Address string
//}

//func main() {
//	ATemp := make(map[string]Person)
//	var Tip string
//	AllName := make([]string, 0)
//	fmt.Println("请输入姓名：")
//	for fmt.Scan(&Tip); Tip != "exit"; fmt.Scan(&Tip) {
//		var age int
//		var sex, address string
//		fmt.Println("请输入年龄，性别和地址：")
//		fmt.Scan(&age, &sex, &address)
//		ATemp[Tip] = Person{
//			Name:    Tip,
//			Age:     age,
//			Sex:     sex,
//			Address: address,
//		}
//		AllName = append(AllName, Tip)
//		fmt.Println("请输入姓名：")
//	}
//	sort.Strings(AllName)
//	fmt.Println("通过名字进行排序后的map：")
//	for _, Name := range AllName {
//		fmt.Println(Name, ":")
//		fmt.Println(ATemp[Name].Age, ATemp[Name].Sex, ATemp[Name].Address)
//	}
//	fmt.Println("没排过序的map:")
//	for _, TemPerson := range ATemp {
//		fmt.Println(TemPerson.Name, TemPerson.Age, TemPerson.Sex, TemPerson.Address)
//	}
//}
