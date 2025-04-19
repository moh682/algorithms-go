package linkedlist

import "log"

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
	size int
}

func New[T comparable]() LinkedList[T] {
	return LinkedList[T]{
		head: nil,
	}
}

/*
[0]
add 1
[1]

add 2
[ 2 -> 1]

add 3
[3 -> 2 -> 1]
*/
func (l *LinkedList[T]) Add(value T) {
	node := &Node[T]{Value: value, Next: nil, Prev: nil}

	if l.head == nil {
		l.head = node

	} else {
		l.head.Prev = node
		node.Next = l.head
		node.Prev = nil
		l.head = node
	}

	l.size++
}

func (l *LinkedList[T]) RemoveLast() {
	if l.head == nil {
		return
	}

	curr := l.head
	for curr.Next != nil {

		if curr.Next.Next == nil {
			curr.Next = nil
			l.size--
			return
		}

		curr = curr.Next

	}

}

func (l *LinkedList[T]) RemoveFirst() {
	if l.head == nil {
		return
	}

	l.head = l.head.Next
	l.size--

}

func (l *LinkedList[T]) Print() {
	node := l.head
	for node != nil {
		log.Println(node.Value)
		node = node.Next
	}
}

func (l *LinkedList[T]) Contains(value T) bool {

	node := l.head
	for node != nil {
		if node.Value == value {
			return true
		}
		node = node.Next
	}
	return false

}

// Get returns the value at the given index
func (l *LinkedList[T]) Len() int {
	return l.size
}
