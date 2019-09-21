package main

import "fmt"


func modifySlice(nums []int) {
	fmt.Println(len(nums), cap(nums))
	nums = append(nums, 10)
}


func main() {
	s1 := []int{1, 2, 3} // len == 3, cap == 3
	modifySlice(s1)
	fmt.Println(s1)

	s2 := make([]int, 0, 3)
	modifySlice(s2)
	fmt.Println(s2)
}
