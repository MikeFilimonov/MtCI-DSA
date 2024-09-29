package queues

import "fmt"

type Node struct {
	data any
	next *Node
}

type QueueOnTwoStacks struct {
	first []any
	last  []any
}

func (qots *QueueOnTwoStacks) Enqueue(value any) {

	length := len(qots.first)

	for i := 0; i < length; i++ {
		qots.last = append(qots.last, qots.first[length-1])
		qots.first = qots.first[:length-1]
	}

	qots.last = append(qots.last, value)

}

func (qots *QueueOnTwoStacks) Dequeue() {

	length := len(qots.last)
	for i := 0; i < length; i++ {
		qots.first = append(qots.first, qots.last[length-1])
		qots.last = qots.last[:length-1]
	}

	qots.first = qots.first[:length-1]

}

func (qots *QueueOnTwoStacks) Peek() {

	if len(qots.first) > 0 {
		fmt.Println(qots.first[len(qots.first)-1])
		return
	}

	fmt.Println(qots.last[0])

}

func (qots *QueueOnTwoStacks) IsEmpty() bool {

	if len(qots.first) > 0 {
		return len(qots.first) == 0
	}

	return len(qots.last) == 0

}

type QueueOnStack struct {
	stack []int
}

func Constructor() QueueOnStack {
	return QueueOnStack{
		stack: []int{},
	}
}

func (qos *QueueOnStack) Push(x int) {
	qos.stack = append(qos.stack, x)
}

func (qos *QueueOnStack) Pop() int {

	firstElement := qos.stack[0]
	qos.stack = qos.stack[1:]
	return firstElement

}

func (qos *QueueOnStack) Peek() int {
	return qos.stack[0]
}

func (qos *QueueOnStack) IsEmpty() bool {
	return len(qos.stack) == 0
}

type Queue struct {
	first  *Node
	last   *Node
	length int
}

func NewQueue() *Queue {
	return &Queue{
		first:  nil,
		last:   nil,
		length: 0,
	}
}

func NewNode(value any) *Node {
	return &Node{
		data: value,
		next: nil,
	}
}

func (q *Queue) Enqueue(data any) {

	newNode := NewNode(data)
	if q.length == 0 {
		q.first = newNode
		q.last = newNode
	} else {
		q.last.next = newNode
		q.last = newNode
	}
	q.length++

}

func (q *Queue) Dequeue() {

	if q.length == 0 {
		return
	}

	if q.length == 1 {
		q.last = nil
	}

	q.first = q.first.next
	q.length--

}

func (q *Queue) Peek() *Node {
	return q.first
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}
