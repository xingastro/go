package main

import (
	"fmt"
	"os"
)

type Person struct {
	name     string
	age      int
	relation map[string][]string
}

func (p Person) Write(w []byte) {
	fmt.Fprintln(os.Stdout, p)
}

func main() {
	p := Person{"Bender", 20, map[string][]string{
		"Astro": {"Bender", "John", "Bill"},
		"John":  {"Adam", "Bill"},
		"Adam":  {"John", "Bill"},
	}}

	fmt.Println(p)
}
