package main

import "fmt"

var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	c <- 0
}

func main() {
	// A send on a channel happens before the corresponding receive from that channel completes
	go f()
	<-c
	fmt.Println(a)
}
