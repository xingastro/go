package main

func adder() func (int, int) int {
	return func(a, b int) int {
		return a + b
	}
}

func main() {
	add := adder()
	var start int
	for i := 0; i < 100; i++ {
		start = add(start, i)
	}
	println(start)
}
