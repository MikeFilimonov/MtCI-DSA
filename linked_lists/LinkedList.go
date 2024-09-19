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

type DoublyLinkedList struct {
	head   *DoublyBoundedNode
	tail   *DoublyBoundedNode
	length int
}

type DoublyBoundedNode struct {
	data     any
	previous *DoublyBoundedNode
	next     *DoublyBoundedNode
}

func NewDoulblyBoundedNode(input any) *DoublyBoundedNode {

	return &DoublyBoundedNode{
		data:     input,
		previous: nil,
		next:     nil,
	}

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

func NewDoublyLinkedList(input any) *DoublyLinkedList {

	firstNode := NewDoulblyBoundedNode(input)

	return &DoublyLinkedList{
		head:   firstNode,
		tail:   firstNode,
		length: 1,
	}

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

func (d *DoublyLinkedList) Prepend(value any) error {

	newNode := NewDoulblyBoundedNode(value)

	d.head.previous = newNode
	newNode.next = d.head
	d.head = newNode
	d.length++

	return nil

}

func (l *LinkedList) Prepend(value any) error {

	newNode := NewNode(value)

	newNode.next = l.head
	l.head = newNode
	l.length++

	return nil
}

func (d *DoublyLinkedList) Append(value any) error {

	newNode := NewDoulblyBoundedNode(value)
	newNode.previous = d.tail
	d.tail.next = newNode
	d.length++

	return nil

}

func (l *LinkedList) Append(value any) error {

	newNode := NewNode(value)

	l.tail.next = newNode
	l.tail = newNode
	l.length++

	return nil
}

func (d *DoublyLinkedList) ShowList() {

	buffer := []any{}

	currentNode := d.head
	for currentNode != nil {

		buffer = append(buffer, currentNode.data)
		currentNode = currentNode.next

	}

	fmt.Println(buffer)

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

func (d *DoublyLinkedList) TraverseToIndex(position int) *DoublyBoundedNode {

	currentNode := d.head
	for i := 0; i < position; {

		if currentNode != nil {
			currentNode = currentNode.next
		}
		i++

	}

	return currentNode

}

func (d *DoublyLinkedList) Insert(value any, position int) error {

	if position < 0 {
		return fmt.Errorf(
			"failed to insert at position %d",
			position,
		)
	}

	if position >= d.length {
		d.Append(value)
		return nil
	}

	newNode := NewDoulblyBoundedNode(value)
	currentNode := d.TraverseToIndex(position - 1)
	newNode.next = currentNode.next
	newNode.previous = currentNode
	currentNode.next = newNode
	currentNode.next.previous = newNode
	d.length++

	return nil

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
	fmt.Println(currentNode)
	newNode.next = currentNode.next
	currentNode.next = newNode
	l.length++

	return nil
}

func (d *DoublyLinkedList) Delete(position int) error {

	if position < 0 || position > d.length {
		return fmt.Errorf("failed to remove the item at [%d]", position)
	}

	node := d.TraverseToIndex(position)

	if position == 0 {
		d.head = node.next
		node.next.previous = nil
	} else if position == d.length-1 {

		preTailNode := d.TraverseToIndex(position - 1)
		if preTailNode != nil {
			preTailNode.next = nil
			d.tail = preTailNode
		}

	} else {

		leftNeighbour := d.TraverseToIndex(position - 1)
		leftNeighbour.next = node.next
		node.next.previous = leftNeighbour
	}

	node = nil
	d.length--

	return nil
}

func (l *LinkedList) Delete(position int) error {

	if position < 0 || position > l.length {
		return fmt.Errorf("failed to remove the item at [%d]", position)
	}

	node := l.TraverseToIndex(position)

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
