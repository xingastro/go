package BubbleSort

func BubbleSort(array []int) {
	isChanged := true
	for beDone := 0; isChanged; beDone++ {
		isChanged = false
		for i := 0; i < len(array)-1-beDone; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				isChanged = true
			}
		}
	}
}
