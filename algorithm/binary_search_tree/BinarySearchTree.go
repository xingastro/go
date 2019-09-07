package BinarySearchTree

import (
	"fmt"
	"sync"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
	Lock sync.RWMutex
}

func (bst *BinarySearchTree) Insert(value int) {
	bst.Lock.Lock()
	defer bst.Lock.Unlock()

	n := &Node{value, nil, nil}
	if bst.Root == nil {
		bst.Root = n
	} else {
		insertNode(bst.Root, n)
	}
}

// internal function to find the correct place for a node
// in a tree
func insertNode(subTreeRoot, node *Node) {
	if subTreeRoot.Value < node.Value {
		if subTreeRoot.Left == nil {
			subTreeRoot.Left = node
		} else {
			insertNode(subTreeRoot.Left, node)
		}
	} else {
		if subTreeRoot.Right == nil {
			subTreeRoot.Right = node
		} else {
			insertNode(subTreeRoot.Right, node)
		}
	}
}

// InOrderTraverse visits all nodes with in-order traversing
func (bst *BinarySearchTree) InOrderTraverse() {
	bst.Lock.Lock()
	defer bst.Lock.Unlock()

	inOrderTraverse(bst.Root)
}

// internal recursive function to traverse in order
func inOrderTraverse(n *Node) {
	if n == nil {
		return
	}
	inOrderTraverse(n.Left)
	fmt.Printf("%d  ", n.Value)
	inOrderTraverse(n.Right)
}

func (bst *BinarySearchTree) PreOrderTraverse() {
	bst.Lock.Lock()
	defer bst.Lock.Unlock()

	preOrderTraverse(bst.Root)
}

func preOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d  ", node.Value)
	preOrderTraverse(node.Left)
	preOrderTraverse(node.Right)
}

func (bst *BinarySearchTree) PostOrderTraverse() {
	bst.Lock.Lock()
	defer bst.Lock.Unlock()

	postOrderTraverse(bst.Root)
}

func postOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	postOrderTraverse(node.Left)
	postOrderTraverse(node.Right)
	fmt.Printf("%d  ", node.Value)
}

func (bst *BinarySearchTree) Min() *Node {
	bst.Lock.Lock()
	defer bst.Lock.Unlock()
	node := bst.Root
	if node == nil {
		return nil
	}
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (bst *BinarySearchTree) Max() *Node {
	bst.Lock.Lock()
	defer bst.Lock.Unlock()

	node := bst.Root
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (bst *BinarySearchTree) Search(value int) (*Node, bool) {
	bst.Lock.Lock()
	defer bst.Lock.Unlock()
	node := search(bst.Root, value)
	if node != nil {
		return node, true
	} else {
		return nil, false
	}
}

func search(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if node.Value == value {
		return node
	} else if node.Value > value {
		return search(node.Left, value)
	} else {
		// node.Value < value
		return search(node.Right, value)
	}
}

func (bst *BinarySearchTree) Remove(value int) {
	bst.Lock.Lock()
	defer bst.Lock.Unlock()

	var parent *Node
	curNode := bst.Root

	for curNode != nil {
		switch {
		case curNode.Value == value: // curNode is going to be deleted
			if curNode.Left != nil {
				rightSubTreeRoot := curNode.Right
				curNode.Value = curNode.Left.Value
				curNode.Right = curNode.Left.Right
				curNode.Left = curNode.Left.Left

				if rightSubTreeRoot != nil {
					insertNodes(curNode, rightSubTreeRoot)
				}

			} else if curNode.Right != nil {
				curNode.Value = curNode.Right.Value
				curNode.Left = curNode.Right.Left
				curNode.Right = curNode.Right.Right
			} else {
				if parent == nil { // root node is ganna be deleted
					bst.Root = nil
					return
				}
				if parent.Left != nil && parent.Left.Value == value {
					// curNode is left node of the parent
					parent.Left = nil
				} else {
					parent.Right = nil
				}
				return
			}
		case curNode.Value > value:
			parent = curNode
			curNode = curNode.Left
		case curNode.Value < value:
			parent = curNode
			curNode = curNode.Right
		}
	}
}

func insertNodes(treeRoot, nodes *Node) {
	if nodes == nil {
		return
	}
	insertNode(treeRoot, &Node{nodes.Value, nil, nil})
	insertNodes(treeRoot, nodes.Left)
	insertNodes(treeRoot, nodes.Right)
}

func (bst *BinarySearchTree) String() {
	bst.Lock.Lock()
	defer bst.Lock.Unlock()
	fmt.Println("------------------------")
	stringify(bst.Root, 0)
	fmt.Println("------------------------")
}

func stringify(node *Node, level int) {
	if node != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "     "
		}
		format += "---[ "
		level++
		stringify(node.Left, level)
		fmt.Printf(format+"%d\n", node.Value)
		stringify(node.Right, level)
	}
}