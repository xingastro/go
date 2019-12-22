package main

import (
	"runtime/pprof"
	"os"
	"log"
)

type user struct {
	name  string
	email string
}

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()
	f2, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal(err)
	}
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", &u2)
	if err := pprof.WriteHeapProfile(f2); err != nil {
		log.Fatal(err)
	}
	f2.Close()
}

//go:noinline
func createUserV1() user {
	u := user {
		name:  "Bill",
		email: "bill@example.com",
	}

	println("V1", &u)
	return u
}

//go:noinline
func createUserV2() *user {
	u := user {
		name: "Bill",
		email: "bill@example.com",
	}

	println("V2", &u)
	return &u
}
