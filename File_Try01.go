package main

//func HandleError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
//func main() {
//	file01, err := os.OpenFile("D:\\TryDir\\Try01.txt", os.O_CREATE|os.O_APPEND, os.ModePerm)
//	HandleError(err)
//	defer file01.Close()
//	RL := []byte("There is a reason")
//	HL := []byte{0, 84, 104, 97, 116, 32, 105, 115, 0}
//	LL := []byte("0123456789Arknights")
//	n, err := file01.Write(RL)
//	HandleError(err)
//	fmt.Println("成功写出：", n)
//	n, err = file01.Write(HL)
//	HandleError(err)
//	fmt.Println("成功写出：", n)
//	n, err = file01.Write(LL[10:])
//	HandleError(err)
//	fmt.Println("成功写出：", n)
//}
