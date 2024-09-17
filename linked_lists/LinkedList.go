package linked_lists

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	data any
	next *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func New(input any) *LinkedList {

	firstNode := &Node{
		data: input,
		next: nil,
	}
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

	newNode := &Node{
		data: value,
		next: nil,
	}
	newNode.next = l.head
	l.head = newNode
	l.length++

	return nil
}

func (l *LinkedList) Append(value any) error {

	newNode := &Node{
		data: value,
		next: nil,
	}

	l.tail.next = newNode
	l.tail = newNode
	l.length++

	return nil
}

func (l *LinkedList) Insert(value any, position int) error {
	return nil
}

func (l *LinkedList) Delete(value any) error {
	return nil
}
