
// Package avl is a Adelson-Velskii and Landis tree implemnation
// avl is self-balancing tree, i.e for all node in a tree, height difference
// between its left and right child will not exceed 1
// more information : https://en.wikipedia.org/wiki/AVL_tree
package main


type Value struct {
	Left, Right, Max int
}

func NewValue(left, right int) *Value {
	return &Value{Left: left, Right: right, Max: right}
}

type Comparable interface{
	More(Value) int
}

func (v *Value) More(key *Value) int {
	if v.Left > key.Left {
		return 1
	} else if v.Left < key.Left {
		return -1
	}
	return 0
}

func overlap(root *Node, key *Value) bool {
	
	if root == nil {
		return false
	}

	if key.Left >= root.Key.Left && key.Right <= root.Key.Right {
		// key входит в root
		return true
	}
	if key.Left <= root.Key.Left && key.Right >= root.Key.Left {
		// пересечение с левой границей
		return true
	}

	if key.Left <= root.Key.Right && key.Right >= root.Key.Right {
		// пересечение с правой границей
		return true
	}

	res := false
	if root.Left != nil {
		if key.Left <= root.Left.Key.Max {
			return overlap(root.Left, key)
		}
	}
	if root.Right != nil {
		if key.Left <= root.Right.Key.Max {
			res = overlap(root.Right, key)
		}
	}
	return res
}


// Node of a tree
type Node struct {
	Key         Value
	Height      int
	Left, Right *Node
}

// NewTree create a new AVL tree
func NewTree() *Node {
	return nil
}

// Get : return node with given key
func Get(root *Node, key *Value) *Node {
	if root == nil {
		return nil
	}
	switch root.Key.More(key){
	case 0:
		return root
	case -1:
		root = root.Right
	case 1:
		root = root.Left
	}
	
	return Get(root, key)
}

// Insert a new item
func Insert(root **Node, key *Value) {
	if *root == nil {
		*root = &Node{
			Key:    *key,
			Height: 1,
		}
		return
	}
	switch (*root).Key.More(key){
	case -1, 0:
		Insert(&(*root).Right, key)
	case 1:
		Insert(&(*root).Left, key)
	}


	// update height
	(*root).Height = height(*root)
	
	// update max of interval
	hook(*root)

	bFactor := balanceFactor(*root)

	if bFactor == 2 { // L
		bFactor = balanceFactor((*root).Left)
		if bFactor == 1 { // LL
			llRotation(root)
		} else if bFactor == -1 { // LR
			lrRotation(root)
		}
	} else if bFactor == -2 { // R
		bFactor = balanceFactor((*root).Right)
		if bFactor == 1 { // RL
			rlRotation(root)
		} else if bFactor == -1 { // RR
			rrRotation(root)
		}
	}
}

// Delete : remove given key from the tree
func Delete(root **Node, key *Value) {
	if root == nil {
		return
	}

	switch (*root).Key.More(key){
	case 0:
		// 3 cases
		// 1. No Child
		// 2. With One Child
		// 3. With Two Child
		if (*root).Left == nil && (*root).Right == nil {
			*root = nil
		} else if (*root).Left == nil {
			*root = (*root).Right
		} else if (*root).Right == nil {
			*root = (*root).Left
		} else {
			minVal := min((*root).Right)
			(*root).Key = minVal
			//Delete(root, &minVal)
			Delete(&(*root).Right, &minVal)
			hook(*root)
		}
		return
	case -1:
		Delete(&(*root).Right, key)
	case 1:
		Delete(&(*root).Left, key)
	}

	// update height
	(*root).Height = height(*root)
	// update max of interval
	hook(*root)

	bFactor := balanceFactor(*root)

	if bFactor == 2 { // L
		switch balanceFactor((*root).Left) {
		case 1: // LL
			llRotation(root)
		case -1: // LR
			lrRotation(root)
		case 0: //  LL OR LR
			llRotation(root)
		}
	} else if bFactor == -2 { // L
		switch balanceFactor((*root).Right) {
		case 1: // RL
			rlRotation(root)
		case -1: // RR
			rrRotation(root)
		case 0: // RL OR RR
			rrRotation(root)
		}
	}
}

// rotations
// 1. LL
// 2. LR
// 3. RR
// 4. RL
func llRotation(root **Node) {
	b := (*root).Left
	br := b.Right
	b.Right = *root
	(*root).Left = br
	(*root).Height = height(*root)
	b.Height = height(b)
	*root = b

	hook((*root).Right)
	hook(*root)
}
func lrRotation(root **Node) {
	rrRotation(&(*root).Left)
	llRotation(root)
	
}
func rrRotation(root **Node) {
	b := (*root).Right
	bl := b.Left
	b.Left = *root

	(*root).Right = bl
	(*root).Height = height(*root)
	b.Height = height(b)
	*root = b

	hook((*root).Left)
	hook(*root)

}
func rlRotation(root **Node) {
	llRotation(&(*root).Right)
	rrRotation(root)
}

// balanceFactor : -ve balance factor means subtree root is heavy toward left
// and +ve balance factor means subtree root is heavy toward right side
func balanceFactor(root *Node) int {
	var leftHeight, rightHeight int
	if root.Left != nil {
		leftHeight = root.Left.Height
	}
	if root.Right != nil {
		rightHeight = root.Right.Height
	}
	return leftHeight - rightHeight
}

func height(root *Node) int {
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

func min(root *Node) Value {
	if root.Left == nil {
		return root.Key
	}
	return min(root.Left)
}


func hook(root *Node) {
	var maxLeft, maxRight, max int
	
	if root.Left == nil && root.Right == nil {
		root.Key.Max = root.Key.Right
		return
	}
	if root.Left != nil {
		maxLeft = root.Left.Key.Max
	}

	if root.Right != nil {
		maxRight = root.Right.Key.Max
	}
	if maxRight > maxLeft {
		max = maxRight
	} else {
		max = maxLeft
	}
	if root.Key.Right < max {
		root.Key.Max = max
	}
}