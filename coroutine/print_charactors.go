package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Allocate 2 CPU cores to scheduler to schedule
	runtime.GOMAXPROCS(4)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()

		for count := 0; count < 1000000; count++ {
			for char := 'a'; char < 'z'+1; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 1000000; count++ {
			for char := 'A'; char < 'Z'+1; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	fmt.Println("Waiting to Finish")
	wg.Wait()
	fmt.Println("Program terminated")
}
