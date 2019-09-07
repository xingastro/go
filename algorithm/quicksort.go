package main

import "fmt"

func quicksort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	pivot := slice[len(slice)-1]
	slice = slice[:len(slice)-1]

	smaller, bigger := []int{}, []int{}
	for _, v := range slice {
		if v <= pivot {
			smaller = append(smaller, v)
		} else {
			bigger = append(bigger, v)
		}
	}
	return append(append(quicksort(smaller), pivot), quicksort(bigger)...)
}

func main() {
	slice := []int{9, 4, 2, 5, 1, 0, 4}
	slice = quicksort(slice)
	fmt.Println(slice)
}
