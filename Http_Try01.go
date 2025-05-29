package main

import (
	"fmt"
	"net/http"
	"os"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	Txt, err := os.ReadFile("./Http_Try01.txt") // 读取当前目录下的Http_Try01.txt文件
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound) // 如果读取文件失败，返回404错误
		return
	}
	_, _ = fmt.Fprintln(w, string(Txt))
}
func main() {
	http.HandleFunc("/Hello", SayHello)      //注册一个处理函数，当访问/Hello路径时调用SayHello函数
	err := http.ListenAndServe(":9090", nil) //通过ListenAndServe函数来启动一个HTTP服务器，监听9090端口
	if err != nil {
		fmt.Println(err)
		return
	}
}
