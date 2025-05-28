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
	return []string{"mio", "耄耋挠头.jpg", "秒开棘背龙形态.jpg"}
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
	大黄 := Dog{
		Animal: Animal{
			Name: "大黄",
			Age:  114514,
		},
		Species: "邪魅的笑",
	}
	耄耋 := Cat{
		Animal: Animal{
			Name: "耄耋",
			Age:  17,
		},
		IQ: 1,
	}
	fmt.Println(大黄.Food(), "And", 耄耋.Food())
	fmt.Println("才艺欣赏：", 大黄.Species, "And", 耄耋.Habit())
	Show(大黄)
	Show(耄耋)
}
