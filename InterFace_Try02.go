package main

import "fmt"

type Characteristic interface {
	Habit() []string
	Food() string
	Position() string
}

type Animal struct {
	Name string
	Age  int
}

type Dog struct {
	Animal
	Species string
}

type Cat struct {
	Animal
	IQ int
}

func (d Dog) Habit() []string {
	return []string{"Bark", "Run", "Fetch"}
}

func (d Dog) Food() string {
	return "Meat"
}

func (d Dog) Position() string {
	return "Human's best friend"
}

func (c Cat) Habit() []string {
	return []string{"老吴", "耄耋挠头.jpg", "秒开棘背龙形态.jpg"}
}

func (c Cat) Food() string {
	return "Fish"
}

func (c Cat) Position() string {
	return "Master or Pet"
}

func Show(c Characteristic) {
	fmt.Println("接下来你将要见证的，是一个哈吉米的全部：")
	fmt.Println(c.Habit())
	fmt.Println(c.Food())
	fmt.Println(c.Position())
}
func main() {
	Dog1 := Dog{
		Animal: Animal{
			Name: "大黄",
			Age:  114514,
		},
		Species: "邪魅的笑",
	}
	Cat1 := Cat{
		Animal: Animal{
			Name: "耄耋",
			Age:  17,
		},
		IQ: 1,
	}
	fmt.Println(Dog1.Food(), "And", Cat1.Food())
	fmt.Println("才艺欣赏：", Dog1.Species, "And", Cat1.Habit())
	Show(Dog1)
	Show(Cat1)
}
