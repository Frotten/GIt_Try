package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func FunctionHttp02(w http.ResponseWriter, r *http.Request) { //w用来向客户端发送响应，r用来获取客户端请求的信息，一个写一个读
	Temp, err := template.ParseFiles("./Template.tmpl")
	if err != nil {
		panic(err)
	}
	err = Temp.Execute(w, "这是传入的数据")
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
