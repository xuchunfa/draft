package main

import (
	"fmt"
	"time"
)

func main() {

	// map同时删除/访问
	//lock := sync.RWMutex{} //加读写锁去访问
	var m = map[string]int{"1": 1, "2": 2, "3": 3}
	go func() {
		//lock.RLock()
		//defer lock.RUnlock()
		for k, v := range m {
			fmt.Println(k, v)
		}
	}()
	go func() {
		//lock.Lock()
		//defer lock.Unlock()
		for k := range m {
			delete(m, k)
		}
	}()
	time.Sleep(time.Second)
}
