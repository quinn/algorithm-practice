package linkedlist

type Node[T any] struct {
	value T
	next  *Node[T]
}

func New[T any](rootVal T) *Node[T] {
	return &Node[T]{
		value: rootVal,
	}
}
