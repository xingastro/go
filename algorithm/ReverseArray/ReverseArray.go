package ReverseArray

func ReverseArray(a []int) {

	for i := 0; i < len(a) / 2; i++ {
		a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}

}
