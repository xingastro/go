package LinkedList

import (
	"fmt"
	"sync"
)

type Node struct {
	Value int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Size int
	Lock sync.RWMutex
}

func (ll *LinkedList) Append(value int) {
	ll.Lock.Lock()
	defer ll.Lock.Unlock()

	node := &Node{value, nil}
	if ll.Head == nil {
		ll.Head = node
	} else {
		curNode := ll.Head
		for curNode.Next != nil {
			curNode = curNode.Next
		}
		curNode.Next = node
	}
	ll.Size++
}

func (ll *LinkedList) Insert(index, value int) error {
	ll.Lock.Lock()
	defer ll.Lock.Unlock()

	node := &Node{value, nil}
	if index < 0 {
		return fmt.Errorf("Index out of bounds")
	}
	if ll.Head == nil {
		ll.Head = node
	} else if index == 0 {
		node.Next = ll.Head
		ll.Head = node
	} else {
		curNode := ll.Head
		for i := 1; i < index && curNode.Next != nil; i++ {
			curNode = curNode.Next
		}
		node.Next = curNode.Next
		curNode.Next = node
	}
	ll.Size++
	return nil
}

func (ll *LinkedList) RemoveAt(index int) (*Node, error) {
	ll.Lock.Lock()
	defer ll.Lock.Unlock()

	if index < 0 || index >= ll.Size{
		return nil, fmt.Errorf("Index out of bounds")
	}

	if ll.Head == nil {
		return nil, fmt.Errorf("Empty linkedlist")
	}

	if index == 0 {
		node := ll.Head
		ll.Head = ll.Head.Next
		ll.Size--
		return node, nil
	}

	curNode := ll.Head
	for i := 1; i < index; i++ {
		curNode = curNode.Next
	}
	node := curNode.Next
	curNode.Next = curNode.Next.Next
	ll.Size--
	return node, nil
}


func (ll *LinkedList) IndexOf (index int) *Node {
	ll.Lock.RLock()
	defer ll.Lock.RUnlock()

	if index < 0 || index >= ll.Size {
		return nil
	}
	if ll.Head == nil {
		return nil
	}

	curNode := ll.Head
	for i := 0; i < index; i++ {
		curNode = curNode.Next
	}
	return curNode
}

func (ll *LinkedList) IsEmpty() bool {
	ll.Lock.RLock()
	defer ll.Lock.RUnlock()

	return ll.Head == nil
}

func (ll *LinkedList) String() string {
	result := ""
	curNode := ll.Head
	for curNode != nil {
		result += fmt.Sprintf("%d ", curNode.Value)
		curNode = curNode.Next
	}
	return result
}

func (ll *LinkedList) GetSize() int {
	ll.Lock.RLock()
	defer ll.Lock.RUnlock()

	return ll.Size
}