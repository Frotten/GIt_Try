package main

import (
	"fmt"
	"reflect"
)

type AnimalReflectTry05 struct {
	Name  string
	Nums  int
	Skill []string
}

func main() {
	Dog := AnimalReflectTry05{
		Name: "Dog",
		Nums: 1,
		Skill: []string{
			"Run",
			"Wolf",
		},
	}
	fmt.Println(Dog)
	VD := reflect.ValueOf(&Dog)
	EVD := VD.Elem()
	if EVD.CanSet() {
		EVD.FieldByName("Nums").SetInt(3)
		DSkill := EVD.FieldByName("Skill")
		DSkill.Set(reflect.ValueOf([]string{
			"Run",
			"Wolf",
			"Search",
			"Guard",
		}))
	} else {
		fmt.Println("Can't Set")
	}
	fmt.Println(Dog)
}
