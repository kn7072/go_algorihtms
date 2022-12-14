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
	node7 := &Node[int]{Value: 15}
	node8 := &Node[int]{Value: 25}
	node9 := &Node[int]{Value: 39}
	node10 := &Node[int]{Value: 45}
	node11 := &Node[int]{Value: 43}
	node12 := &Node[int]{Value: 46}

	insert[int](&root, &node1)
	insert[int](&root, &node2)
	insert[int](&root, &node3)
	insert[int](&root, &node4)
	insert[int](&root, &node5)
	insert[int](&root, &node6)
	insert[int](&root, &node7)
	insert[int](&root, &node8)
	insert[int](&root, &node9)
	insert[int](&root, &node10)
	insert[int](&root, &node11)
	insert[int](&root, &node12)

	// test find
	findNode := find(&root, 50)
	fmt.Println(findNode.Value)
	findNode = find(&root, 20)
	fmt.Println(findNode)
	//----------------------

	// findNode = find(&root, 2)
	// zig(findNode)

	// findNode = find(&root, 20)
	// isRoot := zig(findNode)
	// fmt.Println(isRoot)

	// findNode = find(&root, 50)
	// isRoot := zig(findNode)
	// fmt.Println(isRoot)

	// findNode = find(&root, 45)
	// isRoot := zig(findNode)
	// fmt.Println(isRoot)

	// findNode = find(&root, 200)
	// isRoot := zig(findNode)
	// fmt.Println(isRoot)

	findNode = find(&root, 25)
	isRoot := zig(findNode)
	fmt.Println(isRoot)

	// findNode = find(&root, 200)
	// isRoot = zig(findNode)
	// fmt.Println(isRoot)

	
	splay(findNode)
	fmt.Println()
}