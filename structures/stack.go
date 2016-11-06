package main

import "fmt"

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func (s *Stack) Len() int {
	return s.size
}

func main() {
	stack := new(Stack)
	stack.Push("Things")
	stack.Push("And")
	stack.Push("Stuffs")

	for stack.Len() > 0 {
		fmt.Printf("%s ", stack.Pop().(string))
	}

	fmt.Println()
}
