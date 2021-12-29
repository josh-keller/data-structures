package sliceStack

import "fmt"

type node struct {
	val  int
	next *node
}

type Stack interface {
	Push(val int)
	Pop() int
	Length() int
	String() string
}

type stack struct {
	vals []int
}

func (s *stack) Length() int {
	return len(s.vals)
}

func (s *stack) Push(val int) {
	s.vals = append(s.vals, val)
}

func (s *stack) Pop() int {
	if len(s.vals) == 0 {
		panic("Cannot pop from empty stack!")
	}
	v := s.vals[0]

	s.vals = s.vals[1:]
	return v
}

func NewStack() Stack {
	return &stack{}
}

func (s stack) String() string {
	return fmt.Sprint("Stack(len ", len(s.vals), "): ", s.vals)
}
