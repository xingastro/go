package main

import (
	"fmt"
	"time"
)

func fn1(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Fn1"
}

func fn2(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "Fn2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go fn1(ch1)
	go fn2(ch2)

	//r1 := <-ch1
	//fmt.Println("r1=", r1)
	//r2 := <-ch2
	//fmt.Println("r2=", r2)

	select {
	case r1 := <-ch1:
		fmt.Println("r1:", r1)
	case r2 := <-ch2:
		fmt.Println("r2:", r2)
		//default:
		//	fmt.Println("All channels are not ready")
	}

	ch := make(chan uint64)
	quit := make(chan bool)

	go func() {
		/*
		* Customer
		**/
		for i := 0; i < 100; i++ {
			num := <- ch
			fmt.Println(num)
		}
		quit <- true
	}()

	fibonacci(ch, quit)
}

func fibonacci(ch chan<- uint64, quit <-chan bool) {
	/*
	* Producer
	**/
	var x uint64 = 0
	var y uint64 = 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag: ", flag)
			return
		}
	}
}
