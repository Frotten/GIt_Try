package main

//func main() {
//	fileInfo, err := os.Stat("D:\\TryDir\\Try01.txt")
//	var location string
//	if err != nil {
//		fmt.Println("Error:", err)
//	} else {
//		fmt.Println("文件名:", fileInfo.Name())
//		fmt.Println("文件大小:", fileInfo.Size())
//		fmt.Println("是否为目录:", fileInfo.IsDir())
//		fmt.Println("上一次修改时间:", fileInfo.ModTime())
//		fmt.Println("读写权限:", fileInfo.Mode())
//		location = "D:\\TryDir\\Try01.txt"
//	}
//	fmt.Println("父目录：", filepath.Join(location, "..")) // "."为当前目录，”..”为上级目录
//	fmt.Println("父目录：", filepath.Dir(location))
//	//err = os.Mkdir("D:\\TryDir\\TryDirTemp01", os.ModePerm) //创建目录,os.ModePerm表示权限0777
//	if err != nil {
//		panic(err)
//	}
//	//err = os.Remove("D:\\TryDir\\TryDirTemp01") //删除目录
//}
