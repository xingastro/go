package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			out <- i
			i++
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		}
	}()
	return out
}

func main() {
	c1, c2 := generator(), generator()
	for {
		select {  // select挑选两个channel发送数据更快的那个
		case n := <-c1:
			fmt.Println("Received from c1: ", n)
		case n := <-c2:
			fmt.Println("Received from c2: ", n)
		//default:
		//	fmt.Println("No value received")
		//
		}
	}
}
