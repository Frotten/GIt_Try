package main

import (
	"fmt"
	"net/http"
	"os"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	Txt, err := os.ReadFile("./Http_Try01.txt")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	_, _ = fmt.Fprintln(w, string(Txt))
}
func main() {
	http.HandleFunc("/Hello", SayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
