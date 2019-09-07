package main

import (
	"fmt"
	"time"
)
func worker(id int, c chan int) {
	for i := range c{
		fmt.Printf("Worker %d received %d\n",
			id, i)
	}
}
func createWorker(id int) chan<- int {
	c := make(chan int)

	go worker(id, c)

	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(-1, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	close(c)
}


func main() {
	c := make(chan int)
	c <- 10
	fmt.Println(<- c)
}
