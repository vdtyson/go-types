package btree

type Tree[T any] struct {
	Root *Node[T]
}

func New[T any]() *Tree[T] {
	return &Tree[T]{}
}
