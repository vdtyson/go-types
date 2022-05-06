package btree

type Tree[T any] struct {
	Root *Node[T]
}

func New[T any]() *Tree[T] {
	return &Tree[T]{}
}

// FromSlice creates a new Tree of type T from a slice of items using level order traversal
func FromSlice[T any](items []T) *Tree[T] {
	if len(items) == 0 {
		return &Tree[T]{}
	}
	return &Tree[T]{Root: insertLevelOrder(items)}
}
