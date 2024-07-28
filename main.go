package main

import (
	"fmt"
	"time"
)

func main() {

	// context
	//ctx, cancel1 := context.WithTimeout(context.Background(), time.Second)
	//defer cancel1()
	//
	//ctx, cancel2 := context.WithCancel(context.TODO())
	//defer cancel2()
	//
	//ctx = context.WithValue(context.Background(), "trace", "1")
	//ctx.Value("trace")
	//
	//cpus := runtime.GOMAXPROCS(0)
	//fmt.Println(cpus)

	// 断言
	//var i interface{}
	//i = 5
	//res := i.(int)
	//fmt.Println(res)

	// goroutine
	//for i := 0; i < 5; i++ {
	//	//j := i
	//	go func(j int) {
	//		fmt.Println(j)
	//	}(i)
	//}
	//
	//time.Sleep(time.Second)

	timer := time.NewTimer(time.Second)
	defer timer.Stop()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timer.C:
			fmt.Println("timer end")
			timer.Reset(time.Second)
		case <-ticker.C:
			fmt.Println("ticker end")
		}
	}
}
