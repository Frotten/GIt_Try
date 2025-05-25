package main

//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//	go func() {
//		for {
//			select {
//			case <-ctx.Done():
//				fmt.Println("Done")
//				return
//			default:
//				fmt.Println("Working...")
//				time.Sleep(1 * time.Second)
//			}
//		}
//	}()
//	<-ctx.Done()
//	fmt.Println("Cancel main")
//}
