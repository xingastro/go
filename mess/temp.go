package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var result int
	processors := runtime.GOMAXPROCS(0)

	for i := 0; i < processors; i++ {
		fmt.Println("processors =", processors)
		go func(){
			for { result++ }
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("result =", result)
}
