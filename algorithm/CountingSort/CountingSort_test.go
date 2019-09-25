package countingsort

import (
	"testing"
)

func TestCountingSort(t *testing.T) {
	nums := []int{5, 1, 2, 9, 2, 1, 0, 2, 5, 3, 3, 6, 8}
	CountingSort(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			t.Fail()
		}
	}
}
