package stacks

import "fmt"

type Stack struct {
	top    *Node
	bottom *Node
	length int
}

type Node struct {
	data any
	next *Node
}

type ArrayBasedStack struct {
	data []interface{}
}

func NewArrayBasedStack() *ArrayBasedStack {
	return &ArrayBasedStack{
		data: []any{},
	}
}

func (a *ArrayBasedStack) Peek() {

	if len(a.data) > 0 {
		fmt.Println(a.data[len(a.data)-1])
	} else {
		fmt.Println(nil)
	}

}

func (a *ArrayBasedStack) Pop() {
	if len(a.data) < 1 {
		return
	}

	a.data = a.data[:len(a.data)-1]

}

func (a *ArrayBasedStack) Push(value any) {
	a.data = append(a.data, value)
}

func NewStack() *Stack {
	return &Stack{
		top:    nil,
		bottom: nil,
		length: 0,
	}
}

func NewNode(input any) *Node {
	return &Node{
		data: input,
		next: nil,
	}
}

func (s *Stack) Peek() *Node {

	if s.length == 0 {
		return nil
	} else {
		return s.top
	}

}

func (s *Stack) Pop() {

	if s.length == 0 {
		return
	}

	if s.length == 1 {
		s.bottom = nil
	}

	holdingPointer := s.top
	s.top = holdingPointer.next
	s.length--

}

func (s *Stack) Push(value any) {

	newNode := NewNode(value)

	if s.top == nil {
		s.bottom = newNode
		s.top = newNode
	} else {

		exTop := s.top
		s.top = newNode
		s.top.next = exTop
	}
	s.length++
}

func (s *Stack) IsEmpty() bool {
	return s.length == 0
}
