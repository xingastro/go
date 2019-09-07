package stack

type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value interface{}
		prev *node
	}
)

// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}

// View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}
	top := this.top
	this.top = this.top.prev
	this.length--
	return top.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	this.top = &node{value, this.top}
	this.length++
}