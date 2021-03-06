package main

import (
	"fmt"
)

type item struct {
	value interface{}
	next  *item
}
type Stack struct {
	top  *product
	size int
}

func (stack *Stack) Len() int {
	return stack.size
}
func (stack *Stack) Push(value interface{}) {
	stack.top = &product{
		value: value,
		next:  stack.top,
	}
	stack.size++
}
func (stack *Stack) Pop() (value interface{}) {
	if stack.Len() > 0 {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--
		return
	}
	return nil
}
func main() {
	stack := new(Stack)

	stack.Push(1)
	stack.Push("present")
	stack.Push(3.4)

	for stack.Len() > 0 {
		fmt.Println(stack.Pop())
	}
}
