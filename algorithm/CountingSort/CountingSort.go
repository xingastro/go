package countingsort


// If the given numbers are in a certain range, 
// we can apply countingsort in this case.

// Assume that the numbers start from 0,
// and end at 10
func CountingSort(nums []int) {
	var newArr [10]int

	// Count the times of occrurences of each number
	for _, v := range nums {
		newArr[v]++
	}

	// re-write the value to the original nums
	indexForOriginalNums := 0
	for i := 0; i < 10; i++ {
		if newArr[i] != 0 {
			for j := 0; j < newArr[i]; j++ {
				nums[j+indexForOriginalNums] = i
			}
			indexForOriginalNums += newArr[i]
		}
	}
}
