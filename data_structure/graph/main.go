package main

import (
	queue2 "first/data_structure/queue"
	stack2 "first/data_structure/stack"
	"fmt"
)

type Graph struct {
	nodes map[byte][]byte
}

func dfs(start byte, graph *Graph) {
	stack := stack2.Stack{}
	seen := make(map[byte]bool)
	stack.Push(start)
	seen[start] = true

	for len(stack.Value) != 0 {
		vertex := stack.Pop().(byte)
		fmt.Printf("%c  ", vertex)
		nodes := graph.nodes[vertex]
		for _, node := range nodes {
			if _, ok := seen[node]; !ok {
				// v not in seen
				stack.Push(node)
				seen[node] = true
			}
		}
	}
}

func bfs(start byte, graph *Graph) {
	//
	queue := queue2.Queue{}
	seen := make(map[byte]bool)
	queue.Push(start)
	seen[start] = true
	for len(queue.Value) != 0 {
		vertex := queue.Pop().(byte)
		nodes := graph.nodes[vertex]
		fmt.Printf("%c  ", vertex)
		for _, node := range nodes {
			if !seen[node] {
				// node not in seen
				seen[node] = true
				queue.Push(node)
			}
		}
	}

//	queue := queue2.Queue{}
//	seen := make(map[byte]bool)
//	queue.Push(start)
//	seen[start] = true
//
//	for len(queue.Value) != 0 {
//		vertex := queue.Pop().(byte)
//		fmt.Printf("%c  ", vertex)
//		nodes := graph.nodes[vertex]
//		for _, node := range nodes {
//			if !seen[node] {
//				queue.Push(node)
//				seen[node] = true
//			}
//		}
//	}
}

func main() {

	graph := Graph{
		nodes: map[byte][]byte{
			'A': {'B', 'C'},
			'B': {'A', 'D'},
			'C': {'A', 'E', 'G'},
			'D': {'B', 'E'},
			'E': {'C', 'D', 'F'},
			'F': {'E'},
			'G': {'C'},
		},
	}

	fmt.Println("\n********dfs*******")
	dfs('A', &graph)
	fmt.Println("\n********bfs*******")
	bfs('A', &graph)

}
