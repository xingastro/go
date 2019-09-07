package InsertionSort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestInsertionSort(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	array1 := make([]int, random.Intn(100-10)+10) // make([]int, 10~100)
	for i := range array1 {
		array1[i] = random.Intn(100)
	}

	array2 := make(sort.IntSlice, len(array1))
	copy(array2, array1)
	InsertionSort(array1)
	array2.Sort()
	for i := 0; i < len(array1); i++ {
		if array1[i] != array2[i] {
			t.Fail()
		}
	}
}
