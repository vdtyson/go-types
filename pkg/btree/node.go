package btree

type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{Value: val}
}
