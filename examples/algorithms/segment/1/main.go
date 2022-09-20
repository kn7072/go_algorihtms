package main

import (
	"fmt"
	"math"
)

// https://github.com/ideahitme/segment

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 //|
	//~string
}

// type Ordered[T any] interface {
// 	Less(T) bool
// 	//comparable
// }

type Node[T Ordered] struct { // Ordered[T]
	value T
}

var defaulNode = Node[int]{value: 100000} // 0 ДЛЯ СУММЫ, 10000 ДЛЯ МИНИМУМА

// func Action[T Ordered](l, r T) T {
// 	// sum
// 	return l + r
// }

func Action[T Ordered](l, r T) T {
	// min
	return T(math.Min(float64(l), float64(r)))
}

// Tree implementation of segment tree
type Tree[T Ordered] struct {
	nodes []*Node[T] // elements of the tree
	size  int        // size number of elements in the original array
}

func calcTreeSize(originalSize int) int {
	// fmt.Println(math.Log2(float64(originalSize)))
	// fmt.Println(uint(math.Ceil(math.Log2(float64(originalSize)))+1))
	// fmt.Println(1<<uint(math.Ceil(math.Log2(float64(originalSize)))+1) - 1)

	return 1<<uint(math.Ceil(math.Log2(float64(originalSize)))+1) - 1
}

func (t *Tree[T]) build(from []*Node[T], nodeIndex, leftBound, rightBound int) {
	if leftBound == rightBound {
		t.nodes[nodeIndex] = from[leftBound]
		return
	}

	bisect := (leftBound + rightBound) / 2
	t.build(from, 2*nodeIndex+1, leftBound, bisect)
	t.build(from, 2*nodeIndex+2, bisect+1, rightBound)

	leftChild := t.nodes[2*nodeIndex+1]
	rightChild := t.nodes[2*nodeIndex+2]

	t.nodes[nodeIndex] = &Node[T]{value: Action(leftChild.value, rightChild.value)}
}

// NewTree constructs a segment tree and allows to perform RMQ on provided targetArray
func NewTree[T Ordered](arr []*Node[T]) *Tree[T] {
	treeSize := calcTreeSize(len(arr))
	nodes := make([]*Node[T], treeSize)

	t := &Tree[T]{nodes, len(arr)}
	t.build(arr, 0, 0, len(arr)-1)

	return t
}

func (t *Tree[T]) get(nodeIndex, leftNode, rightNode, leftQuery, rightQuery int) T {
	if leftQuery > rightQuery {
		return T(defaulNode.value)
	}

	if leftQuery == leftNode && rightQuery == rightNode {
		return t.nodes[nodeIndex].value
	}

	bisect := (leftNode + rightNode) / 2
	vLeft := t.get(2*nodeIndex+1, leftNode, bisect, leftQuery, int(math.Min(float64(rightQuery), float64(bisect))))
	vRight := t.get(2*nodeIndex+2, bisect+1, rightNode, int(math.Max(float64(leftQuery), float64(bisect+1))), rightQuery)

	return Action(vLeft, vRight)
}

func (t *Tree[T]) update(nodeIndex, leftNode, rightNode int, position int, newValue T) {
	if leftNode == rightNode {
		t.nodes[nodeIndex].value = newValue
	} else {
		bisect := (leftNode + rightNode) / 2
		if position <= bisect {
			t.update(2*nodeIndex+1, leftNode, bisect, position, newValue)
		} else {
			t.update(2*nodeIndex+2, bisect+1, rightNode, position, newValue)
		}
		leftChild := t.nodes[2*nodeIndex+1]
		rightChild := t.nodes[2*nodeIndex+2]
		Action(leftChild.value, rightChild.value)
	}
}

func main() {
	arr := []*Node[int]{{1},
		{3},
		{5},
		{4},
		{6},
		{10},
		{200},
		{-100},
		{-200},
	}
	tree := NewTree(arr)
	fmt.Println(tree)
	res := tree.get(0, 0, len(arr)-1, 2, 5)
	fmt.Println(res)

	res = tree.get(0, 0, len(arr)-1, 2, 7)
	fmt.Println(res)

	tree.update(0, 0, len(arr)-1, 2, -1)
	res = tree.get(0, 0, len(arr)-1, 2, 5)
	fmt.Println(res)
	fmt.Println()
}
