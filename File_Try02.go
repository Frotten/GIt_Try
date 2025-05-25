package main

//import (
//	"fmt"
//	"io"
//	"os"
//)
//
//func main() {
//	Dir01, _ := os.ReadDir("D:\\TryDir")
//	fmt.Println(Dir01)
//	var name01, name02 string
//	_, _ = fmt.Scan(&name01, &name02)
//	DL, _ := os.Create("D:\\TryDir\\TempCopy\\" + name02 + ".jpg")
//	defer DL.Close()
//	TSrc := "D:\\TryDir\\" + name01 + ".jpg"
//	Src, _ := os.Open(TSrc)
//	defer Src.Close()
//	n, _ := io.Copy(DL, Src)
//	fmt.Println(n)
//}
