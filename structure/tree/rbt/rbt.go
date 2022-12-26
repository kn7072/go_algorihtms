package rbt

import (
	"fmt"
)

type Comparable interface {
	More(b Comparable) bool
}
//type Value any

type Node struct {
	Left, Right *Node
	Color bool
	Value Comparable
}

type Tree struct {
	Root *Node
}

func (t *Node) AddNode(v Comparable) {
	
	// GO to right
	if v.More(t.Value) {
		if t.Right != nil {
			t.Right.AddNode(v)
		} else {
			newNode := &Node{Value: v, Color: false}
			t.Right = newNode
		}
	} else {
		// GO to left
		if t.Left != nil {
			t.Left.AddNode(v)
		} else {
			newNode := &Node{Value: v, Color: false}
			t.Left = newNode
		}
	}
}

func (n *Node) PrintValue() {
	fmt.Println(n.Value)
}

