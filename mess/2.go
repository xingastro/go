package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			if i == 5 {
				runtime.Gosched()
			}
			fmt.Println(i)
		}(i)
	}
}


func main() {
	ch := make(chan int)

	go func(){
		for i := 0; i <= 3; i++ {
			ch <- i
		}
		close(ch) // 关闭掉channel，在使用for range时就不会死锁
	}()

	for i := range ch {
		fmt.Println(i)
	}
}
