package main

import (
	"fmt"
	"sort"
)

type PersonSort struct {
	Name string
	Age  int
}

type Byage []PersonSort

func (a Byage) Len() int           { return len(a) }
func (a Byage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Byage) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []PersonSort{
		{"Alex", 66},
		{"Steve", 55},
		{"Him", 404},
	}
	sort.Sort(Byage(people))
	fmt.Println("By name:", people)
	people = []PersonSort{
		{"Alex", 18},
		{"Blame", 16},
		{"Nell", 19},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Println("By name:", people)
}
