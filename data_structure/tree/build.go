package main

import (
	"fmt"
)

type ElementType int

type Node struct {
	Value  ElementType
	Parent *Node
	Left   *Node
	Right  *Node
}

func PreOrderTraverse(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%d --> ", node.Value)
	PreOrderTraverse(node.Left)
	PreOrderTraverse(node.Right)
}

func InOrderTraverse(node *Node) {
	if node == nil {
		return
	}

	InOrderTraverse(node.Left)
	fmt.Printf("%d --> ", node.Value)
	InOrderTraverse(node.Right)
}

func PostOrderTraverse(node *Node) {
	if node == nil {
		return
	}

	PostOrderTraverse(node.Left)
	PostOrderTraverse(node.Right)
	fmt.Printf("%d --> ", node.Value)
}

func NewNode(v ElementType) *Node {
	return &Node{Value: v, Parent: nil, Left: nil, Right: nil}
}

/*
*节点比较
* n>m:1  n<m:-1  n=m:0
 */
func (n *Node) Compare(m *Node) int {
	if n.Value < m.Value {
		return -1
	} else if n.Value > m.Value {
		return 1
	} else {
		return 0
	}
}

type BinaryTree interface {
	Insert(i ElementType)
	Delete(i ElementType) bool
	Search(i ElementType) *Node
	PreOrderTraverse()
	InOrderTraverse()
	PostOrderTraverse()
}

// Tree
type Tree struct {
	Root *Node
	Size int
}

// Create Tree
func NewTree(n *Node) *Tree {
	if n == nil {
		return &Tree{}
	}
	return &Tree{Root: n, Size: 1}
}

// Insert, 相同的节点值，放到右子树
func (t *Tree) Insert(i ElementType) {
	n := NewNode(i)
	if t.Root == nil {
		t.Root = n
		t.Size++
		return
	}

	curNode := t.Root

	for {
		if n.Compare(curNode) == -1 { //  n的值小于Parent，到左子节点
			if curNode.Left == nil {
				curNode.Left = n
				n.Parent = curNode
				break
			} else {
				curNode = curNode.Left
			}
		} else { // n 的值大于等于Parent,到右子节点
			if curNode.Right == nil {
				curNode.Right = n
				n.Parent = curNode
				break
			} else {
				curNode = curNode.Right
			}
		}
	}

	t.Size++
}

//Search
func (t *Tree) Search(i ElementType) *Node {
	curNode := t.Root
	n := NewNode(i)

	for curNode != nil {
		switch curNode.Compare(n) {
		case -1: //curNode is less than n
			curNode = curNode.Right
		case 1: //curNode is greater than n
			curNode = curNode.Left
		case 0:
			return curNode
		default:
			return nil
		}
	}
	return nil
}

//Delete
func (t *Tree) Delete(i ElementType) bool {
	var parent *Node

	curNode := t.Root
	n := NewNode(i)

	for curNode != nil {
		switch n.Compare(curNode) {
		case -1: //n is less than curNode
			parent = curNode
			curNode = curNode.Left
		case 1: //n is greater than curNode
			parent = curNode
			curNode = curNode.Right
		case 0:
			if curNode.Left != nil {
				right := curNode.Right // right 是当前节点的右子树
				curNode.Value = curNode.Left.Value
				curNode.Right = curNode.Left.Right
				curNode.Left = curNode.Left.Left

				if right != nil {
					subTree := &Tree{Root: curNode}
					IterOnTree(right, func(n *Node) {
						subTree.Insert(n.Value)
					})
				}

				t.Size--
				return true
			}

			if curNode.Right != nil {
				curNode.Value = curNode.Right.Value
				curNode.Left = curNode.Right.Left
				curNode.Right = curNode.Right.Right

				t.Size--
				return true
			}

			if parent == nil {
				// 这棵树的Root节点是要被删除的对象
				t.Root = nil
				t.Size--
				return true
			}

			// 当前节点是要被删除的对象，且当前节点无左节点也无右节点
			if parent.Left == n {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
			t.Size--
			return true
		}
	}

	return false
}

func IterOnTree(node *Node, f func(n *Node)) bool {
	if node == nil {
		return true
	}

	f(node)

	return IterOnTree(node.Left, f) && IterOnTree(node.Right, f)
}

func (tree *Tree) PreOrderTraverse()  { PreOrderTraverse(tree.Root) }
func (tree *Tree) InOrderTraverse()   { InOrderTraverse(tree.Root) }
func (tree *Tree) PostOrderTraverse() { PostOrderTraverse(tree.Root) }

func howMuchLayers(bt BinaryTree) {
	switch v := bt.(type) {
	case *Tree:
		fmt.Println("Tree pointer receiver: ", v.Size)
	}
}

func main() {
	root := NewTree(NewNode(10))
	nodeValue := []int{5, 3, 7, 11, 2, 3}
	for _, v := range nodeValue {
		root.Insert(ElementType(v))
	}

	howMuchLayers(root)

	fmt.Println("-----------PreOrderTraverse---------")
	root.PreOrderTraverse()
	fmt.Println("\n-----------InOrderTraverse----------")
	root.InOrderTraverse()
	fmt.Println("\n-----------PostOrderTraverse--------")
	root.PostOrderTraverse()
}
