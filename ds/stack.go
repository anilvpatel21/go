package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func Mainstack() {
	myStack := Stack{}
	myStack.Push(100)
	myStack.Push(90)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
	myStack.Pop()
	myStack.Push(56)
	myStack.Push(56)
	fmt.Println(myStack)
}