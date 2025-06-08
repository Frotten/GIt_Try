package main

type stu struct {
	Age  int
	Name string
}

func (s stu) getName() string {
	return s.Name
}
func (s *stu) setAge(age int) {
	s.Age = age
}

func (s stu) getAge() int {
	return s.Age
}
func main() {
	s := stu{Age: 16, Name: "Gopher"}
	stu.getName(s)
	(*stu).setAge(&s, 20)
	(*stu).getName(&s)
	(*stu).setAge(&s, 30)
	s.getAge()
}
