package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PersonHttpTry02 struct {
	Name  string
	Skill []string
	Age   int
}

func FunctionHttp02(w http.ResponseWriter, r *http.Request) { //w用来向客户端发送响应，r用来获取客户端请求的信息，一个写一个读
	Temp, err := template.ParseFiles("./Template.tmpl")
	if err != nil {
		panic(err)
	}
	Owl := PersonHttpTry02{
		Name: "Owl",
		Skill: []string{
			"Fly",
			"Catching",
			"Night Vision",
		},
		Age: 10,
	}
	err = Temp.Execute(w, Owl) //传入结构体，则模板出也视为放入一个结构体，如果在模板中使用了Owl.Name等字段，则会从Owl结构体中获取对应的值
	if err != nil {
		panic(err)
	}
}
func main() {
	http.HandleFunc("/Owl", FunctionHttp02)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP Serve Error", err)
		return
	}
}
