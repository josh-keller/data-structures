package llstack

import "fmt"

type node struct {
	val  int
	next *node
}

type Stack interface {
	Push(val int)
	Pop() int
	Peek() int
	Length() int
	IsEmpty() bool
	String() string
}

type llStack struct {
	head   *node
	length int
}

func (s *llStack) Length() int {
	return s.length
}

func (s *llStack) Push(val int) {
	n := node{val, s.head}
	s.head = &n
	s.length += 1
}

func (s *llStack) Pop() int {
	if s.length == 0 {
		panic("Stack is empty")
	}

	n := *s.head
	s.head = n.next
	s.length -= 1

	return n.val
}

func (s *llStack) IsEmpty() bool {
	return s.length == 0
}

func (s *llStack) Peek() int {
	if s.length == 0 {
		panic("Stack is empty")
	}

	return s.head.val
}

func NewStack() Stack {
	return &llStack{nil, 0}
}

func (s llStack) String() string {
	ptr := s.head
	str := fmt.Sprint("Stack(len ", s.length, "): [")

	for ptr != nil {
		str += fmt.Sprint(ptr.val)
		if ptr.next != nil {
			str += fmt.Sprint(" ")
		}

		ptr = ptr.next
	}

	return str + "]"
}
