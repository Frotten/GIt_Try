package main

//func copyFile(Dest, Sour string) (int, error) {
//	file1, err := os.Open(Sour)
//	if err != nil {
//		return 0, err
//	}
//	file2, err := os.OpenFile(Dest, os.O_CREATE|os.O_WRONLY, os.ModePerm)
//	if err != nil {
//		return 0, err
//	}
//	defer file1.Close()
//	defer file2.Close()
//	BS := make([]byte, 1024)
//	var n, total int
//	for {
//		n, err = file1.Read(BS)
//		if err == io.EOF || n == 0 {
//			break
//		} else if err != nil {
//			fmt.Println("报错")
//			return total, err
//		}
//		total += n
//		file2.Write(BS[:n])
//	}
//	return total, nil
//}
//
//func copyFile2(Dest, Sour string) (int64, error) {
//	file1, err := os.Open(Sour)
//	if err != nil {
//		return 0, err
//	}
//	file2, err := os.OpenFile(Dest, os.O_CREATE|os.O_WRONLY, os.ModePerm)
//	if err != nil {
//		return 0, err
//	}
//	defer file1.Close()
//	defer file2.Close()
//	return io.Copy(file2, file1)
//}
//func main() {
//	file1 := "D:\\TryDir\\Try01.txt"
//	file2 := "D:\\TryDir\\Try02.txt"
//	start1 := time.Now()
//	n, err := copyFile(file2, file1)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("函数1成功拷贝：", n)
//	end1 := time.Now()
//	fmt.Println("拷贝函数1拷贝所消耗的时间：", end1.Sub(start1))
//	start2 := time.Now()
//	file1 = "D:\\TryDir\\白咕咕.jpg"
//	file2 = "D:\\TryDir\\TempCopy.jpg"
//	nl, err := copyFile2(file2, file1)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("函数2成功拷贝：", nl)
//	end2 := time.Now()
//	fmt.Println("拷贝函数2拷贝所需要的时间：", end2.Sub(start2))
//	file1 = "D:\\TryDir\\Try02.txt"
//	file, err := os.OpenFile(file1, os.O_RDONLY, os.ModePerm)
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//	_, err = file.Seek(11, io.SeekStart)
//	if err != nil {
//		panic(err)
//	}
//	BS := make([]byte, 0, 40)
//	Temp := make([]byte, 1)
//	A := []byte{0}
//	for {
//		n, err := file.Read(Temp)
//		if err == io.EOF || n == 0 || bytes.Contains(Temp, []byte(" ")) || bytes.Contains(Temp, A) {
//			fmt.Println("Break了")
//			break
//		} else if err != nil {
//			log.Fatal(err)
//		}
//		BS = append(BS, Temp...)
//		fmt.Println("当前读取：", string(Temp))
//	}
//	fmt.Println("读取到的值：", string(BS))
//}
