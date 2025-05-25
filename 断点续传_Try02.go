package main

//func main() {
//	filename01 := "D:\\TryDir\\白咕咕.jpg"
//	filename02 := "D:\\TryDir\\TempCopy\\临时白咕咕.jpg"
//	filename03 := filename02 + ".txt"
//	file01, _ := os.OpenFile(filename01, os.O_RDONLY, os.ModePerm)
//	defer file01.Close()
//	file02, _ := os.OpenFile(filename02, os.O_CREATE|os.O_RDWR, os.ModePerm)
//	defer file02.Close()
//	file03, _ := os.OpenFile(filename03, os.O_CREATE|os.O_RDWR, os.ModePerm)
//	defer file03.Close()
//	TempSlice := make([]byte, 1024)
//	n, _ := file03.Read(TempSlice)
//	CompleteInt, _ := strconv.ParseInt(string(TempSlice[:n]), 10, 64)
//	file01.Seek(CompleteInt, io.SeekStart)
//	file02.Seek(CompleteInt, io.SeekStart)
//	data := make([]byte, 2048)
//	for {
//		//if CompleteInt >= 102400 {
//		//	panic("中断")
//		//}
//		n, err := file01.Read(data)
//		if n == 0 || err == io.EOF {
//			fmt.Println("复制完毕")
//			file03.Close()
//			_ = os.Remove(filename03)
//			break
//		} else {
//			n1, _ := file02.Write(data[:n])
//			CompleteInt += int64(n1)
//			TempByte := []byte(strconv.Itoa(int(CompleteInt)))
//			file03.Seek(0, io.SeekStart)
//			file03.Write(TempByte)
//		}
//	}
//}
