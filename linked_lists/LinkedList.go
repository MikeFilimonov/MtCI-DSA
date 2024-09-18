package linked_lists

import (
	"fmt"
)

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	data any
	next *Node
}

func NewNode(input any) *Node {

	return &Node{
		data: input,
		next: nil,
	}

}

func (n *Node) Next() *Node {
	return n.next
}

func New(input any) *LinkedList {

	firstNode := NewNode(input)

	return &LinkedList{
		head:   firstNode,
		tail:   firstNode,
		length: 1,
	}

}

func (l *LinkedList) Head() *Node {
	return l.head
}

func (l *LinkedList) Tail() *Node {
	return l.tail
}

func (l *LinkedList) Prepend(value any) error {

	newNode := NewNode(value)

	newNode.next = l.head
	l.head = newNode
	l.length++

	return nil
}

func (l *LinkedList) Append(value any) error {

	newNode := NewNode(value)

	l.tail.next = newNode
	l.tail = newNode
	l.length++

	return nil
}

func (l *LinkedList) Find(position int) *Node {
	return nil
}

func (l *LinkedList) ShowList() {

	buffer := []any{}

	currentNode := l.head
	for currentNode != nil {

		buffer = append(buffer, currentNode.data)
		currentNode = currentNode.next

	}

	fmt.Println(buffer)

}

func (l *LinkedList) TraverseToIndex(position int) *Node {

	currentNode := l.head
	for i := 0; i < position; {

		if currentNode != nil {
			currentNode = currentNode.next
		}
		i++

	}

	return currentNode

}

func (l *LinkedList) Insert(value any, position int) error {

	if position < 0 {
		return fmt.Errorf(
			"failed to insert at position %d",
			position,
		)
	}

	if position >= l.length {
		l.Append(value)
		return nil
	}

	newNode := NewNode(value)
	currentNode := l.TraverseToIndex(position)
	newNode.next = currentNode.next
	currentNode.next = newNode
	l.length++

	return nil
}

func (l *LinkedList) Delete(position int) error {

	if position < 0 || position > l.length {
		return fmt.Errorf("failed to remove the item at [%d]", position)
	}

	node := l.TraverseToIndex(position)
	fmt.Println(node)

	if position == 0 {
		l.head = node.next
	} else if position == l.length-1 {
		preTailNode := l.TraverseToIndex(position - 1)
		if preTailNode != nil {
			preTailNode.next = nil
			l.tail = preTailNode
		}
	} else {

		leftNeighbour := l.TraverseToIndex(position - 1)
		leftNeighbour.next = node.next
		node = nil
	}

	node = nil
	l.length--

	return nil
}
