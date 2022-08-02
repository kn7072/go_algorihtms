package rbt

import (
	"fmt"
)

type Value any

type Node struct {
	Left, Right *Node
	Color bool
	Value Value
}

type Tree struct {
	Root *Node
}

func (n *Node) PrintValue() {
	fmt.Println(n.Value)
}

