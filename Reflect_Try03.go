package main

import (
	"fmt"
	"reflect"
)

type Robert struct {
	id       int
	Skill    []string
	Deadline int
}

func main() {
	RobertA := Robert{
		id: 1,
		Skill: []string{
			"Clean",
			"Kill",
			"Steal",
		},
		Deadline: 2025525,
	}
	VA := reflect.ValueOf(&RobertA)
	TA := reflect.TypeOf(&RobertA)
	ETA := TA.Elem()
	EVA := VA.Elem()
	fmt.Println("ETA:", ETA, "  EVA:", EVA)
	TempVA := VA.Interface()
	TempTrack, ok := TempVA.(*Robert)
	if ok {
		fmt.Println(TempTrack)
		fmt.Println(TempTrack.Skill)
	} else {
		fmt.Println("Error")
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover, err:", err)
		}
	}()
	fmt.Println(VA.Float())
}
