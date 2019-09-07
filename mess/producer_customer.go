package main

import (
	"fmt"
	"math/rand"
)

func producer(ch chan<- int) {
	for {
		ch <- rand.Int()
	}
}

func customer(ch <-chan int) {
	for {
		fmt.Println(<- ch)
	}
}

func main() {
	ch := make(chan int, 10)
	go producer(ch)
	go customer(ch)

	fmt.Println("len:", len(ch))
	fmt.Println("cap:", cap(ch))

}
