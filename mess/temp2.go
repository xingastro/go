package main

import (
	"fmt"
	"time"
)

func main() {
	var result int
	go func() { 
		for { }
	}()
	go func() { 
		for { }
	}()
	go func() { 
		for { }
	}()
	go func() { 
		for { }
	}()

	time.Sleep(time.Second)
	fmt.Println("result =", result)
}
