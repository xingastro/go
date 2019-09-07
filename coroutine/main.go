package main

import (
	"time"
)

func main() {
	a := [10]int{}
	for i := 0; i < 10; i++ {
		go func (i int) {
			for {
				//fmt.Println("Hello GOlang from goroutine, ", i)
				a[i]++
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	//fmt.Println(a)
}
