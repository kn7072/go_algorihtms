package main

//import "fmt"


type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

type Node[T Ordered] struct {
	Value               T
	Parent, Left, Right *Node[T]
}

func NewTree[T Ordered](value T) *Node[T] {
	return &Node[T]{Value: value}
}

func insert[T Ordered](root, node **Node[T]) {
	if (*node).Value < (*root).Value {
		if (*root).Left == nil {
			(*root).Left = *node
			(*node).Parent = *root
		} else {
			insert(&(*root).Left, node)
		}
	} else {
		if (*root).Right == nil {
			(*root).Right = *node
			(*node).Parent = *root
		} else {
			insert(&(*root).Right, node)
		}
	}
	
}

func find[T Ordered](root **Node[T], value T) *Node[T] {
	if (*root).Value == value {
		return *root
	}

	if (*root).Left != nil {
		if value < (*root).Value {
			return find(&(*root).Left, value)
		}
	}

	if (*root).Right != nil {
		if value >= (*root).Value {
			return find(&(*root).Right, value)
		}
	}

	return nil
}

func zig[T Ordered](node *Node[T]) {
	parentLink := node.Parent
	
	parentValue := *parentLink
	parentValue.Left = node.Right
	node.Right = &parentValue
	parentValue.Parent = node

	if parentLink.Parent == nil {
		node.Parent = nil // is root
	} else {
		node.Parent = parentLink.Parent
	}
	
	*parentLink = *node
}

func splay[T Ordered](node **Node[T]) {

}