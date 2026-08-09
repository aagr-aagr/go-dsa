package linkedlist

// List represents a singly-linked list that holds
// values of any type.
type Node[T any] struct {
	next *Node[T]
	val  T
}

type List[T any] struct {
	head *Node[T]
}

func (l *List[T]) isEmpty() bool {
	return l.head == nil
}

func (l *List[T]) Push(x T) {
	n := Node[T]{next: l.head, val: x}
	l.head = &n
}
