package main

// https://iq.opengenus.org/interval-tree/

import (
	"fmt"
)

type Comparable interface {
    compare(Comparable) int
	compareMax(Comparable) int
}

type Data struct {
	Left, Right, Max int
}

func NewData(left, right int) *Data {
	return &Data{Left: left, Right: right, Max: right}
}

func (item *Data) compare(other Comparable) int {
	
	otherData, ok := other.(*Data)
	if !ok {
		panic("other was not of type *Data")
	}
	
	if item.Left < otherData.Left {
		return -1
	} else if item.Left == otherData.Left {
		return 0
	}
	return 1
}

func (item *Data) compareMax(other Comparable) int {
	
	otherData, ok := other.(*Data)
	if !ok {
		panic("other was not of type *Data")
	}
	
	if item.Max < otherData.Max {
		return -1
	} else if item.Max == otherData.Max {
		return 0
	}
	return 1
}

// Node of a tree
type Node struct {
	Key         Comparable
	Height      int
	Left, Right *Node
}

// NewTree create a new AVL tree
func NewTree(key Comparable) *Node {
	return &Node{Key: key}
}

// Get : return node with given key
func (root *Node) Get(key Comparable) *Node {
	if root == nil {
		return nil
	}
	switch root.Key.compare(key) {
	case 0: // равны
		return root
	case -1: // key больше чем значение в root.Key
		root = root.Right
	case 1: // key меньше чем значение в root.Key
		root = root.Left
	}
	// if root.Key == key {
	// 	return root
	// } else if root.Key < key {
	// 	root = root.Right
	// } else {
	// 	root = root.Left
	// }
	return root.Get(key)
}

func (root *Node) search(key Comparable) bool {
	if root.Key.(*Data).Left >= (key.(*Data).Left) && root.Key.(*Data).Left <= (key.(*Data).Right) || root.Key.(*Data).Right >= (key.(*Data).Left) && root.Key.(*Data).Right <= (key.(*Data).Right){
		return true
	}
	if root.Left != nil {
		if root.Left.Key.(*Data).Max >= key.(*Data).Left {
			return root.Left.search(key)
		} 
	}

	if root.Right != nil {
		if root.Right.Key.(*Data).Max >= key.(*Data).Left {
			return root.Right.search(key)
		} 
	}
	return false
}


// func (root *Node) Propagation(key Comparable) {
// 	switch root.Key.compare(key) {
// 		case 0: // равны
// 			return
// 		case -1: // key больше чем значение в root.Key
// 			root.Key.(*Data).Max = key.(*Data).Max

// 		case 1: // key меньше чем значение в root.Key
// 		panic("ПОДУМАТЬ")	
// 		root = root.Left
// 		}
// }	

// Insert a new item
func (root *Node) Insert(key Comparable) *Node {
	// if *root == nil {
	// 	*root = &Node{
	// 		Key:    key,
	// 		Height: 1,
	// 	}
	// 	return
	// }
	
	switch root.Key.compare(key) {
		// case 0: // равны
		// 	return root
		case -1: // key больше чем значение в root.Key
			if root.Right != nil {
				root.Right = root.Right.Insert(key)
			} else {
				root.Right = &Node{Key: key, Height: 1}
			}
		
		case 1: // key меньше чем значение в root.Key
			if root.Left != nil {
				root.Left = root.Left.Insert(key)
			} else {
				root.Left = &Node{Key: key, Height: 1}
			}	
	}
	
	// if root.Key < key {
	// 	root.Right.Insert(key)
	// } else if root.Key > key {
	// 	root.Left.Insert(key)
	// }

	// update height
	root.Height = root.height()
	
	if root.Right != nil && root.Left != nil {
		switch root.Right.Key.compareMax(root.Left.Key) {
		case -1: // Left.Key больше Right.Key
			if root.Key.compareMax(root.Left.Key) == -1 {
				// максимум корня меньше чем максимум дочерних узлов(левый узел)
				root.Key.(*Data).Max = root.Left.Key.(*Data).Max
			} 
		case 1: // Right.Key больше Left.Key
			if root.Key.compareMax(root.Right.Key) == -1 {
				// максимум корня меньше чем максимум дочерних узлов(левый узел)
				root.Key.(*Data).Max = root.Right.Key.(*Data).Max
			}
		}
	} else if root.Right != nil {
		if root.Key.compareMax(root.Right.Key) == -1 {
			// максимум корня меньше чем максимум дочерних узлов(левый узел)
			root.Key.(*Data).Max = root.Right.Key.(*Data).Max
		}
	} else {
		if root.Key.compareMax(root.Left.Key) == -1 {
			// максимум корня меньше чем максимум дочерних узлов(левый узел)
			root.Key.(*Data).Max = root.Left.Key.(*Data).Max
		} 
	}

	bFactor := root.balanceFactor()

	if bFactor == 2 { // L
		bFactor = root.Left.balanceFactor()
		if bFactor == 1 { // LL
			root = root.llRotation()
			//fmt.Println()
		} else if bFactor == -1 { // LR
			//root.lrRotation()
			root = root.Left.rrRotation()
			root = root.llRotation()
		}
	} else if bFactor == -2 { // R
		bFactor = root.Right.balanceFactor()
		if bFactor == 1 { // RL
			//root.rlRotation()
			root = root.Right.llRotation()
			root = root.rrRotation()
		} else if bFactor == -1 { // RR
			root = root.rrRotation()
		}
	}
	return root
}

// // Delete : remove given key from the tree
// func (root *Node) Delete(key Comparable) {
// 	if root == nil {
// 		return
// 	}
// 	if root.Key < key {
// 		root.Right.Delete(key)
// 	} else if root.Key > key {
// 		root.Left.Delete(key)
// 	} else {
// 		// 3 cases
// 		// 1. No Child
// 		// 2. With One Child
// 		// 3. With Two Child
// 		if root.Left == nil && root.Right == nil {
// 			root = nil
// 		} else if root.Left == nil {
// 			root = root.Right
// 		} else if root.Right == nil {
// 			root = root.Left
// 		} else {
// 			minVal := root.Right.min()
// 			root.Key = minVal
// 			root.Delete(minVal)
// 		}
// 		return
// 	}

// 	// update height
// 	root.Height = root.height()

// 	bFactor := root.balanceFactor()

// 	if bFactor == 2 { // L
// 		switch root.Left.balanceFactor() {
// 		case 1: // LL
// 			root.llRotation()
// 		case -1: // LR
// 			root.lrRotation()
// 		case 0: //  LL OR LR
// 			root.llRotation()
// 		}
// 	} else if bFactor == -2 { // L
// 		switch root.Right.balanceFactor() {
// 		case 1: // RL
// 			root.rlRotation()
// 		case -1: // RR
// 			root.rrRotation()
// 		case 0: // RL OR RR
// 			root.rrRotation()
// 		}
// 	}
// }

// rotations
// 1. LL
// 2. LR
// 3. RR
// 4. RL
func (root *Node) llRotation() *Node {
	
	// проверить правый узел и root
	b := root.Left
	br := b.Right
	b.Right = root
	root.Left = br
	root.Height = root.height()
	b.Height = b.height()
	root = b
	return root
}
func (root *Node) lrRotation() {
	c := root.Left.Right
	cl := c.Left
	cr := c.Right

	c.Left = root.Left
	c.Right = root
	c.Left.Right = cl

	root.Left = cr

	root.Height = root.height()
	c.Left.Height = c.Left.height()
	c.Height = c.height()

	root = c

}
func (root *Node) rrRotation() *Node {
	// проверить левый узел и root
	b := root.Right
	bl := b.Left
	b.Left = root

	root.Right = bl
	root.Height = root.height()
	b.Height = b.height()
	root = b
	return root
}
func (root *Node) rlRotation() {
	c := root.Right.Left
	cl := c.Left
	cr := c.Right

	c.Right = root.Right
	c.Right.Left = cr
	c.Left = root
	root.Right = cl

	root.Height = root.height()
	c.Right.Height = c.Right.height()
	c.Height = c.height()
	root = c
}

// balanceFactor : -ve balance factor means subtree root is heavy toward left
// and +ve balance factor means subtree root is heavy toward right side
func (root *Node) balanceFactor() int {
	var leftHeight, rightHeight int
	if root.Left != nil {
		leftHeight = root.Left.Height
	}
	if root.Right != nil {
		rightHeight = root.Right.Height
	}
	return leftHeight - rightHeight
}

func (root *Node) height() int {
	if root == nil {
		return 0
	}
	var leftHeight, rightHeight int
	if root.Left != nil {
		leftHeight = root.Left.Height
	}
	if root.Right != nil {
		rightHeight = root.Right.Height
	}
	max := leftHeight
	if rightHeight > leftHeight {
		max = rightHeight
	}
	return 1 + max
}

func (root *Node) min() Comparable {
	if root.Left == nil {
		return root.Key
	}
	return root.Left.min()
}

func main() {
	fmt.Println()

	root := NewTree(NewData(15, 20))
	root.Insert(NewData(10, 30))
	root.Insert(NewData(17, 19))
	root.Insert(NewData(5, 20))
	root.Insert(NewData(12, 15))
	root.Insert(NewData(30, 40))
	root.Insert(NewData(3, 5))
	root.Insert(NewData(1, 2))
	root.Insert(NewData(35, 44))
	root.Insert(NewData(36, 50))
	fmt.Println()

	res1 := root.search(NewData(13, 35))
	fmt.Println(res1)

	res2 := root.search(NewData(-1, 0))
	fmt.Println(res2)

	res3 := root.search(NewData(50, 55))
	fmt.Println(res3)

	fmt.Println()
}