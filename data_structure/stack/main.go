package stack

import "fmt"

type Stack struct {
	Value []interface{}
}

func (stack *Stack) Push(v interface{}) {
	stack.Value = append(stack.Value, v)
}

func (stack *Stack) Pop() interface{} {
	value := stack.Value[len(stack.Value)-1]
	stack.Value = stack.Value[:len(stack.Value)-1]
	return value
}

func (stack *Stack) Print() {
	fmt.Println(stack.Value)
}


