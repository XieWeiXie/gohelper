package main

import (
	"fmt"
	"sync"
)

func main() {
	// method 1: 实现同步
	OK := make(chan bool)
	go func() {
		// do something
		OK <- true
	}()
	fmt.Println(fmt.Sprintf("Done?: %t!", <-OK))
	method2()
}

func method2() {
	// 实现同步的方法二
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// do something
		}()
	}
	wg.Wait()
	fmt.Println("Done?: method2")
}
