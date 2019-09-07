package queue

import "fmt"

type Queue struct {
	Value []interface{}
}

func (queue *Queue) Push(v interface{}) {
	queue.Value = append(queue.Value, v)
}

func (queue *Queue) Pop() interface{} {
	value := queue.Value[0]
	queue.Value = queue.Value[1:]
	return value
}

func (queue *Queue) Print() {
	fmt.Println(queue.Value)
}

