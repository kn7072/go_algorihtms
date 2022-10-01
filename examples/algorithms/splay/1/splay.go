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

func posiotionNode[T Ordered](node *Node[T]) string {
	if node.Parent == nil {
		return "root"
	} else if (node.Parent.Left != nil && node.Parent.Left == node) {
		return "left"
	} else {
		return "right"
	}
}



func zig[T Ordered](node *Node[T]) (isRoot bool){
	

	// parentPtr := node.Parent
	parentValue := *node.Parent

	switch posiotionNode(node){
	case "left":
		if node.Right == nil {
			parentValue.Left = node.Right
		} else {
			parentValue.Left = node.Right
			parentValue.Left.Parent = &parentValue
		}

		node.Right = &parentValue
		parentValue.Parent = node
		
		if parentValue.Right != nil {
			parentValue.Right.Parent = &parentValue
		}
		switch posiotionNode(node.Parent){
		case "left":
			// parent is left child of grandparent
			node.Parent.Parent.Left = node
			node.Parent = node.Parent.Parent
		case "right":
			// parent is right child of grandparent
			node.Parent.Parent.Right = node
			node.Parent = node.Parent.Parent
		case "root":
			// parent is root
			*(node.Parent) = *node
			node.Parent.Parent = nil
			isRoot = true
		}

	case "right":
		// if node.Left == nil {
		// 	parentValue.Right = node.Left
		// } else {
		// 	parentValue.Right = node.Left
		// 	parentValue.Right.Parent = &parentValue
		// 	node.Left = &parentValue
		// 	parentValue.Parent = node
		// }

		if node.Left == nil {
			parentValue.Right = node.Left
		} else {
			parentValue.Right = node.Left
			parentValue.Right.Parent = &parentValue
		}

		node.Left = &parentValue
		parentValue.Parent = node
		
		if parentValue.Left != nil {
			parentValue.Left.Parent = &parentValue
		}
		switch posiotionNode(node.Parent){
		case "left":
			// parent is left child of grandparent
			node.Parent.Parent.Left = node
			node.Parent = node.Parent.Parent
		case "right":
			// parent is right child of grandparent
			node.Parent.Parent.Right = node
			node.Parent = node.Parent.Parent
		case "root":
			// parent is root
			*(node.Parent) = *node
			node.Parent.Parent = nil
			isRoot = true
		}
	case "root":
		return true
	}

	return isRoot
}

func zigOLD[T Ordered](node *Node[T]) (isRoot bool) {
	parentLink := node.Parent
	parentValue := *parentLink
	
	if parentLink.Parent == nil {
		return true
	}
	
	if parentLink.Left != nil && parentLink.Left == node {
		parentValue.Left = node.Right

		
		if node.Right != nil {
			node.Right.Parent = &parentValue
			*(node.Right) = parentValue
		} else {
			node.Right = &parentValue
		}
	} else if (parentLink.Right != nil && parentLink.Right == node) {
		parentValue.Right = node.Left
		
		if node.Left != nil {
			node.Left.Parent = &parentValue
			*(node.Left) = parentValue
		} else {
			node.Left = &parentValue
		}
	}

	parentValue.Parent = node

	if parentLink.Parent == nil {
		node.Parent = nil // is root
		isRoot = true
	} else {
		node.Parent = parentLink.Parent
	}
	
	*parentLink = *node

	return isRoot
}

func splay[T Ordered](node *Node[T]) {
	isRoot := false
	for !isRoot {
		isRoot = zig(node)
	}
}