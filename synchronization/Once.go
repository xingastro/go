package main

import (
	"fmt"
	"sync"
)

var a string
var once sync.Once

func setup() {
	a = "Hello, world"
}

func doprint() {
	once.Do(setup)
	fmt.Println(a)
}

func twoprint() {
	go doprint()
	go doprint()
}

func main() {
	twoprint()
}
