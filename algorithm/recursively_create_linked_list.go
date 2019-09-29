package main

import "fmt"


type ListNode struct {
	Val int
	Next *ListNode
}

func create(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	firstNode := &ListNode{nums[0], nil}
	firstNode.Next = create(nums[1:])
	return firstNode
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	head := create([]int{})
	printLinkedList(head)

	head = create([]int{1})
	printLinkedList(head)

	head = create([]int{1, 2, 3, 4, 5, 6})
	printLinkedList(head)
}
