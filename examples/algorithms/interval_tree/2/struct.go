package main

import "fmt"


type Value struct {
	Value int
}

type Comparable interface {
	More(Value, Value) int
}

func (v *Value) More(second *Value) int {
	if v.Value < second.Value {
		return 1
	} else if v.Value == second.Value {
		return 0
	}
	return -1
}

type Node struct {
	Height int
	Left, Right *Node
	Key *Value
}

func searchSuccessor(node, parent *Node) *Node {
	successor := &Node{}
	if node.Left != nil {
		successor = searchSuccessor(node.Left, node)
		node.Height = heightNode(node)
		balanceNode(node)
	} else {
		// больше нет левых дочерних узлов
		*successor = *node
		if node.Right != nil {
			*node = *node.Right
		} else {
			parent.Left = nil
		}

	}
	return successor
}

func deleteNode(root, parentRoot *Node, delNode *Value) *Node {
	if root == nil {
		return nil
	}
	
	var foundNode *Node

	switch root.Key.More(delNode) {
	case 0:
		// нашли удаляемый элемент

		foundNode = root
		if root.Left != nil && root.Right != nil {
			successor := searchSuccessor(root.Right, root)
			successor.Left = root.Left
			successor.Right = root.Right
			successor.Height = heightNode(successor)
			if root == parentRoot.Left {
				parentRoot.Left = successor
			} else {
				parentRoot.Right = successor
			}
			*root = *successor

		} else if root.Left != nil {
			// только левый дочерний узел
			*root = *root.Left
		} else if root.Right != nil {
			// только правый дочерный узел
			*root = *root.Right
		} else {
			// лист
			if parentRoot.Left == root {
				parentRoot.Left = nil
			} else {
				parentRoot.Right = nil
			}
		}
	case 1:
		// ищем в правом поддереве
		foundNode = deleteNode(root.Right, root, delNode)
	case -1:
		// ищем в левом поддереве
		foundNode = deleteNode(root.Left, root, delNode)
	}

	if foundNode != nil {
		root.Height = heightNode(root)
		balanceNode(root)
	}
	return foundNode
}

func insert(node *Node, addedValue *Value) {

	switch addedValue.More(node.Key) {
	case 1: // addedValue < node.Key
		if node.Left != nil {
			insert(node.Left, addedValue)
			node.Left.Height = heightNode(node.Left)
		} else {
			node.Left = &Node{
				Key: addedValue,
				Height: 1,
			}
			node.Height = heightNode(node)
		}
	default:
		if node.Right != nil {
			insert(node.Right, addedValue)
			node.Right.Height = heightNode(node.Right)
		} else {
			node.Right = &Node{
				Key: addedValue,
				Height: 1,
			}
			node.Height = heightNode(node)
		}
	}

	balanceNode(node)
}

func rotationL(node *Node) {
	leftChild := node.Left
	leftrRghtChild := leftChild.Right
	node.Left = leftrRghtChild
	node.Height = heightNode(node)
	tempValueNode := *node
	leftChild.Right = &tempValueNode

	*node = *leftChild
	node.Height = heightNode(node)
}

func rotationR(node *Node) {
	rightChild := node.Right
	rightLeftChild := rightChild.Left
	node.Right = rightLeftChild
	node.Height = heightNode(node)
	tempValueNode := *node
	rightChild.Left = &tempValueNode
	
	*node = *rightChild
	node.Height = heightNode(node)
}

func balanceNode(node *Node) {
	diffHeight := diffHeightNode(node)
	switch diffHeight {
	case 2: // левый узел имеет большую высоту
		switch diffHeightNode(node.Left) {
		case 1: // ll
			rotationL(node)
		case -1: // lr
			rotationR(node.Left)
			rotationL(node)
		}
	case -2:
		switch diffHeightNode(node.Right) {
		case 1:  // rl
			rotationL(node.Right)
			rotationR(node)
		case -1: // rr
			rotationR(node)
			fmt.Println()
		}
	}
}

func diffHeightNode(node *Node) int {
	hLeft, hRight := 0, 0
	if node.Left != nil {
		hLeft = node.Left.Height
	}
	if node.Right != nil {
		hRight = node.Right.Height
	}
	return hLeft - hRight
}

func heightNode(node *Node) int {
	height := 0

	if node.Left != nil {
		height = node.Left.Height
	}
	if node.Right != nil {
		if node.Right.Height > height {
			height = node.Right.Height
		}
	}
	return height + 1
}