package main

//func HandleError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
//func main() {
//	BasicRoutine := "D:\\TryDir\\"
//	TarDir := BasicRoutine + "TempCopy\\"
//	//err := os.Mkdir(BasicRoutine+"TempCopy", 0777)
//	//HandleError(err)
//	Sour := BasicRoutine + "白咕咕.jpg"
//	Dest := Sour[strings.LastIndex(Sour, "\\")+1:]
//	fmt.Println(Sour)
//	fmt.Println(Dest)
//	file01, err := os.Open(Sour)
//	HandleError(err)
//	file02, err := os.OpenFile(TarDir+Dest, os.O_CREATE|os.O_RDWR, os.ModePerm)
//	HandleError(err)
//	file03, err := os.OpenFile(TarDir+Dest+"Temp.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
//	HandleError(err)
//	defer file01.Close()
//	defer file02.Close()
//	defer file03.Close()
//	file03.Seek(0, io.SeekStart)
//	TempSlice := make([]byte, 1024)
//	n, err := file03.Read(TempSlice)
//	CompleteInt, err := strconv.ParseInt(string(TempSlice[:n]), 10, 64)
//	fmt.Println("已完成字节数：", CompleteInt)
//	file01.Seek(CompleteInt, io.SeekStart)
//	file02.Seek(CompleteInt, io.SeekStart)
//	data := make([]byte, 1024)
//	for {
//		//if CompleteInt >= 270000 {
//		//	panic("突然中断")
//		//}
//		n, err := file01.Read(data)
//		if err == io.EOF || n == 0 {
//			fmt.Println("复制结束,现累积复制了", CompleteInt, "字节内容")
//			file03.Close()
//			os.Remove(TarDir + Dest + "Temp.txt")
//			break
//		}
//		n1, _ := file02.Write(data[:n])
//		CompleteInt += int64(n1)
//		TempByte := []byte(strconv.Itoa(int(CompleteInt)))
//		file03.Seek(0, io.SeekStart)
//		file03.Write(TempByte)
//	}
//}
