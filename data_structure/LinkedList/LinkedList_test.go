package LinkedList

import (
	"testing"
)

var ll LinkedList



func TestLinkedList_Append(t *testing.T) {
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)
	if ll.String() != "1 2 3 4 " {
		t.Fail()
	}
}

func TestLinkedList_Insert(t *testing.T) {
	ll.Insert(0, 0)
	if ll.String() != "0 " {
		t.Fail()
	}
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Insert(2, 0)
	if ll.String() != "0 1 0 2 3 " {
		t.Fail()
	}

	ll.Insert(100, 100)
	if ll.String() != "0 1 0 2 3 100 " {
		t.Fail()
	}
}

func TestLinkedList_IndexOf(t *testing.T) {
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)

	if ll.IndexOf(0).Value != 1 {
		t.Fail()
	}
}

func TestLinkedList_GetSize(t *testing.T) {
	if ll.GetSize() != 0 {
		t.Fail()
	}

	ll.Append(1)
	ll.Append(1)
	ll.Append(1)
	ll.Append(1)
	if ll.GetSize() != 4 {
		t.Fail()
	}

	ll.RemoveAt(0)
	if ll.GetSize() != 3 {
		t.Fail()
	}
}