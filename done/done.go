package main

import (
	"fmt"
	"sync"
	"time"
)

func doWork(id int, w worker) {
	//slice := make([]int, 10, 10)
	for n := range w.in {
		fmt.Printf("%d Worker %d received %c\n",
			doWork, id, n)
		//slice = append(slice, n)
		w.done()
	}
	//fmt.Println(id, slice)
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in:   make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)

	return w
}

func chanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}


	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanDemo()
	time.Sleep(time.Second)
}
