package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Stack[T constraints.Ordered] struct {
	items []T
}

func NewStack[T constraints.Ordered]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func main() {
	s := NewStack[int]()
	s.Push(5)
	s.Push(1)
	s.Push(9)

	top, ok := s.Pop()
	fmt.Println(top, ok)

	fmt.Println(s.Len())
}
