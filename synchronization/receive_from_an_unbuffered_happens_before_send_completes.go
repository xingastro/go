package main

import "fmt"

var c = make(chan int)
var a string

func f() {
	a = "hello world"
	<-c
}

func main() {
	// A receive from an unbuffered channel happens before the send on that channel completes
    go f()
	c <- 0
	fmt.Println(a)
}
