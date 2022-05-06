package btree

type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{Value: val}
}

func insertLevelOrder[T any](items []T) *Node[T] {
	return insertLevelOrderHelper(items, 0)
}

func insertLevelOrderHelper[T any](items []T, i int) *Node[T] {
	if i >= len(items) {
		return nil
	}
	root := NewNode(items[i])

	root.Left = insertLevelOrderHelper(items, 2*i+1)
	root.Right = insertLevelOrderHelper(items, 2*i+2)

	return root
}
