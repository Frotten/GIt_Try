package main //目的，写一个能从0 加到 100 的程序

//const (
//	NUMBER = 10
//)
//
//type task struct {
//	start  int
//	end    int
//	result chan<- int
//}
//
//func (t *task) calculate() {
//	sum := 0
//	for i := t.start; i <= t.end; i++ {
//		sum += i
//	}
//	t.result <- sum
//}
//
//func divide(worker chan task, res chan int, nums int) {
//	mod := nums % NUMBER
//	limit := nums / NUMBER
//	for i := 0; i < limit; i++ {
//		down := (i * NUMBER) + 1
//		up := (i + 1) * NUMBER
//		tsk := task{
//			start:  down,
//			end:    up,
//			result: res,
//		}
//		worker <- tsk
//	}
//	if mod != 0 {
//		tsk := task{
//			start:  (nums/NUMBER)*NUMBER + 1,
//			end:    nums,
//			result: res,
//		}
//		worker <- tsk
//	}
//	close(worker)
//}
//
//func action(worker chan task, sinal chan bool) {
//	for i := 0; i < NUMBER; i++ {
//		go func(worker chan task, sinal chan bool) {
//			for t := range worker {
//				t.calculate()
//			}
//			sinal <- true
//		}(worker, sinal)
//	}
//}
//
//func finish(sinal chan bool, res chan int) {
//	for i := 0; i < NUMBER; i++ {
//		<-sinal
//	}
//	close(res)
//	close(sinal)
//}
//func collect(res chan int) int {
//	sum := 0
//	for value := range res {
//		sum += value
//	}
//	return sum
//}
//func main() {
//	timeHead := time.Now().UnixNano()
//	worker := make(chan task, NUMBER)
//	res := make(chan int, NUMBER)
//	sinal := make(chan bool, NUMBER)
//	go divide(worker, res, 100)
//	action(worker, sinal)
//	go finish(sinal, res)
//	fmt.Println("sum:", collect(res))
//	fmt.Println("Using_time:", time.Now().UnixNano()-timeHead)
//}
