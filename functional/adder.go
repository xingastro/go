package main

import "fmt"

func adder() func (v int) int {
	sum := 0
	return func (v int) int {
		sum += v
		return sum
	}
}

func main() {
	add := adder()
	for i := 0; i <= 10; i++ {
		fmt.Println(add(i))
	}
}
