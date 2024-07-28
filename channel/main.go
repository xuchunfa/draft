package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	reqNums := 10
	//控制协程并发数量
	nums := make(chan struct{}, 2)
	for i := 0; i < reqNums; i++ {
		wg.Add(1)
		nums <- struct{}{}
		go func(i int) {
			defer func() {
				wg.Done()
				<-nums
			}()
			fmt.Printf("goroutine_num: %d, go func: %d \n", runtime.NumGoroutine(), i)
		}(i)
	}
	wg.Wait()

	// 模拟time.After内存泄漏
	//for {
	//	select {
	//	case <-time.After(time.Second):
	//		fmt.Println(time.Now())
	//	default: // 加上default 非阻塞写
	//	}
	//}

	// 往nil的channel读/写都会阻塞当前协程
	//ch := make(chan int)
	//ticker := time.NewTicker(1 * time.Second).C
	//select {
	//case <-ch:
	//case <-time.After(2 * time.Second):
	//	fmt.Println(time.Now())
	////case ch <- 1:
	////default: // 加上default 非阻塞写
	//case <-ticker:
	//	fmt.Println("ticker at: ", time.Now())
	//}

	//done := make(chan struct{})
	//ch := make(chan int)
	////生产者
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		ch <- i
	//	}
	//	close(ch)
	//}()
	////消费者
	//go func() {
	//	for i := range ch {
	//		//time.Sleep(1 * time.Second)
	//		fmt.Println("consume: ", i)
	//	}
	//	done <- struct{}{}
	//}()
	//<-done
	//fmt.Println(rand.Intn(5))
}
