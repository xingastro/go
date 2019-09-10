package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func player(name string, court chan int) {
	defer wg.Done()
	for {
		if ball, ok := <-court; !ok {
			// If channal is closed, we won!
			fmt.Println(name, "won")
			return
		} else if rand.Intn(100)%13 == 0 {
			fmt.Println(name, "Missed")
			close(court)
			return
		} else {
			ball++
			fmt.Println(name, "hit", ball)
			court <- ball
		}
	}
}

var wg sync.WaitGroup

func main() {
	court := make(chan int)

	wg.Add(2)

	go player("Jack", court)
	go player("Bill", court)

	court <- 0
	wg.Wait()
}
