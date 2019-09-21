package main

import "fmt"

func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	pivot := nums[0]
	l, r := 0, len(nums)-1
	for i := 1; i <= r; {
		if nums[i] > pivot {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		} else {
			nums[i], nums[l] = nums[l], nums[i]
			i++
			l++
		}
	}
	quickSort(nums[:l])
	quickSort(nums[l+1:])
}


func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	quickSort(nums)
	n := len(nums)
	ans := [][]int{}
	for i := 0; i < n-2; i++ {
		if nums[i] + nums[n-1] + nums[n-2] < 0 {
			continue
		} else if nums[i] + nums[i+1] + nums[i+2] > 0 {
			break
		} else if i > 0 && nums[i] == nums[i-1] {
			continue
		} else {
			l, r := i+1, n-1
			for l < r {
				temp := nums[i] + nums[l] + nums[r]
				if temp == 0 {
					ans = append(ans, []int{nums[i], nums[l], nums[r]})
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					r--
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					l++
				} else if temp > 0 {
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					r--
				} else {
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					l++
				}
			}
		}
	}
	return ans
}

func main() {
	nums := []int{-4, -1, -1, 0, 1, 2}
	fmt.Println(threeSum(nums))
}