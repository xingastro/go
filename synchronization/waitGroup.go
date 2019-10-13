package main

import (
	"fmt"
	"sync"
)

func dowork(wg *sync.WaitGroup) {
	fmt.Println("hello world")
	wg.Done()
}

func main() {
	var wg  = new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go dowork(wg)
	}
	wg.Wait()
}