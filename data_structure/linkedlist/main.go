package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(index int, v interface{}) {
	node := &Node{v, nil}

	// In this case, list is an empty LinkedList,
	// so just insert the node into list.head
	if list.head == nil {
		list.head = node
		return
	}

	// insert to index 0
	if index == 0 {
		node.next = list.head.next
		list.head = node
		return
	}

	curNode := list.head

	for curIndex := 1; curNode.next != nil; curIndex++ {
		if curIndex == index {
			node.next = curNode.next
			curNode.next = node
			return
		}

		curNode = curNode.next
	}

	// index is bigger than the index of the last element of the list
	// insert it to the tail of the list
	curNode.next = node
}

func (list *LinkedList) Append(v interface{}) {
	node := &Node{v, nil}

	if list.head == nil {
		list.head = node
	}

	curNode := list.head
	for curNode.next != nil {
		curNode = curNode.next
	}
	curNode.next = node
}

func (list *LinkedList) Delete(v interface{}) bool {
	if list.head == nil {
		return false
	}
	if list.head.data == v {
		list.head = list.head.next
		return true
	}

	curNode := list.head
	for curNode.next != nil {
		if curNode.next.data == v {
			curNode.next = curNode.next.next
			return true
		}
		curNode = curNode.next
	}

	return false
}

func (list *LinkedList) Find(v interface{}) (int, *Node) {
	curNode := list.head
	for i := 0; curNode != nil; i++ {
		if curNode.data == v {
			return i, curNode
		}
		curNode = curNode.next
	}
	// Does not found!
	return 0, nil
}

func (list *LinkedList) GetLength() int {
	curNode := list.head
	length := 0
	for curNode != nil {
		length++
		curNode = curNode.next
	}
	return length
}

func (list *LinkedList) GetItems() []interface{} {
	var items []interface{}
	curNode := list.head
	for curNode != nil {
		items = append(items, curNode.data)
		curNode = curNode.next
	}
	return items
}

func (l *LinkedList) Print() {
	curNode := l.head
	for curNode != nil {
		fmt.Printf("%v --> ", curNode.data)
		curNode = curNode.next
	}
	fmt.Println("nil\n")
}

func main() {
	l := LinkedList{}
	l.Print()
	fmt.Println(l.GetItems())
	fmt.Println(l.GetLength())

	l.Insert(0, "hello")
	l.Print()

	l.Append("world")
	l.Print()
	l.Delete("world")
	l.Print()

	l.Insert(10, "golang")
	l.Insert(1, "python")
	l.Print()
	fmt.Println(l.GetItems(), l.GetLength())
	fmt.Println(l.Find("golang"))
}
