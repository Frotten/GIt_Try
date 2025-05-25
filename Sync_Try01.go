package main

//var wg sync.WaitGroup
//var lk sync.RWMutex
//var Temp int = 1
//
//func ReadData(n int) {
//	defer wg.Done()
//	lk.RLock()
//	fmt.Println(n, ": Start Reading")
//	fmt.Println(n, ": Reading :", Temp)
//	defer lk.RUnlock()
//}
//func WriteData(n int) {
//	defer wg.Done()
//	lk.Lock()
//	fmt.Println(n, ": Start Writing : ", Temp)
//	defer lk.Unlock()
//	Temp = n
//}
//func main() { //一般情况下会先将大部分读的协程进行完毕，再进行写的协程，因为读操作之间不会互相锁止，但写操作会
//	wg.Add(20)
//	for i := 1; i <= 10; i++ {
//		go ReadData(i)
//		go WriteData(i)
//	}
//	wg.Wait()
//	fmt.Println("Finished,Temp is :", Temp)
//}
