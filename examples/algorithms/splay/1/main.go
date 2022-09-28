package main

import (
	"fmt"
)

func main() {
	root := NewTree[int](100)
	node1 := &Node[int]{Value: 50}
	node2 := &Node[int]{Value: 30}
	node3 := &Node[int]{Value: 60}
	node4 := &Node[int]{Value: 200}
	node5 := &Node[int]{Value: 20}
	node6 := &Node[int]{Value: 40}

	insert[int](&root, &node1)
	insert[int](&root, &node2)
	insert[int](&root, &node3)
	insert[int](&root, &node4)
	insert[int](&root, &node5)
	insert[int](&root, &node6)

	// test find
	findNode := find(&root, 50)
	fmt.Println(findNode.Value)
	findNode = find(&root, 100)
	fmt.Println(findNode)
	//----------------------

	// findNode = find(&root, 2)
	// zig(findNode)

	findNode = find(&root, 20)
	zig(findNode)

	fmt.Println()
}