package MergeSort

func Merge(arr1, arr2 []int) []int {
	final := []int{}
	for i, j := 0, 0; i < len(arr1) && j < len(arr2); {
		if arr1[i] > arr2[j] {
			final = append(final, arr1[i])
			i++
		} else if arr1[i] < arr2[j] {
			final = append(final, arr2[j])
			j++
		} else {
			final = append(final, arr1[i], arr2[j])
			i++
			j++
		}
	}

	return final
}

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		// When D&C to 1 element, just return it.
		return array
	}

	mid := len(array) / 2
	left := array[:mid]
	right := array[mid:]

	left = MergeSort(left)
	right = MergeSort(right)

	return Merge(left, right)
}
