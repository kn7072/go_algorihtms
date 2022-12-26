package main

import (
	"fmt"
	//"github.com/kn7072/go_algorihtms/structure/tree/rbt"
)

type Comparable interface {
	More(b Comparable) bool
}
//type Value any

type Node struct {
	Left, Right, Parrent *Node
	Color bool
	Value Comparable
}

type Tree struct {
	Root *Node
}

func (n *Node) getColorUncle() (bool, *Node) {
	grandParent := n.Parrent.Parrent
	if grandParent.Left == n {
		// родитель n является левым узлом деда
		if grandParent.Right == nil {
			return true, nil // черный
		} else {
			return grandParent.Right.Color, grandParent.Right  // какой-то
		}
		
	} else {
		// родитель n является правым узлом деда
		if grandParent.Left == nil {
			return true, nil // черный
		} else {
			return grandParent.Left.Color, grandParent.Left // какой-то
		}
	}
}

func (n *Node) getOrientation() (bool, bool){
	// true - левый, false - правый
	parrent := n.Parrent
	orientationNode, orientationParrent := false, false
	if n == parrent.Right {
		// нода правая у родителя
		orientationNode = true
	}
	if parrent == parrent.Parrent.Right {
		// родитель правый у деда
		orientationParrent = true
	}

	return orientationNode, orientationParrent
}


func (i *Node) Balance() {
	// if parrent is black node is ok
	// Color = false - красный цвет
	// Color = true - черный цвет
	
	if i.Parrent != nil {
		if !i.Parrent.Color{
			// родитель красный
			
			colorUncle, uncle := i.getColorUncle()
			if !colorUncle {
				// CASE 1  https://habr.com/ru/company/otus/blog/472040/
				// дядя красный
				uncle.Color = true // дядя стал черным
				i.Parrent.Color = true // родитель стал черным
				i.Parrent.Parrent.Color = false // дед стал красным
				i.Parrent.Parrent.Balance()
			} else {
				// дядя черный
				orientationNode, orientationParrent := i.getOrientation()
				if orientationNode == orientationParrent {
					// нода и родитель на одной стороне
					// CASE 3
					
					tempNode := i.Parrent.Parrent // дед
						i.Parrent.Color = true // родитель черный
						tempNode.Color = false // дед стал красным
						i.Parrent.Parrent = tempNode.Parrent // родитель занял место деда и получил родителя от деда
					
						if orientationNode {
							// левая ориентация
							tempNode.Left = i.Parrent.Right // правая нода родителя стала левой нодой деда
							i.Parrent.Right = tempNode // дед стал правой нодой родителя
							tempNode.Parrent = i.Parrent // у деда - родитель - родидель
						} else {
							// правая ориентация
							tempNode.Right = i.Parrent.Left // левая нода родителя стала правой нодой деда
							i.Parrent.Left = tempNode // дед стал левой нодой родителя
							tempNode.Parrent = i.Parrent // у деда - родитель - родидель
						}
				} else {
					// CASE 2
					if orientationNode {
						// нода находится слева от родителя, родитель - правая нода деда
						// 
					}
				}
			}
			
		} else {
			return
		}

	} else {
		// корень всегда черный
		i.Color = true
	}

	// Case 1

}
func (t *Node) AddNode(v Comparable) {
	
	newNode := new(Node)
	// GO to right
	if v.More(t.Value) {
		if t.Right != nil {
			t.Right.AddNode(v)
		} else {
			newNode = &Node{Value: v, Parrent: t, Color: false}
			t.Right = newNode
		}
	} else {
		// GO to left
		if t.Left != nil {
			t.Left.AddNode(v)
		} else {
			newNode = &Node{Value: v, Parrent: t, Color: false}
			t.Left = newNode
		}
	}
	newNode.Balance()
}

func (n *Node) PrintValue() {
	if n.Left != nil {
		n.Left.PrintValue()
	}
	fmt.Println(n.Value)
	if n.Right != nil {
		n.Right.PrintValue()
	}
}

type MyType struct {
	Value int
}

func (i *MyType) String() int {
	return i.Value
}

func (i *MyType) More(a Comparable) bool{
	return i.Value > a.(*MyType).Value
}



func main() {
	// tree := rbt.Tree{&rbt.Node{Value: 1, Color: true}}
	valueRoot := Comparable(&MyType{Value: 1})
	rootNode := &Node{Value: valueRoot,
					   Color: true}
	tree := Tree{rootNode}

	rootNode.AddNode(Comparable(&MyType{Value: 0}))
	rootNode.AddNode(Comparable(&MyType{Value: 2}))
	rootNode.AddNode(Comparable(&MyType{Value: -1}))
	rootNode.AddNode(Comparable(&MyType{Value: -4}))
	rootNode.PrintValue()
	fmt.Println(tree)

}