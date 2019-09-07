package BinarySearchTree

import (
	"testing"
)

var bst BinarySearchTree


func fillTree(bst *BinarySearchTree) {
	bst.Insert(8)
	bst.Insert(9)
	bst.Insert(4)
	bst.Insert(2)
	bst.Insert(0)
	bst.Insert(10)
	bst.Insert(1)
	bst.Insert(7)
	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(1)
}

func TestBinarySearchTree_Insert(t *testing.T) {
	fillTree(&bst)
	bst.String()

	bst.Insert(11)
	bst.String()
}


// isSameSlice returns true if the 2 slices are identical
func isSameSlice(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestBinarySearchTree_Max(t *testing.T) {
	if bst.Max().Value != 11 {
		t.Errorf("max should be 11")
	}
}

func TestBinarySearchTree_Min(t *testing.T) {
	if bst.Min().Value != 0 {
		t.Errorf("min should be 0")
	}
}

func TestBinarySearchTree_Search(t *testing.T) {
	for _, v := range []int{1, 11, 8} {
		_, result := bst.Search(v)
		if !result {
			t.Errorf("search not working")
		}
	}
}

func TestBinarySearchTree_Remove(t *testing.T) {
	bst.Remove(1)
	if bst.Min().Value != 2 {
		t.Errorf("min should be 2")
	}
}