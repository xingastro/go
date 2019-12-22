package main

import (
	"os"
	"math/rand"
	"log"
	"runtime/pprof"
	"runtime"
)

const (
	row = 10000
	col = 10000
)

func fillMatrix(m *[row][col]int) {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[i][j] = rand.Intn(10000)
		}
	}
}

func calculate(m * [row][col]int) {
	res := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			res += m[i][j]
		}
	}
}

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("Created cpu.prof failed!", err)
	}

	// Retrieve system information
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("Could not start CPU profile", err)
	}
	defer pprof.StopCPUProfile()


	// Main code
	x := [row][col]int{}
	fillMatrix(&x)
	calculate(&x)

	f1, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("Created mem.prof failed!", err)
	}

	runtime.GC()
	if err := pprof.WriteHeapProfile(f1); err != nil {
		log.Fatal("Could not write memory profile:", err)
	}
	f1.Close()

	f2, err := os.Create("goroutine.prof")
	if err != nil {
		log.Fatal("Created goroutine.prof failed!", err)
	}

	if gProf := pprof.Lookup("goroutine"); gProf == nil {
		log.Fatal("Could not write goroutine profile:", err)
	} else {
		gProf.WriteTo(f2, 0)
	}
	f2.Close()
}
