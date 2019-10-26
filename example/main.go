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
		fmt.Println("in main")
		OK <- true
	}()

	method2()
	fmt.Println(fmt.Sprintf("Done?: %t!", <-OK))
}

func method2() {
	// 实现同步的方法二
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("in method2")
			// do something
		}()
	}
	wg.Wait()
	fmt.Println("Done?: method2")
}
